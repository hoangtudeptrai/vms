package model

import (
	"time"

	"github.com/google/uuid"
)

// Document represents a single file stored in the system.
// Document (TàiLiệuLưuTrữ) đại diện cho một file duy nhất được lưu trữ trong toàn bộ hệ thống.
// Đây là bảng trung tâm để quản lý tất cả các file.
type Document struct {
	ID               uuid.UUID `json:"id" db:"document_id"`                        // ID duy nhất của file
	UploaderID       uuid.UUID `json:"uploader_id" db:"uploader_id"`               // ID người đã tải file lên (tham chiếu đến User)
	OriginalFileName string    `json:"original_file_name" db:"original_file_name"` // Tên file gốc
	FilePath         string    `json:"file_path" db:"file_path"`                   // Đường dẫn tới file (URL)
	MimeType         string    `json:"mime_type" db:"mime_type"`                   // Kiểu MIME của file
	FileSize         int64     `json:"file_size" db:"file_size"`                   // Kích thước file (bytes)
	CreatedAt        time.Time `json:"created_at" db:"created_at"`                 // Thời gian tải lên
}

// CreateDocument represents the DTO for creating a new document
type CreateDocument struct {
	UploaderID       uuid.UUID `json:"uploader_id" binding:"required"`
	OriginalFileName string    `json:"original_file_name" binding:"required"`
	FilePath         string    `json:"file_path" binding:"required"`
	MimeType         string    `json:"mime_type" binding:"required"`
	FileSize         int64     `json:"file_size" binding:"required"`
}

// UpdateDocument represents the DTO for updating an existing document
type UpdateDocument struct {
	OriginalFileName string `json:"original_file_name" binding:"required"`
	FilePath         string `json:"file_path" binding:"required"`
	MimeType         string `json:"mime_type" binding:"required"`
	FileSize         int64  `json:"file_size" binding:"required"`
}
