package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type RoleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{DB: db}
}

func (r *RoleRepository) Create(Role *models.Role) error {
	err := r.DB.Create(Role).Error
	return handleDBError(err)
}

func (r *RoleRepository) GetAll() ([]models.Role, error) {
	var Roles []models.Role
	err := r.DB.Find(&Roles).Error
	return Roles, handleDBError(err)
}

func (r *RoleRepository) GetByID(id uint) (models.Role, error) {
	var Role models.Role
	err := r.DB.First(&Role, id).Error
	return Role, handleDBError(err)
}

func (r *RoleRepository) Update(id uint, Role *models.Role) error {
	Role.ID = id
	err := r.DB.Save(Role).Error
	return handleDBError(err)
}

func (r *RoleRepository) Delete(id uint) error {
	err := r.DB.Delete(&models.Role{}, id).Error
	return handleDBError(err)
}
