package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type GenderRepository struct {
	DB *gorm.DB
}

func NewGenderRepository(db *gorm.DB) *GenderRepository {
	return &GenderRepository{DB: db}
}

func (r *GenderRepository) Create(gender *models.Gender) error {
	err := r.DB.Create(gender).Error
	return handleDBError(err)
}

func (r *GenderRepository) GetAll() ([]models.Gender, error) {
	var genders []models.Gender
	err := r.DB.Find(&genders).Error
	return genders, handleDBError(err)
}

func (r *GenderRepository) GetByID(id uint) (models.Gender, error) {
	var gender models.Gender
	err := r.DB.First(&gender, id).Error
	return gender, handleDBError(err)
}

func (r *GenderRepository) Update(id uint, gender *models.Gender) error {
	gender.ID = id
	err := r.DB.Save(gender).Error
	return handleDBError(err)
}

func (r *GenderRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Gender{}, id).Error
	return handleDBError(err)
}
