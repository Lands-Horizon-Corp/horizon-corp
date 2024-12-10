package models

import (
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
	Email        string `json:"email"`
	Description  string `json:"description"`
	FeedbackType string `json:"feedbackType"`
}

func (m *ModelResource) FeedbackToResource(feedback *Feedback) *FeedbackResource {
	if feedback == nil {
		return nil
	}
	return &FeedbackResource{
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

func (m *ModelResource) FeedbackSeeders() error {
	m.logger.Info("Seeding Feedback")
	return nil
}
