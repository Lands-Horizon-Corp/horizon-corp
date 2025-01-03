package auth

import (
	"horizon/server/internal/models"
	"time"

	"go.uber.org/zap"
)

type EmployeeAuthService struct {
	tokenService TokenService
	logger       *zap.Logger
}

func NewEmployeeAuthService(tokenService TokenService, logger *zap.Logger) *EmployeeAuthService {
	return &EmployeeAuthService{
		tokenService: tokenService,
		logger:       logger,
	}
}

func (s *EmployeeAuthService) GenerateEmployeeToken(employee *models.Employee, expiration time.Duration) (string, error) {
	claims := &UserClaims{
		ID:          employee.ID,
		AccountType: "Employee",
	}

	token, err := s.tokenService.GenerateToken(claims, expiration)
	if err != nil {
		s.logger.Error("Failed to generate employee token", zap.Error(err))
		return "", err
	}
	return token, nil
}
