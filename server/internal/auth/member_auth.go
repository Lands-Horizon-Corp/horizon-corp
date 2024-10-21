package auth

import (
	"horizon/server/internal/models"
	"time"

	"go.uber.org/zap"
)

type MemberAuthService struct {
	tokenService TokenService
	logger       *zap.Logger
}

func NewMemberAuthService(tokenService TokenService, logger *zap.Logger) *MemberAuthService {
	return &MemberAuthService{
		tokenService: tokenService,
		logger:       logger,
	}
}

func (s *MemberAuthService) GenerateMemberToken(member models.Member, expiration time.Duration) (string, error) {

	claims := &UserClaims{
		ID:          member.ID,
		AccountType: "Member",
	}

	token, err := s.tokenService.GenerateToken(claims, 0)
	if err != nil {
		s.logger.Error("Failed to generate member token", zap.Error(err))
		return "", err
	}
	return token, nil
}
