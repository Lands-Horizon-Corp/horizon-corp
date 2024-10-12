package auth

import (
	"horizon/server/internal/models"

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

func (s *EmployeeAuthService) GenerateEmployeeToken(employee models.Employee) (string, error) {
	claims := &UserClaims{
		ID:          employee.ID,
		AccountType: "Employee",
	}

	token, err := s.tokenService.GenerateToken(claims)
	if err != nil {
		s.logger.Error("Failed to generate employee token", zap.Error(err))
		return "", err
	}
	return token, nil
}
