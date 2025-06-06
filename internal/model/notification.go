package model

import (
	"time"

	"github.com/google/uuid"
)

// Notification struct remains unchanged by this refactoring.
// Struct Notification không thay đổi trong lần tái cấu trúc này.
type Notification struct {
	ID               uuid.UUID `json:"id" db:"notification_id"`
	UserID           uuid.UUID `json:"user_id" db:"user_id"`
	Message          string    `json:"message" db:"message"`
	Link             *string   `json:"link,omitempty" db:"link"`
	IsRead           bool      `json:"is_read" db:"is_read"`
	NotificationType string    `json:"notification_type" db:"notification_type"`
	Priority         string    `json:"priority" db:"priority"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}

// CreateNotification represents the DTO for creating a new notification
type CreateNotification struct {
	UserID           uuid.UUID `json:"user_id" binding:"required"`
	Message          string    `json:"message" binding:"required"`
	Link             *string   `json:"link"`
	NotificationType string    `json:"notification_type" binding:"required"`
	Priority         string    `json:"priority" binding:"required"`
}

// UpdateNotification represents the DTO for updating an existing notification
type UpdateNotification struct {
	Message          string  `json:"message" binding:"required"`
	Link             *string `json:"link"`
	IsRead           bool    `json:"is_read"`
	NotificationType string  `json:"notification_type" binding:"required"`
	Priority         string  `json:"priority" binding:"required"`
}
