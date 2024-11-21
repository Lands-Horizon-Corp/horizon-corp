package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type MediaRepository struct {
	*Repository[models.Media]
}

func NewMediaRepository(db *gorm.DB) *MediaRepository {
	return &MediaRepository{
		Repository: NewRepository[models.Media](db),
	}
}
