package models

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"gorm.io/gorm"
)

type Feedback struct {
	gorm.Model

	// Fields
	Email        string `gorm:"type:varchar(255);not null" json:"email"`
	Description  string `gorm:"type:text;not null" json:"description"`
	FeedbackType string `gorm:"type:enum('bug', 'feature', 'general');not null" json:"feedback_type"`
}

type FeedbackRepository struct {
	*Repository[Feedback]
}

func NewFeedbackRepository(db *providers.DatabaseService) *FeedbackRepository {
	return &FeedbackRepository{
		Repository: NewRepository[Feedback](db),
	}
}
