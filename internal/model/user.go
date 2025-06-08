package model

import (
	"time"

	"github.com/google/uuid"
)

// User is the base model for users
type User struct {
	ID                uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Email             string    `json:"email" db:"email"`
	FullName          string    `json:"full_name" db:"full_name"`
	UserName          string    `json:"user_name" db:"user_name"`
	Password          string    `json:"password" db:"password"`
	Role              string    `json:"role" db:"role"` // 'student', 'teacher', 'admin'
	ProfilePictureURL string    `json:"profile_picture_url,omitempty" db:"profile_picture_url"`
	PhoneNumber       string    `json:"phone_number,omitempty" db:"phone_number"`
	DateOfBirth       time.Time `json:"date_of_birth,omitempty" db:"date_of_birth"` // Sử dụng time.Time cho DATE có thể NULL
	Address           string    `json:"address,omitempty" db:"address"`
	Bio               string    `json:"bio,omitempty" db:"bio"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}

// CreateUser DTO for creating a new User
type CreateUser struct {
	ID                uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	FullName          string    `json:"full_name" binding:"required"`
	UserName          string    `json:"user_name"`
	Email             string    `json:"email" binding:"required,email"`
	Password          string    `json:"password" binding:"required"` // Changed from password since this is raw password
	Role              string    `json:"role" binding:"required,oneof=student teacher admin"`
	ProfilePictureURL string    `json:"profile_picture_url,omitempty"`
	PhoneNumber       string    `json:"phone_number,omitempty"`
	DateOfBirth       time.Time `json:"date_of_birth,omitempty"`
	Address           string    `json:"address,omitempty"`
	Bio               string    `json:"bio,omitempty"`
}

// UpdateUser DTO for updating an existing User
type UpdateUser struct {
	ID                uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	FullName          string    `json:"full_name"`
	UserName          string    `json:"user_name"`
	Email             string    `json:"email"`
	Role              string    `json:"role"`
	ProfilePictureURL string    `json:"profile_picture_url,omitempty"`
	PhoneNumber       string    `json:"phone_number,omitempty"`
	DateOfBirth       time.Time `json:"date_of_birth,omitempty"`
	Address           string    `json:"address,omitempty"`
	Bio               string    `json:"bio,omitempty"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
