package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Feedback struct {
	ID        uuid.UUID      `gorm:"type:char(36);primary_key"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// Fields
	Email        string `gorm:"type:varchar(255);unsigned" json:"email"`
	Description  string `gorm:"type:text;unsigned" json:"description"`
	FeedbackType string `gorm:"type:enum('bug', 'feature', 'general');unsigned" json:"feedback_type"`
}

type FeedbackResource struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt string    `json:"createdAt"`
	UpdatedAt string    `json:"updatedAt"`
	DeletedAt string    `json:"deletedAt"`

	Email        string `json:"email"`
	Description  string `json:"description"`
	FeedbackType string `json:"feedbackType"`
}

func (m *ModelTransformer) FeedbackToResource(feedback *Feedback) *FeedbackResource {
	if feedback == nil {
		return nil
	}

	return &FeedbackResource{

		ID:        feedback.ID,
		CreatedAt: feedback.CreatedAt.Format(time.RFC3339),
		UpdatedAt: feedback.UpdatedAt.Format(time.RFC3339),
		DeletedAt: feedback.DeletedAt.Time.Format(time.RFC3339),

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
