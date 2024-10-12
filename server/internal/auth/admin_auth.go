package auth

import (
	"horizon/server/config"
	"horizon/server/internal/models"

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

func (s *AdminAuthService) GenerateAdminToken(admin models.Admin) (string, error) {
	claims := &UserClaims{
		ID:          admin.ID,
		AccountType: "Admin",
	}

	token, err := s.tokenService.GenerateToken(claims)
	if err != nil {
		s.logger.Error("Failed to generate admin token", zap.Error(err))
		return "", err
	}
	return token, nil
}
