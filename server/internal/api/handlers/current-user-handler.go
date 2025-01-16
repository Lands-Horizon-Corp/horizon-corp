package handlers

import (
	"fmt"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	IPAddress      string `json:"ipAddress"`
	UserAgent      string `json:"userAgent"`
	Referer        string `json:"referer"`
	Location       string `json:"location"`
	AcceptLanguage string `json:"acceptLanguage"`
}

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

func (c *CurrentUser) Claims(ctx *gin.Context) (*providers.UserClaims, *UserInfo, error) {
	claims, exists := ctx.Get("claims")
	if !exists {
		c.tokenProvider.ClearTokenCookie(ctx)
		return nil, nil, fmt.Errorf("claims not found in context")
	}
	userClaims, ok := claims.(*providers.UserClaims)
	if !ok {
		c.tokenProvider.ClearTokenCookie(ctx)
		return nil, nil, fmt.Errorf("failed to cast claims to *auth.UserClaims")
	}

	// Populate UserInfo
	userInfo := &UserInfo{
		IPAddress:      ctx.ClientIP(),
		UserAgent:      ctx.GetHeader("User-Agent"),
		Referer:        ctx.GetHeader("Referer"),
		Location:       ctx.GetHeader("X-Location"),
		AcceptLanguage: ctx.GetHeader("Accept-Language"),
	}

	return userClaims, userInfo, nil
}
