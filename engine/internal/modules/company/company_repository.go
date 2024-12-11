package company

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
)

type CompanyRepository struct {
	cfg   *config.AppConfig
	model *models.ModelResource
}

func NewCompanyRepository(
	cfg *config.AppConfig,
	model *models.ModelResource,
) *CompanyRepository {
	return &CompanyRepository{
		cfg:   cfg,
		model: model,
	}
}
