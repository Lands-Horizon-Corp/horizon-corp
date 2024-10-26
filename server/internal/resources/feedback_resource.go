package resources

import (
	"horizon/server/internal/models"
	"time"
)

type FeedbackResource struct {
	ID           uint   `json:"id"`
	Email        string `json:"email"`
	Description  string `json:"description"`
	FeedbackType string `json:"feedbackType"`

	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

// Convert models.Feedback to *FeedbackResource
func ToResourceFeedback(feedback *models.Feedback) *FeedbackResource {
	if feedback == nil {
		return nil
	}

	return &FeedbackResource{
		ID:           feedback.ID,
		Email:        feedback.Email,
		Description:  feedback.Description,
		FeedbackType: feedback.FeedbackType,
		CreatedAt:    feedback.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    feedback.UpdatedAt.Format(time.RFC3339),
	}
}

// Convert []*models.Feedback to []*FeedbackResource
func ToResourceListFeedback(feedbackList []*models.Feedback) []*FeedbackResource {
	resourceList := make([]*FeedbackResource, len(feedbackList))
	for i, feedback := range feedbackList {
		resourceList[i] = ToResourceFeedback(feedback)
	}
	return resourceList
}
