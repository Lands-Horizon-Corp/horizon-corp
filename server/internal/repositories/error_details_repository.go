package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type ErrorDetailsRepository struct {
	DB *gorm.DB
}

func NewErrorDetailsRepository(db *gorm.DB) *ErrorDetailsRepository {
	return &ErrorDetailsRepository{DB: db}
}

func (r *ErrorDetailsRepository) Create(errorDetails *models.ErrorDetails) error {
	return r.DB.Create(errorDetails).Error
}

func (r *ErrorDetailsRepository) GetAll() ([]models.ErrorDetails, error) {
	var errorDetailss []models.ErrorDetails
	err := r.DB.Find(&errorDetailss).Error
	return errorDetailss, err
}

func (r *ErrorDetailsRepository) GetByID(id uint) (models.ErrorDetails, error) {
	var errorDetails models.ErrorDetails
	err := r.DB.First(&errorDetails, id).Error
	return errorDetails, err
}

func (r *ErrorDetailsRepository) Update(errorDetails *models.ErrorDetails) error {
	return r.DB.Save(errorDetails).Error
}

func (r *ErrorDetailsRepository) Delete(id uint) error {
	return r.DB.Delete(&models.ErrorDetails{}, id).Error
}
