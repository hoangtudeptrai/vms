package model

import (
	"time"

	"github.com/google/uuid"
)

// Course now references a Document for its cover image.
// Course (KhóaHọc) giờ đây tham chiếu đến một Document cho ảnh bìa.
type Course struct {
	ID                   uuid.UUID  `json:"id" db:"course_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title                string     `json:"title" db:"title"`                                               // Tiêu đề khóa học
	Description          *string    `json:"description,omitempty" db:"description"`                         // Mô tả chi tiết (có thể null)
	TeacherID            uuid.UUID  `json:"teacher_id" db:"teacher_id"`                                     // ID của giáo viên (tham chiếu đến User)
	CoverImageDocumentID *uuid.UUID `json:"cover_image_document_id,omitempty" db:"cover_image_document_id"` // ID ảnh bìa (tham chiếu đến Document, có thể null)
	CreatedAt            time.Time  `json:"created_at" db:"created_at"`                                     // Thời gian tạo khóa học
	UpdatedAt            time.Time  `json:"updated_at" db:"updated_at"`                                     // Thời gian cập nhật gần nhất
}

// CreateCourse represents the DTO for creating a new course
type CreateCourse struct {
	Title                string     `json:"title" binding:"required"`
	Description          *string    `json:"description"`
	TeacherID            uuid.UUID  `json:"teacher_id" binding:"required"`
	CoverImageDocumentID *uuid.UUID `json:"cover_image_document_id"`
}

// UpdateCourse represents the DTO for updating an existing course
type UpdateCourse struct {
	Title                string     `json:"title" binding:"required"`
	Description          *string    `json:"description"`
	TeacherID            uuid.UUID  `json:"teacher_id" binding:"required"`
	CoverImageDocumentID *uuid.UUID `json:"cover_image_document_id"`
}
