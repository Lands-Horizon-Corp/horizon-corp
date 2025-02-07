package auth

import (
	"fmt"
	"strconv"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers/filter"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/auth/auth_accounts"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
)

// AuthFootstep "ctx", "Action", "Description"

type AuthFootstep struct {
	middle        *middleware.Middleware
	tokenProvider *providers.TokenService
	authAccount   *auth_accounts.AuthAccount
	modelResource *models.ModelResource
}

func NewAuthFootstep(
	middle *middleware.Middleware,
	tokenProvider *providers.TokenService,
	authAccount *auth_accounts.AuthAccount,
	modelResource *models.ModelResource,
) *AuthFootstep {
	return &AuthFootstep{
		middle:        middle,
		tokenProvider: tokenProvider,
		authAccount:   authAccount,
		modelResource: modelResource,
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

func (af *AuthFootstep) Save(ctx *gin.Context, action, description string) (*models.Footstep, error) {
	userClaims, err := af.UserClaims(ctx)
	if err != nil {
		return nil, fmt.Errorf("user not authenticated")
	}
	return af.authAccount.AccountFootstep(userClaims.AccountType, userClaims.ID, action, description)
}

func (af *AuthFootstep) Self(ctx *gin.Context) (filter.FilterPages[models.Footstep], error) {
	userClaims, err := af.UserClaims(ctx)
	if err != nil {
		return filter.FilterPages[models.Footstep]{}, fmt.Errorf("user not authenticated")
	}
	filterParam := ctx.Query("filter")
	if filterParam == "" {
		return filter.FilterPages[models.Footstep]{}, fmt.Errorf("filter parameter is required")
	}
	pageIndexStr := ctx.Query("pageIndex")
	pageSizeStr := ctx.Query("pageSize")
	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		pageIndex = 0
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}
	return af.modelResource.FootstepFilterForSelf(userClaims.AccountType, userClaims.ID, filterParam, pageIndex, pageSize)
}

func (af *AuthFootstep) Peers(ctx *gin.Context) (filter.FilterPages[models.Footstep], error) {
	userClaims, err := af.UserClaims(ctx)
	if err != nil {
		return filter.FilterPages[models.Footstep]{}, fmt.Errorf("user not authenticated")
	}
	filterParam := ctx.Query("filter")
	if filterParam == "" {
		return filter.FilterPages[models.Footstep]{}, fmt.Errorf("filter parameter is required")
	}
	pageIndexStr := ctx.Query("pageIndex")
	pageSizeStr := ctx.Query("pageSize")
	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil {
		pageIndex = 0
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		pageSize = 10
	}
	return af.modelResource.FootstepFilterForPeers(userClaims.AccountType, userClaims.ID, filterParam, pageSize, pageIndex)
}
