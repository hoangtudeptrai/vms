package model

import (
	"time"

	"github.com/google/uuid"
)

// Assignment DTOs
type CreateAssignment struct {
	CourseID    uuid.UUID  `json:"course_id" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Description *string    `json:"description"`
	DueDate     *time.Time `json:"due_date"`
	MaxScore    int        `json:"max_score"`
}

type UpdateAssignment struct {
	Title            string     `json:"title"`
	Description      *string    `json:"description"`
	DueDate          *time.Time `json:"due_date"`
	Status           string     `json:"status"`
	MaxScore         int        `json:"max_score"`
	AssignmentStatus string     `json:"assignment_status"`
}

// Course DTOs
type CreateCourse struct {
	Title        string    `json:"title" binding:"required"`
	Description  *string   `json:"description"`
	InstructorID uuid.UUID `json:"instructor_id" binding:"required"`
	Thumbnail    *string   `json:"thumbnail"`
	Duration     *string   `json:"duration"`
}

type UpdateCourse struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Thumbnail   *string `json:"thumbnail"`
	Duration    *string `json:"duration"`
	Status      string  `json:"status"`
}

// Lesson DTOs
type CreateLesson struct {
	CourseID   uuid.UUID `json:"course_id" binding:"required"`
	Title      string    `json:"title" binding:"required"`
	Content    *string   `json:"content"`
	Duration   *int      `json:"duration"`
	OrderIndex int       `json:"order_index" binding:"required"`
	Type       string    `json:"type"`
}

type UpdateLesson struct {
	Title      string  `json:"title"`
	Content    *string `json:"content"`
	Duration   *int    `json:"duration"`
	OrderIndex int     `json:"order_index"`
	Type       string  `json:"type"`
}

// Profile DTOs
type CreateProfile struct {
	FullName        string   `json:"full_name" binding:"required"`
	Email           string   `json:"email" binding:"required,email"`
	Role            string   `json:"role" binding:"required"`
	Bio             *string  `json:"bio"`
	PhoneNumber     *string  `json:"phone_number"`
	Address         *string  `json:"address"`
	Education       *string  `json:"education"`
	Experience      *string  `json:"experience"`
	Specializations []string `json:"specializations"`
}

type UpdateProfile struct {
	FullName        string   `json:"full_name"`
	AvatarURL       *string  `json:"avatar_url"`
	Bio             *string  `json:"bio"`
	PhoneNumber     *string  `json:"phone_number"`
	Address         *string  `json:"address"`
	Education       *string  `json:"education"`
	Experience      *string  `json:"experience"`
	Specializations []string `json:"specializations"`
}

// Message DTOs
type CreateMessage struct {
	ReceiverID uuid.UUID  `json:"receiver_id" binding:"required"`
	Subject    *string    `json:"subject"`
	Content    string     `json:"content" binding:"required"`
	RepliedTo  *uuid.UUID `json:"replied_to"`
}

type UpdateMessage struct {
	IsRead bool `json:"is_read"`
}

// Notification DTOs
type CreateNotification struct {
	UserID    uuid.UUID  `json:"user_id" binding:"required"`
	Title     string     `json:"title" binding:"required"`
	Content   string     `json:"content" binding:"required"`
	Type      string     `json:"type"`
	RelatedID *uuid.UUID `json:"related_id"`
}

// CourseEnrollment DTOs
type CreateCourseEnrollment struct {
	StudentID uuid.UUID `json:"student_id" binding:"required"`
	CourseID  uuid.UUID `json:"course_id" binding:"required"`
	Status    string    `json:"status" binding:"required,oneof=enrolled pending dropped completed"`
}

type UpdateCourseEnrollment struct {
	Status string `json:"status" binding:"required,oneof=enrolled pending dropped completed"`
}

// Grade DTOs
type CreateGrade struct {
	StudentID    uuid.UUID `json:"student_id" binding:"required"`
	AssignmentID uuid.UUID `json:"assignment_id" binding:"required"`
	CourseID     uuid.UUID `json:"course_id" binding:"required"`
	Score        float64   `json:"score" binding:"required,min=0,max=100"`
	Feedback     string    `json:"feedback"`
}

type UpdateGrade struct {
	Score    float64 `json:"score" binding:"required,min=0,max=100"`
	Feedback string  `json:"feedback"`
}

// Submission DTOs
type CreateSubmission struct {
	StudentID    uuid.UUID `json:"student_id" binding:"required"`
	AssignmentID uuid.UUID `json:"assignment_id" binding:"required"`
	Content      string    `json:"content" binding:"required"`
	FileURL      string    `json:"file_url"`
}

type UpdateSubmission struct {
	Content string `json:"content"`
	FileURL string `json:"file_url"`
	Status  string `json:"status" binding:"required,oneof=submitted graded"`
}

// Comment DTOs
type CreateComment struct {
	UserID       uuid.UUID `json:"user_id" binding:"required"`
	SubmissionID uuid.UUID `json:"submission_id" binding:"required"`
	Content      string    `json:"content" binding:"required"`
}

type UpdateComment struct {
	Content string `json:"content" binding:"required"`
}

// AssignmentDocument DTOs
type CreateAssignmentDocument struct {
	AssignmentID uuid.UUID `json:"assignment_id" binding:"required"`
	Title        string    `json:"title" binding:"required"`
	Description  *string   `json:"description"`
	FileName     string    `json:"file_name" binding:"required"`
	FilePath     string    `json:"file_path" binding:"required"`
	FileSize     *int64    `json:"file_size"`
	FileType     *string   `json:"file_type"`
	UploadedBy   string    `json:"uploaded_by" binding:"required"`
}

type UpdateAssignmentDocument struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	FileName    string  `json:"file_name" binding:"required"`
	FilePath    string  `json:"file_path" binding:"required"`
	FileSize    *int64  `json:"file_size"`
	FileType    *string `json:"file_type"`
}

// CourseDocument DTOs
type CreateCourseDocument struct {
	CourseID    uuid.UUID `json:"course_id" binding:"required"`
	Title       string    `json:"title" binding:"required"`
	Description *string   `json:"description"`
	FileName    string    `json:"file_name" binding:"required"`
	FilePath    string    `json:"file_path" binding:"required"`
	FileSize    *int64    `json:"file_size"`
	FileType    *string   `json:"file_type"`
	UploadedBy  string    `json:"uploaded_by" binding:"required"`
}

type UpdateCourseDocument struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description"`
	FileName    string  `json:"file_name" binding:"required"`
	FilePath    string  `json:"file_path" binding:"required"`
	FileSize    *int64  `json:"file_size"`
	FileType    *string `json:"file_type"`
}

// AssignmentSubmissionFile DTOs
type CreateAssignmentSubmissionFile struct {
	SubmissionID uuid.UUID `json:"submission_id"`
	FileName     string    `json:"file_name" binding:"required"`
	FilePath     string    `json:"file_path" binding:"required"`
	FileSize     *int64    `json:"file_size"`
	FileType     *string   `json:"file_type"`
}

type UpdateAssignmentSubmissionFile struct {
	FileName string  `json:"file_name" binding:"required"`
	FilePath string  `json:"file_path" binding:"required"`
	FileSize *int64  `json:"file_size"`
	FileType *string `json:"file_type"`
}

// AssignmentSubmission DTOs
type CreateAssignmentSubmission struct {
	AssignmentID uuid.UUID `json:"assignment_id" binding:"required"`
	StudentID    uuid.UUID `json:"student_id" binding:"required"`
	Content      *string   `json:"content"`
	Grade        *float64  `json:"grade" binding:"omitempty,min=0,max=10"`
	Feedback     *string   `json:"feedback"`
	Status       string    `json:"status" binding:"required,oneof=pending submitted graded"`
}

type UpdateAssignmentSubmission struct {
	Content  *string  `json:"content"`
	Grade    *float64 `json:"grade" binding:"omitempty,min=0,max=10"`
	Feedback *string  `json:"feedback"`
	Status   string   `json:"status" binding:"required,oneof=pending submitted graded"`
}
