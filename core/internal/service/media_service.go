package service

import (
	"horizon-core/internal/models"
	"horizon-core/internal/repository"

	"gorm.io/gorm"
)

type MediaRepository struct {
	*repository.ModelRepository[models.Media]
}

func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{
		ModelRepository: repository.NewModelRepository[models.Media](db),
	}
}
