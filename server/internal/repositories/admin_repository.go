package repositories

import (
	"horizon/server/config"
	"horizon/server/helpers"
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type AdminRepository struct {
	*Repository[models.Admin]
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{
		Repository: NewRepository[models.Admin](db),
	}
}

func (r *AdminRepository) GetByEmail(email string) (*models.Admin, error) {
	var admin models.Admin
	err := r.DB.Preload("Media").Where("email = ?", email).First(&admin).Error
	return &admin, handleDBError(err)
}

func (r *AdminRepository) GetByContactNumber(contactNumber string) (*models.Admin, error) {
	var admin models.Admin
	err := r.DB.Preload("Media").Where("contact_number = ?", contactNumber).First(&admin).Error
	return &admin, handleDBError(err)
}

func (r *AdminRepository) GetByUsername(username string) (*models.Admin, error) {
	var admin models.Admin
	err := r.DB.Preload("Media").Where("username = ?", username).First(&admin).Error
	return &admin, handleDBError(err)
}

func (r *AdminRepository) FindByEmailUsernameOrContact(input string) (*models.Admin, error) {
	switch helpers.GetKeyType(input) {
	case "contact":
		return r.GetByContactNumber(input)
	case "email":
		return r.GetByEmail(input)
	default:
		return r.GetByUsername(input)
	}
}

func (r *AdminRepository) UpdatePassword(id uint, password string) error {
	newPassword, err := config.HashPassword(password)
	if err != nil {
		return err
	}
	updated := &models.Admin{Password: newPassword}
	_, err = r.Repository.UpdateColumns(id, *updated, []string{})
	return err
}
