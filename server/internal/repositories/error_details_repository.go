package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type ErrorDetailRepository struct {
	DB *gorm.DB
}

func NewErrorDetailRepository(db *gorm.DB) *ErrorDetailRepository {
	return &ErrorDetailRepository{DB: db}
}

func (r *ErrorDetailRepository) Create(errorDetails *models.ErrorDetail) error {
	return r.DB.Create(errorDetails).Error
}

func (r *ErrorDetailRepository) GetAll() ([]models.ErrorDetail, error) {
	var errorDetailss []models.ErrorDetail
	err := r.DB.Find(&errorDetailss).Error
	return errorDetailss, err
}

func (r *ErrorDetailRepository) GetByID(id uint) (models.ErrorDetail, error) {
	var errorDetails models.ErrorDetail
	err := r.DB.First(&errorDetails, id).Error
	return errorDetails, err
}

func (r *ErrorDetailRepository) Update(errorDetails *models.ErrorDetail) error {
	return r.DB.Save(errorDetails).Error
}

func (r *ErrorDetailRepository) Delete(id uint) error {
	return r.DB.Delete(&models.ErrorDetail{}, id).Error
}
