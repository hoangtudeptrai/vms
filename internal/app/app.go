package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/hoangtu1372k2/vms/docs/swagger"
	"github.com/hoangtu1372k2/vms/internal/auth"
	controllers "github.com/hoangtu1372k2/vms/internal/controller"
	"github.com/hoangtu1372k2/vms/internal/model"
	"gorm.io/gorm"

	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"net/http/pprof"

	"github.com/hoangtu1372k2/common-go/reposity"
	"github.com/hoangtu1372k2/common-go/telemetry"
)

var (
	ErrShutdown = fmt.Errorf("application shutdown gracefully")
)

var srv *http.Server
var runCtx context.Context
var runCancel context.CancelFunc
var cfg *viper.Viper
var log *logrus.Logger
var stats = telemetry.New()
var db *gorm.DB
var TrustedProxies = []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"}

// Run application.
func Run(c *viper.Viper) error {
	var err error

	runCtx, runCancel = context.WithCancel(context.Background())
	cfg = c
	log = logrus.New()

	if cfg.GetBool("debug") {
		log.Level = logrus.DebugLevel
		log.Debug("Enabling Debug Logging")
	}
	if cfg.GetBool("trace") {
		log.Level = logrus.TraceLevel
		log.Debug("Enabling Trace Logging")
	}
	if cfg.GetBool("disable_logging") {
		log.Level = logrus.FatalLevel
	}

	if cfg.GetBool("enable_sql") {
		log.Infof("Connecting to DB")
		db, err = reposity.Connect(
			cfg.GetString("sql_host"),
			cfg.GetString("sql_port"),
			cfg.GetString("sql_dbname"),
			cfg.GetString("sql_sslmode"),
			cfg.GetString("sql_user"),
			cfg.GetString("sql_password"),
			cfg.GetString("sql_schema"))
		if err != nil {
			return fmt.Errorf("could not establish database connection - %s", err)
		}
	}
	if !reposity.Connected {
		log.Infof("SQL DB not configured, skipping")
	} else {
		err = reposity.Ping()
		if err != nil {
			log.Errorf(err.Error())
		} else {
			log.Info("Successfully connected to database")
			// Migrate tables
			err = reposity.Migrate(
				&model.User{},
				&model.Course{},
				&model.Assignment{},
				&model.Notification{},
				&model.Profile{},
				&model.Message{},
				&model.Grade{},
				&model.AssignmentSubmission{},
				&model.AssignmentSubmissionFile{},
				&model.AssignmentDocument{},
				&model.CourseDocument{},
				&model.CourseEnrollment{},
				&model.Lesson{},
			)
			if err != nil {
				panic("Failed to AutoMigrate table! err: " + err.Error())
			}
		}
	}
	defer reposity.Close()
	controllers.InitMinIO()

	// Cấu hình Swagger
	swagger.SwaggerInfo.Title = "VMS"
	swagger.SwaggerInfo.Description = "A video management service API."
	swagger.SwaggerInfo.Version = "1.0"
	swagger.SwaggerInfo.Host = cfg.GetString("base_url")
	swagger.SwaggerInfo.BasePath = cfg.GetString("service_path")
	if strings.Contains(cfg.GetString("service_path"), "127.0.0.1") || strings.Contains(cfg.GetString("service_path"), "localhost") {
		swagger.SwaggerInfo.Schemes = []string{"https", "http"}
	} else {
		swagger.SwaggerInfo.Schemes = []string{"http", "https"}
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowCredentials:       true,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowOrigins:           []string{"*"},
		AllowMethods:           []string{"POST", "PUT", "PATCH", "DELETE", "GET", "OPTIONS", "UPDATE"},
		AllowHeaders: []string{
			"Content-Type, content-length, accept-encoding, X-CSRF-Token, " +
				"access-control-allow-origin, Authorization, X-Max, access-control-allow-headers, " +
				"accept, origin, Cache-Control, X-Requested-With, X-Request-Source"},
		MaxAge: 12 * time.Hour,
	}))

	apiV0 := router.Group(cfg.GetString("service_path"))
	{
		// Initialize auth service and handler
		authService := &auth.Service{
			DB:         db,
			JWTSecret:  []byte(cfg.GetString("jwt_secret")),
			ExpireTime: 24 * time.Hour,
		}
		authHandler := auth.NewHandler(authService)

		// Public routes (không yêu cầu xác thực)
		apiV0.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

		// Debug endpoints
		if cfg.GetBool("enable_pprof") {
			apiV0.GET("/debug/pprof/", gin.WrapH(http.HandlerFunc(pprof.Index)))
			apiV0.GET("/debug/pprof/cmdline", gin.WrapH(http.HandlerFunc(pprof.Cmdline)))
			apiV0.GET("/debug/pprof/profile", gin.WrapH(http.HandlerFunc(pprof.Profile)))
			apiV0.GET("/debug/pprof/symbol", gin.WrapH(http.HandlerFunc(pprof.Symbol)))
			apiV0.GET("/debug/pprof/trace", gin.WrapH(http.HandlerFunc(pprof.Trace)))
			apiV0.GET("/debug/pprof/allocs", gin.WrapH(pprof.Handler("allocs")))
			apiV0.GET("/debug/pprof/mutex", gin.WrapH(pprof.Handler("mutex")))
			apiV0.GET("/debug/pprof/goroutine", gin.WrapH(pprof.Handler("goroutine")))
			apiV0.GET("/debug/pprof/heap", gin.WrapH(pprof.Handler("heap")))
			apiV0.GET("/debug/pprof/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))
			apiV0.GET("/debug/pprof/block", gin.WrapH(pprof.Handler("block")))
		}

		apiV0.POST("/auth/login", handleWrapper(authHandler.Login, false))

		// User routes
		apiV0.GET("/users", handleWrapper(controllers.Get, false))
		apiV0.POST("/users", handleWrapper(controllers.CreateUser, false))
		apiV0.GET("/users/:id", handleWrapper(controllers.GetUserByID, false))
		apiV0.PUT("/users/:id", handleWrapper(controllers.UpdateUser, false))
		apiV0.DELETE("/users/:id", handleWrapper(controllers.DeleteUser, false))
		apiV0.GET("/users/me", handleWrapper(controllers.GetCurrentUser, false))

		// Course routes
		apiV0.GET("/courses", handleWrapper(controllers.GetCourses, false))
		apiV0.POST("/courses", handleWrapper(controllers.CreateCourse, false))
		apiV0.GET("/courses/:id", handleWrapper(controllers.GetCourseByID, false))
		apiV0.PUT("/courses/:id", handleWrapper(controllers.UpdateCourse, false))
		apiV0.DELETE("/courses/:id", handleWrapper(controllers.DeleteCourse, false))
		apiV0.GET("/courses/instructor", handleWrapper(controllers.GetCoursesByInstructor, false))
		apiV0.GET("/courses/enrolled", handleWrapper(controllers.GetEnrolledCourses, false))

		// Course Enrollment routes
		apiV0.GET("/enrollments", handleWrapper(controllers.GetCourseEnrollments, false))
		apiV0.POST("/enrollments", handleWrapper(controllers.CreateCourseEnrollment, false))
		apiV0.GET("/enrollments/:id", handleWrapper(controllers.GetCourseEnrollmentByID, false))
		apiV0.PUT("/enrollments/:id", handleWrapper(controllers.UpdateCourseEnrollment, false))
		apiV0.DELETE("/enrollments/:id", handleWrapper(controllers.DeleteCourseEnrollment, false))
		apiV0.GET("/enrollments/student/:id", handleWrapper(controllers.GetStudentEnrollments, false))
		apiV0.GET("/enrollments/course/:id", handleWrapper(controllers.GetCourseEnrollmentsByCourse, false))

		// Course Document routes
		apiV0.GET("/course-documents", handleWrapper(controllers.GetCourseDocuments, false))
		apiV0.POST("/course-documents", handleWrapper(controllers.CreateCourseDocument, false))
		apiV0.GET("/course-documents/:id", handleWrapper(controllers.GetCourseDocumentByID, false))
		apiV0.PUT("/course-documents/:id", handleWrapper(controllers.UpdateCourseDocument, false))
		apiV0.DELETE("/course-documents/:id", handleWrapper(controllers.DeleteCourseDocument, false))
		apiV0.GET("/course-documents/course/:id", handleWrapper(controllers.GetCourseDocumentsByCourse, false))

		// Assignment routes
		apiV0.GET("/assignments", handleWrapper(controllers.GetAssignments, false))
		apiV0.POST("/assignments", handleWrapper(controllers.CreateAssignment, false))
		apiV0.GET("/assignments/:id", handleWrapper(controllers.GetAssignmentByID, false))
		apiV0.PUT("/assignments/:id", handleWrapper(controllers.UpdateAssignment, false))
		apiV0.DELETE("/assignments/:id", handleWrapper(controllers.DeleteAssignment, false))
		apiV0.GET("/assignments/course/:course_id", handleWrapper(controllers.GetAssignmentsByCourseID, false))

		// Assignment Submission routes
		apiV0.GET("/assignment-submissions", handleWrapper(controllers.GetAssignmentSubmissions, false))
		apiV0.POST("/assignment-submissions", handleWrapper(controllers.CreateAssignmentSubmission, false))
		apiV0.GET("/assignment-submissions/:id", handleWrapper(controllers.GetAssignmentSubmissionByID, false))
		apiV0.PUT("/assignment-submissions/:id", handleWrapper(controllers.UpdateAssignmentSubmission, false))
		apiV0.DELETE("/assignment-submissions/:id", handleWrapper(controllers.DeleteAssignmentSubmission, false))
		apiV0.GET("/assignment-submissions/assignment/:id", handleWrapper(controllers.GetAssignmentSubmissionsByAssignment, false))
		apiV0.GET("/assignment-submissions/student/:id", handleWrapper(controllers.GetAssignmentSubmissionsByStudent, false))

		// Assignment Submission File routes
		apiV0.GET("/assignment-submission-files", handleWrapper(controllers.GetAssignmentSubmissionFiles, false))
		apiV0.POST("/assignment-submission-files", handleWrapper(controllers.CreateAssignmentSubmissionFile, false))
		apiV0.GET("/assignment-submission-files/:id", handleWrapper(controllers.GetAssignmentSubmissionFileByID, false))
		apiV0.PUT("/assignment-submission-files/:id", handleWrapper(controllers.UpdateAssignmentSubmissionFile, false))
		apiV0.DELETE("/assignment-submission-files/:id", handleWrapper(controllers.DeleteAssignmentSubmissionFile, false))
		apiV0.GET("/assignment-submission-files/submission/:id", handleWrapper(controllers.GetAssignmentSubmissionFilesBySubmission, false))

		// Assignment Document routes
		apiV0.GET("/assignment-documents", handleWrapper(controllers.GetAssignmentDocuments, false))
		apiV0.POST("/assignment-documents", handleWrapper(controllers.CreateAssignmentDocument, false))
		apiV0.GET("/assignment-documents/:id", handleWrapper(controllers.GetAssignmentDocumentByID, false))
		apiV0.PUT("/assignment-documents/:id", handleWrapper(controllers.UpdateAssignmentDocument, false))
		apiV0.DELETE("/assignment-documents/:id", handleWrapper(controllers.DeleteAssignmentDocument, false))
		apiV0.GET("/assignment-documents/assignment/:id", handleWrapper(controllers.GetAssignmentDocumentsByAssignment, false))

		// Lesson routes
		apiV0.GET("/lessons", handleWrapper(controllers.GetLessons, false))
		apiV0.POST("/lessons", handleWrapper(controllers.CreateLesson, false))
		apiV0.GET("/lessons/:id", handleWrapper(controllers.GetLessonByID, false))
		apiV0.PUT("/lessons/:id", handleWrapper(controllers.UpdateLesson, false))
		apiV0.DELETE("/lessons/:id", handleWrapper(controllers.DeleteLesson, false))
		apiV0.GET("/lessons/course/:id", handleWrapper(controllers.GetLessonsByCourse, false))

		// Comment routes
		apiV0.GET("/comments", handleWrapper(controllers.GetComments, false))
		apiV0.POST("/comments", handleWrapper(controllers.CreateComment, false))
		apiV0.GET("/comments/:id", handleWrapper(controllers.GetCommentByID, false))
		apiV0.PUT("/comments/:id", handleWrapper(controllers.UpdateComment, false))
		apiV0.DELETE("/comments/:id", handleWrapper(controllers.DeleteComment, false))
		apiV0.GET("/comments/user/:id", handleWrapper(controllers.GetUserComments, false))
		apiV0.GET("/comments/submission/:id", handleWrapper(controllers.GetSubmissionComments, false))

		// Grade routes
		apiV0.GET("/grades", handleWrapper(controllers.GetGrades, false))
		apiV0.POST("/grades", handleWrapper(controllers.CreateGrade, false))
		apiV0.GET("/grades/:id", handleWrapper(controllers.GetGradeByID, false))
		apiV0.PUT("/grades/:id", handleWrapper(controllers.UpdateGrade, false))
		apiV0.DELETE("/grades/:id", handleWrapper(controllers.DeleteGrade, false))
		apiV0.GET("/grades/student/:id", handleWrapper(controllers.GetStudentGrades, false))
		apiV0.GET("/grades/assignment/:id", handleWrapper(controllers.GetAssignmentGrades, false))
		apiV0.GET("/grades/course/:id", handleWrapper(controllers.GetCourseGrades, false))

		// Submission routes
		apiV0.GET("/submissions", handleWrapper(controllers.GetSubmissions, false))
		apiV0.POST("/submissions", handleWrapper(controllers.CreateSubmission, false))
		apiV0.GET("/submissions/:id", handleWrapper(controllers.GetSubmissionByID, false))
		apiV0.PUT("/submissions/:id", handleWrapper(controllers.UpdateSubmission, false))
		apiV0.DELETE("/submissions/:id", handleWrapper(controllers.DeleteSubmission, false))
		apiV0.GET("/submissions/student/:id", handleWrapper(controllers.GetStudentSubmissions, false))
		apiV0.GET("/submissions/latest/:student_id/:assignment_id", handleWrapper(controllers.GetLatestSubmission, false))

		// Notification routes
		apiV0.GET("/notifications", handleWrapper(controllers.GetNotifications, false))

		// Other utility routes
		apiV0.GET("/health", handleWrapper(controllers.Health, false))
		apiV0.POST("/upload", handleWrapper(controllers.UploadFile, false))
		apiV0.GET("/file", handleWrapper(controllers.GetFile, false))

		srv = &http.Server{
			Addr:    cfg.GetString("listen_addr"),
			Handler: router,
		}

		// Xử lý tín hiệu shutdown
		go func() {
			trap := make(chan os.Signal, 1)
			signal.Notify(trap, syscall.SIGTERM, syscall.SIGINT)
			s := <-trap
			log.Infof("Received shutdown signal %s", s)
			Stop()
		}()

		log.Infof("Starting HTTP Listener on %s, service path is %s", cfg.GetString("listen_addr"), cfg.GetString("service_path"))
		err = srv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				<-runCtx.Done()
				return ErrShutdown
			}
			return fmt.Errorf("unable to start HTTP Server - %s", err)
		}

		return nil
	}
}

// Stop is used to gracefully shutdown the server with a timeout.
func Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	runCancel() // Hủy runCtx trước để báo hiệu cho các goroutine dừng
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Failed to shutdown HTTP server gracefully: %s", err)
		if err := srv.Close(); err != nil {
			log.Errorf("Failed to force close HTTP server: %s", err)
		}
	} else {
		log.Info("HTTP server shutdown gracefully")
	}
	// Chờ một khoảng thời gian ngắn để các goroutine kết thúc
	time.Sleep(100 * time.Millisecond)
}

var isPProf = regexp.MustCompile(`.*debug\/pprof.*`)

// Lấy IP đáng tin cậy từ request
func GetTrustedClientIP(c *gin.Context) string {
	// Lấy IP từ X-Forwarded-For
	xff := c.Request.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		for i := len(ips) - 1; i >= 0; i-- {
			ip := strings.TrimSpace(ips[i])
			if isTrustedProxy(ip) {
				continue
			}
			fmt.Println("IP X-Forwarded-For : ", ip)
			return ip
		}
	}
	// Fallback về RemoteAddr nếu không có proxy đáng tin cậy
	ip, _, _ := net.SplitHostPort(c.Request.RemoteAddr)
	fmt.Println("IP SplitHostPort : ", ip)
	return ip
}

// isTrustedProxy kiểm tra xem IP có thuộc danh sách proxy đáng tin cậy không
func isTrustedProxy(ip string) bool {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return false
	}
	for _, cidr := range TrustedProxies {
		_, subnet, err := net.ParseCIDR(cidr)
		if err != nil {
			continue
		}
		if subnet.Contains(parsedIP) {
			return true
		}
	}
	return false
}

func handleWrapper(n gin.HandlerFunc, tokenRequired bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()

		// Lọc header nhạy cảm cho logging
		safeHeaders := make(http.Header)
		for k, v := range c.Request.Header {
			if k != "Authorization" && k != "Cookie" {
				safeHeaders[k] = v
			}
		}

		// Log request với header đã lọc
		log.WithFields(logrus.Fields{
			"method":         c.Request.Method,
			"remote-addr":    c.Request.RemoteAddr,
			"http-protocol":  c.Request.Proto,
			"headers":        safeHeaders,
			"content-length": c.Request.ContentLength,
		}).Debugf("HTTP Request to %s", c.Request.URL)

		// Kiểm tra PProf
		if isPProf.MatchString(c.Request.URL.Path) && !cfg.GetBool("enable_pprof") {
			log.WithFields(logrus.Fields{
				"method":         c.Request.Method,
				"remote-addr":    c.Request.RemoteAddr,
				"http-protocol":  c.Request.Proto,
				"headers":        safeHeaders,
				"content-length": c.Request.ContentLength,
			}).Debugf("Request to PProf Address failed, PProf disabled")
			c.AbortWithStatus(http.StatusForbidden)
			stats.Srv.WithLabelValues(c.Request.URL.Path).Observe(time.Since(now).Seconds())
			return
		}

		// Kiểm tra token nếu endpoint yêu cầu xác thực
		if tokenRequired {
			// Lấy token từ header Authorization
			authHeader := c.GetHeader("Authorization")
			log.WithFields(logrus.Fields{
				"auth_header": authHeader,
				"path":        c.Request.URL.Path,
			}).Debug("Checking authorization header")

			if authHeader == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
				return
			}

			// Kiểm tra format của token
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header format"})
				return
			}

			tokenString := parts[1]
			log.WithFields(logrus.Fields{
				"token": tokenString[:10] + "...", // Log một phần của token để debug
			}).Debug("Validating token")

			// Xác thực JWT token
			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(cfg.GetString("jwt_secret")), nil
			})

			if err != nil {
				log.WithError(err).Error("Token validation failed")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token: " + err.Error()})
				return
			}

			if !token.Valid {
				log.Error("Token is invalid")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
				return
			}

			// Lấy thông tin user từ token
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				c.Set("user_id", claims["user_id"])
				c.Set("username", claims["username"])
				c.Set("role", claims["role"])
				log.WithFields(logrus.Fields{
					"user_id":  claims["user_id"],
					"username": claims["username"],
				}).Debug("User authenticated successfully")
			}
		}

		// Thực thi handler
		n(c)
		stats.Srv.WithLabelValues(c.Request.URL.Path).Observe(time.Since(now).Seconds())
	}
}
