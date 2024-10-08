package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{DB: db}
}

func (r *AdminRepository) Create(admin *models.Admin) error {
	return r.DB.Create(admin).Error
}

func (r *AdminRepository) GetAll() ([]models.Admin, error) {
	var admin []models.Admin
	err := r.DB.Find(&admin).Error
	return admin, err
}

func (r *AdminRepository) GetByID(id uint) (models.Admin, error) {
	var admin models.Admin
	err := r.DB.First(&admin, id).Error
	return admin, err
}

func (r *AdminRepository) Update(id uint, admin *models.Admin) error {
	admin.ID = id
	return r.DB.Save(admin).Error
}

func (r *AdminRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Admin{}, id).Error
}
