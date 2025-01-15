package handlers

import (
	"fmt"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type CurrentUser struct {
	tokenProvider *providers.TokenService
}

func NewCurrentUser(
	tokenProvider *providers.TokenService,
) *CurrentUser {
	return &CurrentUser{
		tokenProvider: tokenProvider,
	}
}

func (c *CurrentUser) User(ctx *gin.Context) {

}

func (c *CurrentUser) Claims(ctx *gin.Context) (*providers.UserClaims, error) {
	claims, exists := ctx.Get("claims")
	if !exists {
		c.tokenProvider.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("claims not found in context")
	}
	userClaims, ok := claims.(*providers.UserClaims)
	if !ok {
		c.tokenProvider.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("failed to cast claims to *auth.UserClaims")
	}
	return userClaims, nil
}
