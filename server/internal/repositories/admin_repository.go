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
	err := r.DB.Create(admin).Error
	return handleDBError(err)
}

func (r *AdminRepository) GetAll() ([]models.Admin, error) {
	var admin []models.Admin
	err := r.DB.Find(&admin).Error
	return admin, handleDBError(err)
}

func (r *AdminRepository) GetByID(id uint) (models.Admin, error) {
	var admin models.Admin
	err := r.DB.First(&admin, id).Error
	return admin, handleDBError(err)
}

func (r *AdminRepository) Update(id uint, admin *models.Admin) error {
	admin.ID = id
	err := r.DB.Save(admin).Error
	return handleDBError(err)
}

func (r *AdminRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Admin{}, id).Error
	return handleDBError(err)
}

func (r *AdminRepository) GetByEmail(email string) (models.Admin, error) {
	var admin models.Admin
	err := r.DB.Where("email = ?", email).First(&admin).Error
	return admin, handleDBError(err)
}

func (r *AdminRepository) FindByEmailOrUsername(email, username string) (models.Admin, error) {
	var admin models.Admin
	if email != "" {
		err := r.DB.Where("email = ?", email).First(&admin).Error
		return admin, handleDBError(err)
	} else {
		err := r.DB.Where("username = ?", username).First(&admin).Error
		return admin, handleDBError(err)
	}
}
