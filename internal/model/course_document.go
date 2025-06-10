package model

import (
	"time"

	"github.com/google/uuid"
)

type CourseDocument struct {
	ID          uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CourseID    uuid.UUID `json:"course_id" gorm:"type:uuid;not null"`
	Title       string    `json:"title" gorm:"not null"`
	Description *string   `json:"description"`
	FileName    string    `json:"file_name" gorm:"not null"`
	FilePath    string    `json:"file_path" gorm:"not null"`
	FileSize    *int64    `json:"file_size"`
	FileType    *string   `json:"file_type"`
	UploadedBy  uuid.UUID `json:"uploaded_by" gorm:"type:uuid;not null"`
	CreatedAt   time.Time `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"not null;default:now()"`
}
