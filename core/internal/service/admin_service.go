package service

import (
	"horizon-core/internal/models"
	"horizon-core/internal/repository"

	"gorm.io/gorm"
)

type AdminRepository struct {
	*repository.ModelRepository[models.Admin]
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{
		ModelRepository: repository.NewModelRepository[models.Admin](db),
	}
}
