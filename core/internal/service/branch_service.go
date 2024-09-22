package service

import (
	"horizon-core/internal/models"
	"horizon-core/internal/repository"

	"gorm.io/gorm"
)

type BranchRepository struct {
	*repository.ModelRepository[models.Branch]
}

func NewBranchRepository(db *gorm.DB) *BranchRepository {
	return &BranchRepository{
		ModelRepository: repository.NewModelRepository[models.Branch](db),
	}
}
