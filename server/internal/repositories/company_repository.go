package repositories

import (
	"horizon/server/internal/models"

	"gorm.io/gorm"
)

type CompaniesRepository struct {
	*Repository[models.Company]
}

func NewCompaniesRepository(db *gorm.DB) *CompaniesRepository {
	return &CompaniesRepository{
		Repository: NewRepository[models.Company](db),
	}
}
