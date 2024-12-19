package auth

import (
	"horizon/server/internal/models"
	"time"

	"go.uber.org/zap"
)

type OwnerAuthService struct {
	tokenService TokenService
	logger       *zap.Logger
}

func NewOwnerAuthService(tokenService TokenService, logger *zap.Logger) *OwnerAuthService {
	return &OwnerAuthService{
		tokenService: tokenService,
		logger:       logger,
	}
}

func (s *OwnerAuthService) GenerateOwnerToken(owner *models.Owner, expiration time.Duration) (string, error) {
	claims := &UserClaims{
		ID:          owner.ID,
		AccountType: "Owner",
	}

	token, err := s.tokenService.GenerateToken(claims, expiration)
	if err != nil {
		s.logger.Error("Failed to generate owner token", zap.Error(err))
		return "", err
	}
	return token, nil
}
