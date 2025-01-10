package models

import (
	"time"

	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model

	// Fields
	Email        string `gorm:"type:varchar(255);unsigned" json:"email"`
	Description  string `gorm:"type:text;unsigned" json:"description"`
	FeedbackType string `gorm:"type:enum('bug', 'feature', 'general');unsigned" json:"feedback_type"`
}

type FeedbackResource struct {
	ID        uint   `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`

	Email        string `json:"email"`
	Description  string `json:"description"`
	FeedbackType string `json:"feedbackType"`
}

func (m *ModelTransformer) FeedbackToResource(feedback *Feedback) *FeedbackResource {
	if feedback == nil {
		return nil
	}

	return &FeedbackResource{
		ID:           feedback.ID,
		CreatedAt:    feedback.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    feedback.UpdatedAt.Format(time.RFC3339),
		Email:        feedback.Email,
		Description:  feedback.Description,
		FeedbackType: feedback.FeedbackType,
	}
}

func (m *ModelTransformer) FeedbackToResourceList(feedbackList []*Feedback) []*FeedbackResource {
	if feedbackList == nil {
		return nil
	}

	var feedbackResources []*FeedbackResource
	for _, feedback := range feedbackList {
		feedbackResources = append(feedbackResources, m.FeedbackToResource(feedback))
	}
	return feedbackResources
}
