package auth

import (
	"horizon/server/internal/models"

	"go.uber.org/zap"
)

type EmployeeService struct {
	tokenService TokenService
	logger       *zap.Logger
}

func NewEmployeeAuthService(tokenService TokenService, logger *zap.Logger) *EmployeeService {
	return &EmployeeService{
		tokenService: tokenService,
		logger:       logger,
	}
}

func (s *EmployeeService) GenerateEmployeeToken(employee models.Employee) (string, error) {
	claims := &UserClaims{
		ID:                employee.ID,
		Mode:              "employee",
		FirstName:         employee.FirstName,
		LastName:          employee.LastName,
		PermanentAddress:  employee.PermanentAddress,
		Description:       employee.Description,
		Birthdate:         employee.Birthdate,
		Email:             employee.Email,
		IsEmailVerified:   employee.IsEmailVerified,
		IsContactVerified: employee.IsContactVerified,
		ContactNumber:     employee.ContactNumber,
		MediaID:           employee.MediaID,
	}

	token, err := s.tokenService.GenerateToken(claims)
	if err != nil {
		s.logger.Error("Failed to generate employee token", zap.Error(err))
		return "", err
	}
	return token, nil
}
