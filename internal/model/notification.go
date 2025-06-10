package model

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:uuid"`
	Title     string    `json:"title" gorm:"not null"`
	Content   string    `json:"content" gorm:"not null"`
	Type      string    `json:"type" gorm:"default:'system'"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	RelatedID uuid.UUID `json:"related_id" gorm:"type:uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
}
