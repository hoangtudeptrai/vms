package model

import (
	"time"

	"github.com/google/uuid"
)

// User is the base model for users
type User struct {
	UserID     uuid.UUID `json:"UserID" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Username   string    `json:"Username" gorm:"unique;not null"`
	Password   string    `json:"Password" gorm:"not null"`
	Email      string    `json:"Email" gorm:"unique;not null"`
	FullName   string    `json:"FullName" gorm:"not null"`
	Role       string    `json:"Role" gorm:"not null"`
	AvatarURL  string    `json:"AvatarURL"`
	TelegramID string    `json:"TelegramID"`
	CreatedAt  time.Time `json:"CreatedAt"`
	UpdatedAt  time.Time `json:"UpdatedAt"`
}

// CreateUser DTO for creating/updating a User
type CreateUser struct {
	UserID     uuid.UUID `json:"UserID"`
	Username   string    `json:"Username" binding:"required"`
	Password   string    `json:"Password" binding:"required"`
	Email      string    `json:"Email" binding:"required,email"`
	FullName   string    `json:"FullName" binding:"required"`
	Role       string    `json:"Role" binding:"required,oneof=student teacher admin"`
	AvatarURL  string    `json:"AvatarURL"`
	TelegramID string    `json:"TelegramID"`
}

// UpdateUser DTO for reading a User
type UpdateUser struct {
	UserID     uuid.UUID `json:"UserID"`
	Username   string    `json:"Username"`
	Email      string    `json:"Email"`
	FullName   string    `json:"FullName"`
	Role       string    `json:"Role"`
	AvatarURL  string    `json:"AvatarURL"`
	TelegramID string    `json:"TelegramID"`
	CreatedAt  time.Time `json:"CreatedAt"`
	UpdatedAt  time.Time `json:"UpdatedAt"`
}

// CreateCourse DTO for creating/updating a Course
type CreateCourse struct {
	Title       string `json:"Title" binding:"required"`
	Description string `json:"Description"`
	TeacherID   uint   `json:"TeacherID" binding:"required"`
}

// UpdateCourse DTO for reading a Course
type UpdateCourse struct {
	CourseID    uint      `json:"CourseID"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	TeacherID   uint      `json:"TeacherID"`
	CreatedAt   time.Time `json:"CreatedAt"`
	UpdatedAt   time.Time `json:"UpdatedAt"`
}

// CreateEnrollment DTO for creating/updating an Enrollment
type CreateEnrollment struct {
	UserID   uint `json:"UserID" binding:"required"`
	CourseID uint `json:"CourseID" binding:"required"`
}

// UpdateEnrollment DTO for reading an Enrollment
type UpdateEnrollment struct {
	EnrollmentID uint      `json:"EnrollmentID"`
	UserID       uint      `json:"UserID"`
	CourseID     uint      `json:"CourseID"`
	EnrolledAt   time.Time `json:"EnrolledAt"`
}

// CreateDocument DTO for creating/updating a Document
type CreateDocument struct {
	CourseID   uint   `json:"CourseID" binding:"required"`
	Title      string `json:"Title" binding:"required"`
	FileURL    string `json:"FileURL" binding:"required"`
	FileType   string `json:"FileType" binding:"required,oneof=pdf video image"`
	UploadedBy uint   `json:"UploadedBy" binding:"required"`
}

// UpdateDocument DTO for reading a Document
type UpdateDocument struct {
	DocumentID uint      `json:"DocumentID"`
	CourseID   uint      `json:"CourseID"`
	Title      string    `json:"Title"`
	FileURL    string    `json:"FileURL"`
	FileType   string    `json:"FileType"`
	UploadedBy uint      `json:"UploadedBy"`
	UploadedAt time.Time `json:"UploadedAt"`
}

// CreateAssignment DTO for creating/updating an Assignment
type CreateAssignment struct {
	CourseID    uint      `json:"CourseID" binding:"required"`
	Title       string    `json:"Title" binding:"required"`
	Description string    `json:"Description"`
	DueDate     time.Time `json:"DueDate" binding:"required"`
	CreatedBy   uint      `json:"CreatedBy" binding:"required"`
}

// UpdateAssignment DTO for reading an Assignment
type UpdateAssignment struct {
	AssignmentID uint      `json:"AssignmentID"`
	CourseID     uint      `json:"CourseID"`
	Title        string    `json:"Title"`
	Description  string    `json:"Description"`
	DueDate      time.Time `json:"DueDate"`
	CreatedBy    uint      `json:"CreatedBy"`
	CreatedAt    time.Time `json:"CreatedAt"`
}

// CreateSubmission DTO for creating/updating a Submission
type CreateSubmission struct {
	AssignmentID uint   `json:"AssignmentID" binding:"required"`
	UserID       uint   `json:"UserID" binding:"required"`
	FileURL      string `json:"FileURL"`
	TextContent  string `json:"TextContent"`
}

// UpdateSubmission DTO for reading a Submission
type UpdateSubmission struct {
	SubmissionID uint      `json:"SubmissionID"`
	AssignmentID uint      `json:"AssignmentID"`
	UserID       uint      `json:"UserID"`
	FileURL      string    `json:"FileURL"`
	TextContent  string    `json:"TextContent"`
	SubmittedAt  time.Time `json:"SubmittedAt"`
}

// CreateGrade DTO for creating/updating a Grade
type CreateGrade struct {
	SubmissionID uint    `json:"SubmissionID" binding:"required"`
	Score        float64 `json:"Score" binding:"required,gte=0,lte=100"`
	Feedback     string  `json:"Feedback"`
	GradedBy     uint    `json:"GradedBy" binding:"required"`
}

// UpdateGrade DTO for reading a Grade
type UpdateGrade struct {
	GradeID      uint      `json:"GradeID"`
	SubmissionID uint      `json:"SubmissionID"`
	Score        float64   `json:"Score"`
	Feedback     string    `json:"Feedback"`
	GradedBy     uint      `json:"GradedBy"`
	GradedAt     time.Time `json:"GradedAt"`
}

// CreateNotification DTO for creating/updating a Notification
type CreateNotification struct {
	UserID  uint   `json:"UserID" binding:"required"`
	Content string `json:"Content" binding:"required"`
	Type    string `json:"Type" binding:"required,oneof=assignment grade course system"`
	IsRead  bool   `json:"IsRead"`
}

// UpdateNotification DTO for reading a Notification
type UpdateNotification struct {
	NotificationID uint      `json:"NotificationID"`
	UserID         uint      `json:"UserID"`
	Content        string    `json:"Content"`
	Type           string    `json:"Type"`
	IsRead         bool      `json:"IsRead"`
	CreatedAt      time.Time `json:"CreatedAt"`
}
