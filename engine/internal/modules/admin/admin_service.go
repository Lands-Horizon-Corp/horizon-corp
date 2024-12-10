package admin

import "github.com/Lands-Horizon-Corp/horizon-corp/internal/config"

type AdminService struct {
	cfg *config.AppConfig
}

func NewAdminService(cfg *config.AppConfig) *AdminService {
	return &AdminService{
		cfg: cfg,
	}
}
