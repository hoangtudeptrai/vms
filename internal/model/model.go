package model

import (
	"time"

	"github.com/google/uuid"
)

// // User represents the users table
// type User struct {
// 	UserID     uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
// 	Username   string    `json:"Username"`
// 	Password   string    `json:"Password"`
// 	Email      string    `json:"Email"`
// 	FullName   string    `json:"FullName"`
// 	Role       string    `json:"Role"`
// 	AvatarURL  string    `json:"AvatarURL"`
// 	TelegramID string    `json:"TelegramID"`
// 	DeleteMark bool      `json:"deletedMark" gorm:"column:deleted_mark;default:false"`
// 	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime:true"`
// 	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime:true"`
// 	DeletedAt  time.Time `json:"deletedAt" gorm:"column:deleted_at"`
// }

// Course represents the courses table
type Course struct {
	CourseID    uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Title       string    `json:"Title"`
	Description string    `json:"Description"`
	TeacherID   uuid.UUID `json:"TeacherID"`
	Teacher     User      `json:"Teacher"`
	DeleteMark  bool      `json:"deletedMark" gorm:"column:deleted_mark;default:false"`
	CreatedAt   time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime:true"`
	DeletedAt   time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}

// Enrollment represents the enrollments table
type Enrollment struct {
	EnrollmentID uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID       uuid.UUID `json:"UserID"`
	CourseID     uuid.UUID `json:"CourseID"`
	User         User      `json:"User"`
	Course       Course    `json:"Course"`
	EnrolledAt   time.Time `json:"EnrolledAt"`
}

// Document represents the documents table
type Document struct {
	DocumentID uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CourseID   uuid.UUID `json:"CourseID"`
	Title      string    `json:"Title"`
	FileURL    string    `json:"FileURL"`
	FileType   string    `json:"FileType"`
	UploadedBy uuid.UUID `json:"UploadedBy"`
	Uploader   User      `json:"Uploader"`
	UploadedAt time.Time `json:"UploadedAt"`
}

// Assignment represents the assignments table
type Assignment struct {
	AssignmentID uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CourseID     uuid.UUID `json:"CourseID"`
	Title        string    `json:"Title"`
	Description  string    `json:"Description"`
	DueDate      time.Time `json:"DueDate"`
	CreatedBy    uuid.UUID `json:"CreatedBy"`
	Creator      User      `json:"Creator"`
	CreatedAt    time.Time `json:"CreatedAt"`
}

// Submission represents the submissions table
type Submission struct {
	SubmissionID uuid.UUID  `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	AssignmentID uuid.UUID  `json:"AssignmentID"`
	UserID       uuid.UUID  `json:"UserID"`
	FileURL      string     `json:"FileURL"`
	TextContent  string     `json:"TextContent"`
	Assignment   Assignment `json:"Assignment"`
	User         User       `json:"User"`
	SubmittedAt  time.Time  `json:"SubmittedAt"`
}

// Grade represents the grades table
type Grade struct {
	GradeID      uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	SubmissionID uuid.UUID `json:"SubmissionID"`
	Score        float64   `json:"Score"`
	Feedback     string    `json:"Feedback"`
	GradedBy     uuid.UUID `json:"GradedBy"`
	Grader       User      `json:"Grader"`
	GradedAt     time.Time `json:"GradedAt"`
}

// Notification represents the notifications table
type Notification struct {
	NotificationID uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	UserID         uuid.UUID `json:"UserID"`
	Content        string    `json:"Content"`
	Type           string    `json:"Type"`
	IsRead         bool      `json:"IsRead"`
	User           User      `json:"User"`
	DeleteMark     bool      `json:"deletedMark" gorm:"column:deleted_mark;default:false"`
	CreatedAt      time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime:true"`
	DeletedAt      time.Time `json:"deletedAt" gorm:"column:deleted_at"`
}
