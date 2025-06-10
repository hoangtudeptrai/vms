package model

import (
	"time"

	"github.com/google/uuid"
)

type AssignmentSubmission struct {
	ID           uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	AssignmentID uuid.UUID `json:"assignment_id" gorm:"type:uuid;not null"`
	StudentID    uuid.UUID `json:"student_id" gorm:"type:uuid;not null"`
	SubmittedAt  time.Time `json:"submitted_at" gorm:"default:now()"`
	Grade        *float64  `json:"grade"` //gorm:"check:grade >= 0 AND grade <= 10"
	Feedback     *string   `json:"feedback"`
	Content      *string   `json:"content"`
	Status       string    `json:"status"` // gorm:"default:'pending';type:submission_status"
}
