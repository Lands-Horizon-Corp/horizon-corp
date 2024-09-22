package service

import (
	"horizon-core/internal/models"
	"horizon-core/internal/repository"

	"gorm.io/gorm"
)

type OwnerRepository struct {
	*repository.ModelRepository[models.Owner]
}

func NewOwnerRepository(db *gorm.DB) *OwnerRepository {
	return &OwnerRepository{
		ModelRepository: repository.NewModelRepository[models.Owner](db),
	}
}
