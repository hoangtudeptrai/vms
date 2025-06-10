package model

import (
	"time"

	"github.com/google/uuid"
)

type CourseEnrollment struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CourseID   uuid.UUID `json:"course_id" gorm:"type:uuid;not null"`
	StudentID  uuid.UUID `json:"student_id" gorm:"type:uuid;not null"`
	EnrolledAt time.Time `json:"enrolled_at" gorm:"not null;default:now()"`
	Progress   int       `json:"progress" gorm:"default:0;check:progress >= 0 AND progress <= 100"`
	Status     string    `json:"status" gorm:"default:'enrolled';type:text;check:status IN ('enrolled','completed','dropped')"`
	LastActive time.Time `json:"last_active" gorm:"default:now()"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime:true"`
}
