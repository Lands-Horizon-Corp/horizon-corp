package admin

import "github.com/Lands-Horizon-Corp/horizon-corp/internal/config"

type AdminRepository struct {
	cfg *config.AppConfig
}

func NewAdminRepository(cfg *config.AppConfig) *AdminRepository {
	return &AdminRepository{
		cfg: cfg,
	}
}
