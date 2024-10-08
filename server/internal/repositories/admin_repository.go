package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type AdminsRepository struct {
	DB *gorm.DB
}

func NewAdminsRepository(db *gorm.DB) *AdminsRepository {
	return &AdminsRepository{DB: db}
}

func (r *AdminsRepository) Create(admin *models.Admin) error {
	return r.DB.Create(admin).Error
}

func (r *AdminsRepository) GetAll() ([]models.Admin, error) {
	var admin []models.Admin
	err := r.DB.Find(&admin).Error
	return admin, err
}

func (r *AdminsRepository) GetByID(id uint) (models.Admin, error) {
	var admin models.Admin
	err := r.DB.First(&admin, id).Error
	return admin, err
}

func (r *AdminsRepository) Update(admin *models.Admin) error {
	return r.DB.Save(admin).Error
}

func (r *AdminsRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Admin{}, id).Error
}
