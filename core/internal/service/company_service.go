package service

import (
	"horizon-core/internal/models"
	"horizon-core/internal/repository"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	*repository.ModelRepository[models.Company]
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{
		ModelRepository: repository.NewModelRepository[models.Company](db),
	}
}
