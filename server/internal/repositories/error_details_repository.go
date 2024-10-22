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
	err := r.DB.Create(errorDetails).Error
	return handleDBError(err)
}

func (r *ErrorDetailRepository) GetAll() ([]models.ErrorDetail, error) {
	var errorDetailss []models.ErrorDetail
	err := r.DB.Find(&errorDetailss).Error
	return errorDetailss, handleDBError(err)
}

func (r *ErrorDetailRepository) GetByID(id uint) (models.ErrorDetail, error) {
	var errorDetails models.ErrorDetail
	err := r.DB.First(&errorDetails, id).Error
	return errorDetails, handleDBError(err)
}

func (r *ErrorDetailRepository) Update(id uint, errorDetails *models.ErrorDetail) error {
	errorDetails.ID = id
	err := r.DB.Save(errorDetails).Error
	return handleDBError(err)
}

func (r *ErrorDetailRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.ErrorDetail{}, id).Error
	return handleDBError(err)
}
