package model

import (
	"time"

	"github.com/google/uuid"
)

type Grade struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	StudentID    uuid.UUID `json:"student_id" gorm:"type:uuid"`
	CourseID     uuid.UUID `json:"course_id" gorm:"type:uuid"`
	AssignmentID uuid.UUID `json:"assignment_id" gorm:"type:uuid"`
	Score        float64   `json:"score" gorm:"not null"`
	MaxScore     float64   `json:"max_score" gorm:"not null;default:100"`
	Percentage   float64   `json:"percentage"` //gorm:"default:round((score / max_score) * 100, 2)"
	GradedBy     uuid.UUID `json:"graded_by" gorm:"type:uuid"`
	GradedAt     time.Time `json:"graded_at" gorm:"default:now()"`
	CreatedAt    time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt    time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime:true"`
	Comments     *string   `json:"comments"`
}
