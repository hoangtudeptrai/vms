package model

import (
	"time"

	"github.com/google/uuid"
)

// Submission now references a Document for the submitted file.
// Submission (BàiNộp) giờ đây tham chiếu đến một Document cho file bài làm.
type Submission struct {
	ID                   uuid.UUID  `json:"id" db:"submission_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	AssignmentID         uuid.UUID  `json:"assignment_id" db:"assignment_id"`
	StudentID            uuid.UUID  `json:"student_id" db:"student_id"`
	SubmissionContent    *string    `json:"submission_content,omitempty" db:"submission_content"`
	SubmissionDocumentID *uuid.UUID `json:"submission_document_id,omitempty" db:"submission_document_id"`
	SubmittedAt          time.Time  `json:"submitted_at" db:"submitted_at"`
	Grade                *float64   `json:"grade,omitempty" db:"grade"`
	Feedback             *string    `json:"feedback,omitempty" db:"feedback"`
	GradedByID           *uuid.UUID `json:"graded_by_id,omitempty" db:"graded_by_id"`
	GradedAt             *time.Time `json:"graded_at,omitempty" db:"graded_at"`
}

// CreateSubmission represents the DTO for creating a new submission
type CreateSubmission struct {
	AssignmentID         uuid.UUID  `json:"assignment_id" binding:"required"`
	StudentID            uuid.UUID  `json:"student_id" binding:"required"`
	SubmissionContent    *string    `json:"submission_content"`
	SubmissionDocumentID *uuid.UUID `json:"submission_document_id"`
}

// UpdateSubmission represents the DTO for updating an existing submission
type UpdateSubmission struct {
	SubmissionContent    *string    `json:"submission_content"`
	SubmissionDocumentID *uuid.UUID `json:"submission_document_id"`
	Grade                *float64   `json:"grade"`
	Feedback             *string    `json:"feedback"`
	GradedByID           *uuid.UUID `json:"graded_by_id"`
}
