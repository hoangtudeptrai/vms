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
				&model.Enrollment{},
				&model.Assignment{},
				&model.Document{},
				&model.Notification{},
				&model.Submission{},
				&model.CourseMaterial{},
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
		apiV0.GET("/users/:id", handleWrapper(controllers.GetUserByID, true))
		apiV0.PUT("/users/:id", handleWrapper(controllers.UpdateUser, true))
		apiV0.DELETE("/users/:id", handleWrapper(controllers.DeleteUser, true))
		apiV0.GET("/users/me", handleWrapper(controllers.GetCurrentUser, false))

		// Course routes
		apiV0.POST("/courses", handleWrapper(controllers.CreateCourse, true))
		apiV0.GET("/courses", handleWrapper(controllers.GetCourses, true))
		apiV0.GET("/courses/:id", handleWrapper(controllers.GetCourseByID, true))
		apiV0.PUT("/courses/:id", handleWrapper(controllers.UpdateCourse, true))
		apiV0.DELETE("/courses/:id", handleWrapper(controllers.DeleteCourse, true))

		// Enrollment routes
		apiV0.POST("/enrollments", handleWrapper(controllers.CreateEnrollment, true))
		apiV0.GET("/enrollments", handleWrapper(controllers.GetEnrollments, true))
		apiV0.GET("/enrollments/:id", handleWrapper(controllers.GetEnrollmentByID, true))
		apiV0.PUT("/enrollments/:id", handleWrapper(controllers.UpdateEnrollment, true))
		apiV0.DELETE("/enrollments/:id", handleWrapper(controllers.DeleteEnrollment, true))

		// Assignment routes
		apiV0.POST("/assignments", handleWrapper(controllers.CreateAssignment, true))
		apiV0.GET("/assignments", handleWrapper(controllers.GetAssignments, true))
		apiV0.GET("/assignments/:id", handleWrapper(controllers.GetAssignmentByID, true))
		apiV0.PUT("/assignments/:id", handleWrapper(controllers.UpdateAssignment, true))
		apiV0.DELETE("/assignments/:id", handleWrapper(controllers.DeleteAssignment, true))

		// Document routes
		apiV0.POST("/documents", handleWrapper(controllers.CreateDocument, true))
		apiV0.GET("/documents", handleWrapper(controllers.GetDocuments, false))
		apiV0.GET("/documents/:id", handleWrapper(controllers.GetDocumentByID, true))
		apiV0.PUT("/documents/:id", handleWrapper(controllers.UpdateDocument, true))
		apiV0.DELETE("/documents/:id", handleWrapper(controllers.DeleteDocument, true))

		// Notification routes
		apiV0.POST("/notifications", handleWrapper(controllers.CreateNotification, true))
		apiV0.GET("/notifications", handleWrapper(controllers.GetNotifications, true))
		apiV0.GET("/notifications/:id", handleWrapper(controllers.GetNotificationByID, true))
		apiV0.PUT("/notifications/:id", handleWrapper(controllers.UpdateNotification, true))
		apiV0.DELETE("/notifications/:id", handleWrapper(controllers.DeleteNotification, true))

		// Submission routes
		apiV0.POST("/submissions", handleWrapper(controllers.CreateSubmission, true))
		apiV0.GET("/submissions", handleWrapper(controllers.GetSubmissions, true))
		apiV0.GET("/submissions/:id", handleWrapper(controllers.GetSubmissionByID, true))
		apiV0.PUT("/submissions/:id", handleWrapper(controllers.UpdateSubmission, true))
		apiV0.DELETE("/submissions/:id", handleWrapper(controllers.DeleteSubmission, true))

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
