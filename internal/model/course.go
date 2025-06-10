package model

import (
	"time"

	"github.com/google/uuid"
)

type Course struct {
	ID            uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title         string    `json:"title" gorm:"not null"`
	Description   *string   `json:"description"`
	InstructorID  uuid.UUID `json:"instructor_id" gorm:"type:uuid;not null"`
	Thumbnail     *string   `json:"thumbnail"`
	Duration      *string   `json:"duration"`
	LessonsCount  int       `json:"lessons_count" gorm:"default:0"`
	StudentsCount int       `json:"students_count" gorm:"default:0"`
	Status        string    `json:"status" gorm:"default:'draft';type:text;check:status IN ('draft','active','completed','archived')"`
	CreatedAt     time.Time `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"not null;default:now()"`
}
