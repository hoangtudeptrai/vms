package model

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	SenderID   uuid.UUID `json:"sender_id" gorm:"type:uuid"`
	ReceiverID uuid.UUID `json:"receiver_id" gorm:"type:uuid"`
	Subject    *string   `json:"subject"`
	Content    string    `json:"content" gorm:"not null"`
	IsRead     bool      `json:"is_read" gorm:"default:false"`
	RepliedTo  uuid.UUID `json:"replied_to" gorm:"type:uuid"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:now()"`
}
