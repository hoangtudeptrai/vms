package model

import (
	"time"

	"github.com/google/uuid"
)

type AssignmentDocument struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	AssignmentID uuid.UUID `json:"assignment_id" gorm:"type:uuid;not null"`
	Title        string    `json:"title" gorm:"not null"`
	Description  *string   `json:"description"`
	FileName     string    `json:"file_name" gorm:"not null"`
	FilePath     string    `json:"file_path" gorm:"not null"`
	FileSize     *int64    `json:"file_size"`
	FileType     *string   `json:"file_type"`
	UploadedBy   uuid.UUID `json:"uploaded_by" gorm:"type:uuid;not null"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime:true"`
}
type DTOAssignmentDocument struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	AssignmentID uuid.UUID `json:"assignment_id" gorm:"type:uuid;not null"`
	Title        string    `json:"title" gorm:"not null"`
	Description  *string   `json:"description"`
	FileName     string    `json:"file_name" gorm:"not null"`
	FilePath     string    `json:"file_path" gorm:"not null"`
	FileSize     *int64    `json:"file_size"`
	FileType     *string   `json:"file_type"`
	UploadedBy   uuid.UUID `json:"uploaded_by" gorm:"type:uuid;not null"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime:true"`
}
