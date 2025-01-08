package auth

import (
	"fmt"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	tokenProvider *providers.TokenService
}

func NewAuthHandler(
	tokenProvider *providers.TokenService,
) *AuthHandler {
	return &AuthHandler{
		tokenProvider: tokenProvider,
	}
}

func (as *AuthHandler) GetUserClaims(ctx *gin.Context) (*providers.UserClaims, error) {
	claims, exists := ctx.Get("claims")
	if !exists {
		as.tokenProvider.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("claims not found in context")
	}
	userClaims, ok := claims.(*providers.UserClaims)
	if !ok {
		as.tokenProvider.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("failed to cast claims to *auth.UserClaims")
	}
	return userClaims, nil
}

func (as *AuthHandler) GetUser(ctx *gin.Context) (*providers.UserClaims, error) {
	claims, exists := ctx.Get("claims")
	if !exists {
		as.tokenProvider.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("claims not found in context")
	}
	userClaims, ok := claims.(*providers.UserClaims)
	if !ok {
		as.tokenProvider.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("failed to cast claims to *auth.UserClaims")
	}
	return userClaims, nil
}
