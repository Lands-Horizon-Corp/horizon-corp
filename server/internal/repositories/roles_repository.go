package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type RolesRepository struct {
	DB *gorm.DB
}

func NewRolesRepository(db *gorm.DB) *RolesRepository {
	return &RolesRepository{DB: db}
}

func (r *RolesRepository) Create(roles *models.Roles) error {
	err := r.DB.Create(roles).Error
	return handleDBError(err)
}

func (r *RolesRepository) GetAll() ([]models.Roles, error) {
	var roless []models.Roles
	err := r.DB.Find(&roless).Error
	return roless, handleDBError(err)
}

func (r *RolesRepository) GetByID(id uint) (models.Roles, error) {
	var roles models.Roles
	err := r.DB.First(&roles, id).Error
	return roles, handleDBError(err)
}

func (r *RolesRepository) Update(id uint, roles *models.Roles) error {
	roles.ID = id
	err := r.DB.Save(roles).Error
	return handleDBError(err)
}

func (r *RolesRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Roles{}, id).Error
	return handleDBError(err)
}
