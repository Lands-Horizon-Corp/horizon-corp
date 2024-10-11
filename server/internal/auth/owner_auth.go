package auth

import (
	"horizon/server/internal/models"

	"go.uber.org/zap"
)

type OwnerService struct {
	tokenService TokenService
	logger       *zap.Logger
}

func NewOwnerAuthService(tokenService TokenService, logger *zap.Logger) *OwnerService {
	return &OwnerService{
		tokenService: tokenService,
		logger:       logger,
	}
}

func (s *OwnerService) GenerateOwnerToken(owner models.Owner) (string, error) {
	claims := &UserClaims{
		ID:                owner.ID,
		Mode:              "owner",
		FirstName:         owner.FirstName,
		LastName:          owner.LastName,
		PermanentAddress:  owner.PermanentAddress,
		Description:       owner.Description,
		Birthdate:         owner.Birthdate,
		Email:             owner.Email,
		IsEmailVerified:   owner.IsEmailVerified,
		IsContactVerified: owner.IsContactVerified,
		ContactNumber:     owner.ContactNumber,
		MediaID:           owner.MediaID,
	}

	token, err := s.tokenService.GenerateToken(claims)
	if err != nil {
		s.logger.Error("Failed to generate owner token", zap.Error(err))
		return "", err
	}
	return token, nil
}
