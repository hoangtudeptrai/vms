package model

import (
	"time"

	"github.com/google/uuid"
)

// Assignment now references a Document for its attachment.
// Assignment (BàiTập) giờ đây tham chiếu đến một Document cho file đính kèm.
type Assignment struct {
	ID                   uuid.UUID  `json:"id" db:"assignment_id" gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	CourseID             uuid.UUID  `json:"course_id" db:"course_id"`
	CreatedByID          uuid.UUID  `json:"created_by_id" db:"created_by_id"`
	Title                string     `json:"title" db:"title"`
	Description          *string    `json:"description,omitempty" db:"description"`
	DueDate              time.Time  `json:"due_date" db:"due_date"`
	AttachmentDocumentID *uuid.UUID `json:"attachment_document_id,omitempty" db:"attachment_document_id"`
	CreatedAt            time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt            time.Time  `json:"updated_at" db:"updated_at"`
}

// CreateAssignment represents the DTO for creating a new assignment
type CreateAssignment struct {
	CourseID             uuid.UUID  `json:"course_id" binding:"required"`
	CreatedByID          uuid.UUID  `json:"created_by_id" binding:"required"`
	Title                string     `json:"title" binding:"required"`
	Description          *string    `json:"description"`
	DueDate              time.Time  `json:"due_date" binding:"required"`
	AttachmentDocumentID *uuid.UUID `json:"attachment_document_id"`
}

// UpdateAssignment represents the DTO for updating an existing assignment
type UpdateAssignment struct {
	CourseID             uuid.UUID  `json:"course_id" binding:"required"`
	Title                string     `json:"title" binding:"required"`
	Description          *string    `json:"description"`
	DueDate              time.Time  `json:"due_date" binding:"required"`
	AttachmentDocumentID *uuid.UUID `json:"attachment_document_id"`
}
