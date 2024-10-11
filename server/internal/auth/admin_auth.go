package auth

import (
	"horizon/server/internal/models"

	"go.uber.org/zap"
)

type AdminService struct {
	tokenService TokenService
	logger       *zap.Logger
}

func NewAdminAuthService(tokenService TokenService, logger *zap.Logger) *AdminService {
	return &AdminService{
		tokenService: tokenService,
		logger:       logger,
	}
}

func (s *AdminService) GenerateAdminToken(admin models.Admin) (string, error) {
	claims := &UserClaims{
		ID:                admin.ID,
		Mode:              "admin",
		FirstName:         admin.FirstName,
		LastName:          admin.LastName,
		PermanentAddress:  admin.PermanentAddress,
		Description:       admin.Description,
		Birthdate:         admin.Birthdate,
		Email:             admin.Email,
		IsEmailVerified:   admin.IsEmailVerified,
		IsContactVerified: admin.IsContactVerified,
		ContactNumber:     admin.ContactNumber,
		MediaID:           admin.MediaID,
	}

	token, err := s.tokenService.GenerateToken(claims)
	if err != nil {
		s.logger.Error("Failed to generate admin token", zap.Error(err))
		return "", err
	}
	return token, nil
}
