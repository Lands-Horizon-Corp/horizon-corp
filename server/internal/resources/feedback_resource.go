package resources

import (
	"horizon/server/internal/models"
	"time"
)

type FeedbackResource struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	Description  string `json:"description"`
	FeedbackType string `json:"feedback_type"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

func ToResourceFeedback(feedback models.Feedback) FeedbackResource {
	return FeedbackResource{
		ID:           feedback.ID,
		Email:        feedback.Email,
		Description:  feedback.Description,
		FeedbackType: feedback.FeedbackType,
		CreatedAt:    feedback.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    feedback.UpdatedAt.Format(time.RFC3339),
	}
}

func ToResourceListFeedback(feedbacks []models.Feedback) []FeedbackResource {
	var resources []FeedbackResource
	for _, feedback := range feedbacks {
		resources = append(resources, ToResourceFeedback(feedback))
	}
	return resources
}
