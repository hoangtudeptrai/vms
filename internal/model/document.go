package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Document struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title       string    `gorm:"type:varchar(255);not null" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	Content     string    `gorm:"type:text" json:"content"`
	FileURL     string    `gorm:"type:text" json:"file_url"`
	Type        string    `gorm:"type:varchar(50);not null" json:"type"` // lecture, assignment, material
	CourseID    uuid.UUID `gorm:"type:uuid;not null" json:"course_id"`
	LessonID    uuid.UUID `gorm:"type:uuid" json:"lesson_id"`
	CreatedBy   uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt
}
