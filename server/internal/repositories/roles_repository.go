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
	return r.DB.Create(roles).Error
}

func (r *RolesRepository) GetAll() ([]models.Roles, error) {
	var roless []models.Roles
	err := r.DB.Find(&roless).Error
	return roless, err
}

func (r *RolesRepository) GetByID(id uint) (models.Roles, error) {
	var roles models.Roles
	err := r.DB.First(&roles, id).Error
	return roles, err
}

func (r *RolesRepository) Update(id uint, roles *models.Roles) error {
	roles.ID = id
	return r.DB.Save(roles).Error
}

func (r *RolesRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Roles{}, id).Error
}
