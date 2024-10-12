package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type FeedbackRepository struct {
	DB *gorm.DB
}

func NewFeedbackRepository(db *gorm.DB) *FeedbackRepository {
	return &FeedbackRepository{DB: db}
}

func (r *FeedbackRepository) Create(feedback *models.Feedback) error {
	err := r.DB.Create(feedback).Error
	return handleDBError(err)
}

func (r *FeedbackRepository) GetAll() ([]models.Feedback, error) {
	var feedback []models.Feedback
	err := r.DB.Find(&feedback).Error
	return feedback, handleDBError(err)
}

func (r *FeedbackRepository) GetByID(id uint) (models.Feedback, error) {
	var feedback models.Feedback
	err := r.DB.First(&feedback, id).Error
	return feedback, handleDBError(err)
}

func (r *FeedbackRepository) Update(id uint, feedback *models.Feedback) error {
	feedback.ID = id
	err := r.DB.Save(feedback).Error
	return handleDBError(err)
}

func (r *FeedbackRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Feedback{}, id).Error
	return handleDBError(err)
}
