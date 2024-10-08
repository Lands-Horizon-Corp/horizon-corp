package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type OwnerRepository struct {
	DB *gorm.DB
}

func NewOwnerRepository(db *gorm.DB) *OwnerRepository {
	return &OwnerRepository{DB: db}
}

func (r *OwnerRepository) Create(owner *models.Owner) error {
	return r.DB.Create(owner).Error
}

func (r *OwnerRepository) GetAll() ([]models.Owner, error) {
	var owner []models.Owner
	err := r.DB.Find(&owner).Error
	return owner, err
}

func (r *OwnerRepository) GetByID(id uint) (models.Owner, error) {
	var owner models.Owner
	err := r.DB.First(&owner, id).Error
	return owner, err
}

func (r *OwnerRepository) Update(owner *models.Owner) error {
	return r.DB.Save(owner).Error
}

func (r *OwnerRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Owner{}, id).Error
}
