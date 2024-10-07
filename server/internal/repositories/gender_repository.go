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
	return r.DB.Create(gender).Error
}

func (r *GenderRepository) GetAll() ([]models.Gender, error) {
	var genders []models.Gender
	err := r.DB.Find(&genders).Error
	return genders, err
}

func (r *GenderRepository) GetByID(id uint) (models.Gender, error) {
	var gender models.Gender
	err := r.DB.First(&gender, id).Error
	return gender, err
}

func (r *GenderRepository) Update(gender *models.Gender) error {
	return r.DB.Save(gender).Error
}

func (r *GenderRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Gender{}, id).Error
}
