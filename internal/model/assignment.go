package model

import (
	"time"

	"github.com/google/uuid"
)

type Assignment struct {
	ID               uuid.UUID  `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CourseID         uuid.UUID  `json:"course_id" gorm:"type:uuid;not null"`
	Title            string     `json:"title" gorm:"not null"`
	Description      *string    `json:"description"`
	DueDate          *time.Time `json:"due_date"`
	CreatedBy        uuid.UUID  `json:"created_by" gorm:"type:uuid;not null"`
	Status           string     `json:"status"` //gorm:"default:'draft';type:text;check:status IN ('draft','active','completed')
	CreatedAt        time.Time  `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt        time.Time  `json:"updated_at" gorm:"not null;default:now()"`
	MaxScore         int        `json:"max_score" gorm:"default:100"`
	AssignmentStatus string     `json:"assignment_status"` // gorm:"default:'published';type:assignment_status"
}
