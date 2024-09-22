package service

import (
	"horizon-core/internal/models"
	"horizon-core/internal/repository"

	"gorm.io/gorm"
)

type PermissionRepository struct {
	*repository.ModelRepository[models.Permission]
}

func NewPermissionRepository(db *gorm.DB) *PermissionRepository {
	return &PermissionRepository{
		ModelRepository: repository.NewModelRepository[models.Permission](db),
	}
}
