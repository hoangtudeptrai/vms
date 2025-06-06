package model

import (
	"time"

	"github.com/google/uuid"
)

// CourseMaterial represents a Document being used as a learning material in a Course.
// CourseMaterial (TàiLiệuHọcTậpCủaKhóaHọc) liên kết một Document với một Course.
// Bảng này thay thế cho bảng "Materials" cũ, giúp trả lời câu hỏi: "Khóa học này có những tài liệu nào?".
type CourseMaterial struct {
	ID           uuid.UUID `json:"id" db:"course_material_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CourseID     uuid.UUID `json:"course_id" db:"course_id"`
	DocumentID   uuid.UUID `json:"document_id" db:"document_id"`
	Title        string    `json:"title" db:"title"`
	Description  *string   `json:"description,omitempty" db:"description"`
	UploadedByID uuid.UUID `json:"uploaded_by_id" db:"uploaded_by_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
}

// CreateCourseMaterial represents the DTO for creating a new course material
type CreateCourseMaterial struct {
	CourseID     uuid.UUID `json:"course_id" binding:"required"`
	DocumentID   uuid.UUID `json:"document_id" binding:"required"`
	Title        string    `json:"title" binding:"required"`
	Description  *string   `json:"description"`
	UploadedByID uuid.UUID `json:"uploaded_by_id" binding:"required"`
}

// UpdateCourseMaterial represents the DTO for updating an existing course material
type UpdateCourseMaterial struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
}
