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
	return r.DB.Create(feedback).Error
}

func (r *FeedbackRepository) GetAll() ([]models.Feedback, error) {
	var feedback []models.Feedback
	err := r.DB.Find(&feedback).Error
	return feedback, err
}

func (r *FeedbackRepository) GetByID(id uint) (models.Feedback, error) {
	var feedback models.Feedback
	err := r.DB.First(&feedback, id).Error
	return feedback, err
}

func (r *FeedbackRepository) Update(feedback *models.Feedback) error {
	return r.DB.Save(feedback).Error
}

func (r *FeedbackRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Feedback{}, id).Error
}
