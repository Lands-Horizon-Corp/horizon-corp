package admin

import "github.com/Lands-Horizon-Corp/horizon-corp/internal/config"

type AdminProvider struct {
	cfg *config.AppConfig
}

func NewAdminProvider(cfg *config.AppConfig) *AdminProvider {
	return &AdminProvider{
		cfg: cfg,
	}
}
