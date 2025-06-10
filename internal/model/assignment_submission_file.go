package model

import (
	"time"

	"github.com/google/uuid"
)

type AssignmentSubmissionFile struct {
	ID           uuid.UUID  `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	SubmissionID *uuid.UUID `json:"submission_id" gorm:"type:uuid"`
	FileName     string     `json:"file_name" gorm:"not null"`
	FilePath     string     `json:"file_path" gorm:"not null"`
	FileSize     *int64     `json:"file_size"`
	FileType     *string    `json:"file_type"`
	UploadedAt   time.Time  `json:"uploaded_at" gorm:"default:now()"`
}
