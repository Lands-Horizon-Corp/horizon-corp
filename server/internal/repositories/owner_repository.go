package repositories

import (
	"horizon/server/config"
	"horizon/server/helpers"
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type OwnerRepository struct {
	*Repository[models.Owner]
}

func NewOwnerRepository(db *gorm.DB) *OwnerRepository {
	return &OwnerRepository{
		Repository: NewRepository[models.Owner](db),
	}
}

func (r *OwnerRepository) GetByEmail(email string) (*models.Owner, error) {
	var owner models.Owner
	err := r.DB.Preload("Media").Where("email = ?", email).First(&owner).Error
	return &owner, handleDBError(err)
}

func (r *OwnerRepository) GetByContactNumber(contactNumber string) (*models.Owner, error) {
	var owner models.Owner
	err := r.DB.Preload("Media").Where("contact_number = ?", contactNumber).First(&owner).Error
	return &owner, handleDBError(err)
}

func (r *OwnerRepository) GetByUsername(username string) (*models.Owner, error) {
	var owner models.Owner
	err := r.DB.Preload("Media").Where("username = ?", username).First(&owner).Error
	return &owner, handleDBError(err)
}

func (r *OwnerRepository) FindByEmailUsernameOrContact(input string) (*models.Owner, error) {
	switch helpers.GetKeyType(input) {
	case "contact":
		return r.GetByContactNumber(input)
	case "email":
		return r.GetByEmail(input)
	default:
		return r.GetByUsername(input)
	}
}

func (r *OwnerRepository) UpdatePassword(id uint, password string) error {
	newPassword, err := config.HashPassword(password)
	if err != nil {
		return err
	}
	updated := &models.Owner{Password: newPassword}
	_, err = r.Repository.UpdateColumns(id, *updated, []string{})
	return err
}
