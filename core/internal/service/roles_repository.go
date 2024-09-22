package service

import (
	"horizon-core/internal/models"
	"horizon-core/internal/repository"

	"gorm.io/gorm"
)

type RoleRepository struct {
	*repository.ModelRepository[models.Role]
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		ModelRepository: repository.NewModelRepository[models.Role](db),
	}
}
