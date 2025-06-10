package model

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID              uuid.UUID `json:"id" gorm:"primary_key;type:uuid"`
	FullName        string    `json:"full_name" gorm:"not null;default:''"`
	Email           string    `json:"email" gorm:"not null;default:''"`
	AvatarURL       *string   `json:"avatar_url"`
	Role            string    `json:"role"` //gorm:"not null;default:'student';type:user_role"
	Bio             *string   `json:"bio"`
	PhoneNumber     *string   `json:"phone_number"`
	Address         *string   `json:"address"`
	Education       *string   `json:"education"`
	Experience      *string   `json:"experience"`
	Specializations string    `json:"specializations"` // gorm:"type:text[]"
	CreatedAt       time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"default:now()"`
}
