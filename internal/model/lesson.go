package model

import (
	"time"

	"github.com/google/uuid"
)

type Lesson struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CourseID   uuid.UUID `json:"course_id" gorm:"type:uuid"`
	Title      string    `json:"title" gorm:"not null"`
	Content    *string   `json:"content"`
	Duration   *int      `json:"duration"`
	OrderIndex int       `json:"order_index" gorm:"not null"`
	Type       string    `json:"type" gorm:"default:'text'"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:now()"`
}
