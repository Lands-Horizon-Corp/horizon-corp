package auth_accounts

import (
	"errors"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"go.uber.org/zap"
)

type AuthAccount struct {
	cfg         *config.AppConfig
	engine      *providers.EngineService
	middle      *middleware.Middleware
	otpProvider *providers.OTPService

	tokenProvider *providers.TokenService
	logger        *providers.LoggerService
}

func NewAuthAccount(
	cfg *config.AppConfig,
	engine *providers.EngineService,
	middle *middleware.Middleware,
	otpProvider *providers.OTPService,

	tokenProvider *providers.TokenService,
	logger *providers.LoggerService,
) *AuthAccount {
	return &AuthAccount{
		cfg:           cfg,
		engine:        engine,
		middle:        middle,
		otpProvider:   otpProvider,
		tokenProvider: tokenProvider,
	}
}

const tokenExpiration = time.Hour * 12

func (at *AuthAccount) GenerateUserToken(id uint, accountType string) (*string, error) {
	validTypes := []string{"Member", "Employee", "Admin", "Owner"}
	for _, validType := range validTypes {
		if accountType == validType {
			return nil, errors.New("invalid account type")
		}
	}
	claims := &providers.UserClaims{
		ID:          id,
		AccountType: accountType,
	}
	token, err := at.tokenProvider.GenerateToken(claims, tokenExpiration)
	if err != nil {
		at.logger.Error("Failed to generate admin token", zap.Error(err))
		return nil, err
	}
	return &token, nil
}
