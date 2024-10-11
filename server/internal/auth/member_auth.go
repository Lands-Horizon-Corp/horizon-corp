package auth

import (
	"horizon/server/internal/models"

	"go.uber.org/zap"
)

type MemberService struct {
	tokenService TokenService
	logger       *zap.Logger
}

func NewMemberAuthService(tokenService TokenService, logger *zap.Logger) *MemberService {
	return &MemberService{
		tokenService: tokenService,
		logger:       logger,
	}
}

func (s *MemberService) GenerateMemberToken(member models.Member) (string, error) {
	claims := &UserClaims{
		ID:                member.ID,
		Mode:              "member",
		FirstName:         member.FirstName,
		LastName:          member.LastName,
		PermanentAddress:  member.PermanentAddress,
		Description:       member.Description,
		Birthdate:         member.Birthdate,
		Email:             member.Email,
		IsEmailVerified:   member.IsEmailVerified,
		IsContactVerified: member.IsContactVerified,
		ContactNumber:     member.ContactNumber,
		MediaID:           member.MediaID,
	}

	token, err := s.tokenService.GenerateToken(claims)
	if err != nil {
		s.logger.Error("Failed to generate member token", zap.Error(err))
		return "", err
	}
	return token, nil
}
