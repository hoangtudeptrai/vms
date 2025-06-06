package model

import (
	"time"

	"github.com/google/uuid"
)

// Enrollment struct remains unchanged by this refactoring.
// Struct Enrollment không thay đổi trong lần tái cấu trúc này.
type Enrollment struct {
	ID             uuid.UUID  `json:"id" db:"enrollment_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID         uuid.UUID  `json:"user_id" db:"user_id"`
	CourseID       uuid.UUID  `json:"course_id" db:"course_id"`
	TeacherID      *uuid.UUID `json:"teacher_id,omitempty" db:"teacher_id"`
	EnrollmentDate time.Time  `json:"enrollment_date" db:"enrollment_date"`
}

// CreateEnrollment represents the DTO for creating a new enrollment
type CreateEnrollment struct {
	UserID    uuid.UUID  `json:"user_id" binding:"required"`
	CourseID  uuid.UUID  `json:"course_id" binding:"required"`
	TeacherID *uuid.UUID `json:"teacher_id"`
}

// UpdateEnrollment represents the DTO for updating an existing enrollment
type UpdateEnrollment struct {
	TeacherID *uuid.UUID `json:"teacher_id"`
}
