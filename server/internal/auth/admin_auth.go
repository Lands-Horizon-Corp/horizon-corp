package auth

import (
	"horizon/server/config"
	"horizon/server/internal/models"
	"time"

	"go.uber.org/zap"
)

type AdminAuthService struct {
	tokenService TokenService
	logger       *zap.Logger
	cfg          *config.AppConfig
}

func NewAdminAuthService(
	tokenService TokenService, logger *zap.Logger, cfg *config.AppConfig) *AdminAuthService {
	return &AdminAuthService{
		tokenService: tokenService,
		logger:       logger,
		cfg:          cfg,
	}
}

func (s *AdminAuthService) GenerateAdminToken(admin models.Admin, expiration time.Duration) (string, error) {
	claims := &UserClaims{
		ID:          admin.ID,
		AccountType: "Admin",
	}

	token, err := s.tokenService.GenerateToken(claims, 0)
	if err != nil {
		s.logger.Error("Failed to generate admin token", zap.Error(err))
		return "", err
	}
	return token, nil
}
