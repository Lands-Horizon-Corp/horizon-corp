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

func (r *FeedbackRepository) Create(gender *models.Feedback) error {
	return r.DB.Create(gender).Error
}

func (r *FeedbackRepository) GetAll() ([]models.Feedback, error) {
	var genders []models.Feedback
	err := r.DB.Find(&genders).Error
	return genders, err
}

func (r *FeedbackRepository) GetByID(id uint) (models.Feedback, error) {
	var gender models.Feedback
	err := r.DB.First(&gender, id).Error
	return gender, err
}

func (r *FeedbackRepository) Update(gender *models.Feedback) error {
	return r.DB.Save(gender).Error
}

func (r *FeedbackRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Feedback{}, id).Error
}
