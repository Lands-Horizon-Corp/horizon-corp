package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type FeedbackRepository struct {
	*Repository[models.Feedback]
}

func NewFeedbackRepository(db *gorm.DB) *FeedbackRepository {
	return &FeedbackRepository{
		Repository: NewRepository[models.Feedback](db),
	}
}
