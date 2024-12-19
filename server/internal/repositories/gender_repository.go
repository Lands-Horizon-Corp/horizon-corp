package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type GenderRepository struct {
	*Repository[models.Gender]
}

func NewGenderRepository(db *gorm.DB) *GenderRepository {
	return &GenderRepository{
		Repository: NewRepository[models.Gender](db),
	}
}
