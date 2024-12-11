package branch

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
)

type BranchRepository struct {
	cfg   *config.AppConfig
	model *models.ModelResource
}

func NewBranchRepository(
	cfg *config.AppConfig,
	model *models.ModelResource,
) *BranchRepository {
	return &BranchRepository{
		cfg:   cfg,
		model: model,
	}
}
