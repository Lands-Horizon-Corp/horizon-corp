package models

import (
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
