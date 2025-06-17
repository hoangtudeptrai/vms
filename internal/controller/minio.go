package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// MinIO client configuration
var minioClient *minio.Client

// InitMinIO initializes the MinIO client
func InitMinIO() error {
	endpoint := "127.0.0.1:9000"
	accessKey := "minioadmin"
	secretKey := "minioadmin"
	useSSL := false

	log.Printf("Đang khởi tạo MinIO client với endpoint: %s, accessKey: %s", endpoint, accessKey)

	var err error
	minioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Printf("Lỗi khởi tạo MinIO client: %v", err)
		return fmt.Errorf("failed to initialize MinIO client: %v", err)
	}

	// Kiểm tra kết nối tới MinIO server
	ctx := context.Background()
	buckets, err := minioClient.ListBuckets(ctx)
	if err != nil {
		log.Printf("Lỗi kết nối tới MinIO server: %v", err)
		minioClient = nil // Đặt lại client để tránh sử dụng instance không hợp lệ
		return fmt.Errorf("failed to connect to MinIO server: %v", err)
	}
	log.Printf("Kết nối thành công tới MinIO server, tìm thấy %d bucket", len(buckets))

	// Kiểm tra bucket tồn tại, tạo nếu chưa có
	bucketName := "lms"
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		log.Printf("Lỗi kiểm tra bucket: %v", err)
		return fmt.Errorf("failed to check bucket existence: %v", err)
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Printf("Lỗi tạo bucket: %v", err)
			return fmt.Errorf("failed to create bucket: %v", err)
		}
		log.Printf("Đã tạo bucket %s", bucketName)
	}

	return nil
}

// UploadFile godoc
// @Summary      Tải file lên MinIO
// @Description  Tải file lên bucket 'lms' của MinIO. Định dạng hỗ trợ: JPG, PNG, MP4, PDF, DOC, DOCX, XLS, XLSX.
// @Tags         minio
// @Accept       multipart/form-data
// @Produce      json
// @Param        file formData file true "File cần tải lên"
// @Param        id_user  query  string  true  "Read user by id"
// @Success      200 {object} map[string]interface{} "Tải file thành công"
// @Failure      400 {object} map[string]string "File hoặc yêu cầu không hợp lệ"
// @Failure      500 {object} map[string]string "Lỗi server"
// @Router       /upload [post]
func UploadFile(c *gin.Context) {
	// Kiểm tra MinIO client đã được khởi tạo
	if minioClient == nil {
		log.Println("Lỗi: MinIO client chưa được khởi tạo")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "MinIO client chưa được khởi tạo"})
		return
	}

	// Lấy file từ form-data
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("Lỗi lấy file: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Không có file được tải lên"})
		return
	}

	// Kiểm tra định dạng file
	allowedExtensions := map[string]string{
		".jpg":  "image/jpeg",
		".jpeg": "image/jpeg",
		".png":  "image/png",
		".mp4":  "video/mp4",
		".pdf":  "application/pdf",
		".doc":  "application/msword",
		".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
		".xls":  "application/vnd.ms-excel",
		".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	}
	ext := strings.ToLower(filepath.Ext(file.Filename))
	contentType, valid := allowedExtensions[ext]
	if !valid {
		log.Printf("Định dạng file không hỗ trợ: %s", ext)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Định dạng file không hỗ trợ. Sử dụng JPG, PNG, MP4, PDF, DOC, DOCX, XLS, hoặc XLSX"})
		return
	}

	// Mở file
	f, err := file.Open()
	if err != nil {
		log.Printf("Lỗi mở file: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi mở file"})
		return
	}
	defer f.Close()

	// Tạo tên object dựa trên loại file
	var folder string
	switch ext {
	case ".mp4":
		folder = "videos"
	case ".pdf", ".doc", ".docx":
		folder = "documents"
	case ".xls", ".xlsx":
		folder = "spreadsheets"
	default:
		folder = "images"
	}
	objectName := fmt.Sprintf("%s/%s", folder, file.Filename)

	// Tải lên MinIO
	bucketName := "lms"
	ctx := context.Background()
	info, err := minioClient.PutObject(ctx, bucketName, objectName, f, file.Size, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Printf("Lỗi tải file lên MinIO: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Lỗi tải file: %v", err)})
		return
	}

	// Trả về phản hồi thành công
	log.Printf("Tải file %s thành công, kích thước: %d bytes", objectName, info.Size)
	c.JSON(http.StatusOK, gin.H{"message": "Tải file thành công",
		"file_name": objectName,
		"file_path": fmt.Sprintf("http://%s/%s/%s", minioClient.EndpointURL().Host, bucketName, objectName),
		"file_size": info.Size,
		"file_type": contentType,
		"file_id":   info.Key,
	})
}

// GetFile godoc
// @Summary      Lấy tệp từ MinIO
// @Description  Lấy tệp từ bucket 'lms' của MinIO hoặc trả về URL đã ký trước để truy cập tệp. Yêu cầu tên object của tệp.
// @Tags         minio
// @Accept       json
// @Produce      json
// @Param        objectName query string true "Tên object của tệp (bao gồm thư mục, ví dụ: images/example.jpg)"
// @Success      200 {object} map[string]interface{} "Truy xuất tệp thành công"
// @Failure      400 {object} map[string]string "Yêu cầu không hợp lệ"
// @Failure      404 {object} map[string]string "Tệp không tồn tại"
// @Failure      500 {object} map[string]string "Lỗi server"
// @Router       /file [get]
func GetFile(c *gin.Context) {
	// Kiểm tra MinIO client đã được khởi tạo
	if minioClient == nil {
		log.Println("Lỗi: MinIO client chưa được khởi tạo")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "MinIO client chưa được khởi tạo"})
		return
	}

	// Lấy tên object từ query parameter
	objectName := c.Query("objectName")
	if objectName == "" {
		log.Println("Lỗi: Thiếu tên object")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Yêu cầu tên object của tệp"})
		return
	}

	// Kiểm tra tệp tồn tại trong bucket
	bucketName := "lms"
	ctx := context.Background()
	_, err := minioClient.StatObject(ctx, bucketName, objectName, minio.StatObjectOptions{})
	if err != nil {
		log.Printf("Lỗi kiểm tra tệp %s: %v", objectName, err)
		// Kiểm tra nếu tệp không tồn tại
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tệp không tồn tại"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Lỗi kiểm tra tệp: %v", err)})
		}
		return
	}

	// Tạo URL đã ký trước (presigned URL) để truy cập tệp
	presignedURL, err := minioClient.PresignedGetObject(ctx, bucketName, objectName, time.Hour*24, nil)
	if err != nil {
		log.Printf("Lỗi tạo URL đã ký cho %s: %v", objectName, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi tạo URL truy cập tệp"})
		return
	}

	// Trả về phản hồi thành công
	log.Printf("Truy xuất tệp %s thành công", objectName)
	c.JSON(http.StatusOK, gin.H{
		"message":    "Truy xuất tệp thành công",
		"objectName": objectName,
		"url":        presignedURL.String(),
	})
}
