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
	err := r.DB.Create(owner).Error
	return handleDBError(err)
}

func (r *OwnerRepository) GetAll() ([]models.Owner, error) {
	var owner []models.Owner
	err := r.DB.Find(&owner).Error
	return owner, handleDBError(err)
}

func (r *OwnerRepository) GetByID(id uint) (models.Owner, error) {
	var owner models.Owner
	err := r.DB.First(&owner, id).Error
	return owner, handleDBError(err)
}

func (r *OwnerRepository) Update(id uint, owner *models.Owner) error {
	owner.ID = id
	err := r.DB.Save(owner).Error
	return handleDBError(err)
}

func (r *OwnerRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Owner{}, id).Error
	return handleDBError(err)
}
