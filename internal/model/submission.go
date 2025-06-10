package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Submission struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	StudentID    uuid.UUID `gorm:"type:uuid;not null" json:"student_id"`
	AssignmentID uuid.UUID `gorm:"type:uuid;not null" json:"assignment_id"`
	Content      string    `gorm:"type:text;not null" json:"content"`
	FileURL      string    `gorm:"type:text" json:"file_url"`
	Status       string    `gorm:"type:varchar(20);not null;default:'submitted'" json:"status"` // submitted, graded
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime:true"`
	DeletedAt    gorm.DeletedAt
}
