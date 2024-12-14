package auth

import "github.com/Lands-Horizon-Corp/horizon-corp/internal/config"

type AuthProvider struct {
	cfg *config.AppConfig
}

func NewAuthProvider(cfg *config.AppConfig) *AuthProvider {
	return &AuthProvider{
		cfg: cfg,
	}
}
