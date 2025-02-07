package admin

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
)

type AdminRepository struct {
	cfg   *config.AppConfig
	model *models.ModelResource
}

func NewAdminRepository(
	cfg *config.AppConfig,
	model *models.ModelResource,
) *AdminRepository {
	return &AdminRepository{
		cfg:   cfg,
		model: model,
	}
}
