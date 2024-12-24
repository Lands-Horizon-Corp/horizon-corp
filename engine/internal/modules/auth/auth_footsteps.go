package auth

import (
	"fmt"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
)

// AuthFootstep "ctx", "Action", "Description"

type AuthFootstep struct {
	middle        *middleware.Middleware
	tokenProvider *providers.TokenService
}

func NewAuthFootstep(
	middle *middleware.Middleware,
	tokenProvider *providers.TokenService,
) *AuthFootstep {
	return &AuthFootstep{
		middle:        middle,
		tokenProvider: tokenProvider,
	}
}

func (ts *AuthFootstep) UserClaims(ctx *gin.Context) (*providers.UserClaims, error) {
	claims, exists := ctx.Get("claims")
	if !exists {
		ts.tokenProvider.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("claims not found in context")
	}

	userClaims, ok := claims.(*providers.UserClaims)
	if !ok {
		ts.tokenProvider.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("failed to cast claims to *auth.UserClaims")
	}

	return userClaims, nil
}

func (af *AuthFootstep) Save(ctx *gin.Context, action, description string) {
	// userClaims, err := af.UserClaims(ctx)
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
	// 	return
	// }

}

// type AuthService struct {
// 	engine        *providers.EngineService
// 	middle        *middleware.Middleware
// 	tokenProvider *providers.TokenService
// 	otpProvider   *providers.OTPService
// 	authProvider  *AuthProvider
// 	authAccount   *auth_accounts.AuthAccount
// 	modelResource *models.ModelResource
// }

// func NewAuthService(
// 	engine *providers.EngineService,
// 	middle *middleware.Middleware,
// 	tokenProvider *providers.TokenService,
// 	otpProvider *providers.OTPService,
// 	authProvider *AuthProvider,
// 	authAccount *auth_accounts.AuthAccount,
// 	modelResource *models.ModelResource,
// ) *AuthService {
// 	return &AuthService{
// 		engine:        engine,
// 		middle:        middle,
// 		otpProvider:   otpProvider,
// 		tokenProvider: tokenProvider,
// 		authProvider:  authProvider,
// 		authAccount:   authAccount,
// 		modelResource: modelResource,
// 	}
// }
