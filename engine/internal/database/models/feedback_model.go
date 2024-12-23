package models

import (
	"strconv"
	"time"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model

	// Fields
	Email        string `gorm:"type:varchar(255);not null" json:"email"`
	Description  string `gorm:"type:text;not null" json:"description"`
	FeedbackType string `gorm:"type:enum('bug', 'feature', 'general');not null" json:"feedback_type"`
}

type FeedbackResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Email        string `json:"email"`
	Description  string `json:"description"`
	FeedbackType string `json:"feedbackType"`
}

type FeedbackRequest struct {
	Email        string `json:"email" validate:"required,email,max=255"`
	Description  string `json:"description" validate:"required"`
	FeedbackType string `json:"feedbackType" validate:"required,oneof=bug feature general"`
}

func (m *ModelResource) FeedbackToResource(feedback *Feedback) *FeedbackResource {
	if feedback == nil {
		return nil
	}
	return &FeedbackResource{
		ID:        feedback.ID,
		CreatedAt: feedback.CreatedAt.Format(time.RFC3339),
		UpdatedAt: feedback.UpdatedAt.Format(time.RFC3339),

		Email:        feedback.Email,
		Description:  feedback.Description,
		FeedbackType: feedback.FeedbackType,
	}
}

func (m *ModelResource) FeedbackToResourceList(feedbacks []*Feedback) []*FeedbackResource {
	if feedbacks == nil {
		return nil
	}
	var feedbackResources []*FeedbackResource
	for _, feedback := range feedbacks {
		feedbackResources = append(feedbackResources, m.FeedbackToResource(feedback))
	}
	return feedbackResources
}

// FeedbackToRecord converts a slice of Feedback pointers into CSV records and headers.
func (m *ModelResource) FeedbackToRecord(feedbacks []*Feedback) ([][]string, []string) {
	// Convert Feedback structs to FeedbackResource structs
	resource := m.FeedbackToResourceList(feedbacks)
	records := make([][]string, 0, len(resource))

	for _, feedback := range resource {
		// Basic Fields
		id := strconv.Itoa(int(feedback.ID))
		email := sanitizeCSVField(feedback.Email)
		description := sanitizeCSVField(feedback.Description)
		feedbackType := sanitizeCSVField(feedback.FeedbackType)
		createdAt := sanitizeCSVField(feedback.CreatedAt)
		updatedAt := sanitizeCSVField(feedback.UpdatedAt)

		// Assemble the record
		record := []string{
			id,
			email,
			description,
			feedbackType,
			createdAt,
			updatedAt,
		}
		records = append(records, record)
	}

	headers := []string{
		"ID",
		"Email",
		"Description",
		"Feedback Type",
		"Created At",
		"Updated At",
	}

	return records, headers
}

func (m *ModelResource) ValidateFeedbackRequest(req *FeedbackRequest) error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return m.helpers.FormatValidationError(err)
	}
	return nil
}

func (m *ModelResource) FeedbackSeeders() error {
	m.logger.Info("Seeding Feedback")
	return nil
}
