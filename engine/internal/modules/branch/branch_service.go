package branch

import (
	"fmt"
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers/filter"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
)

type BranchService struct {
	controller    *managers.Controller[models.Branch, models.BranchRequest, models.BranchResource]
	db            *providers.DatabaseService
	engine        *providers.EngineService
	modelResource *models.ModelResource
	middle        *middleware.Middleware
	tokenProvider *providers.TokenService
}

func NewBranchService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	modelResource *models.ModelResource,
	middle *middleware.Middleware,
	tokenProvider *providers.TokenService,
) *BranchService {
	controller := managers.NewController(
		modelResource.BranchDB,
		modelResource.ValidateBranchRequest,
		modelResource.BranchToResource,
		modelResource.BranchToResourceList,
	)

	return &BranchService{
		controller:    controller,
		db:            db,
		engine:        engine,
		modelResource: modelResource,
		middle:        middle,
		tokenProvider: tokenProvider,
	}
}

// BranchFilterForAdmin
// BranchFilterForOwner

func (ts *BranchService) getUserClaims(ctx *gin.Context) (*providers.UserClaims, error) {
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

func (as *BranchService) SearchFilter(ctx *gin.Context) {
	// Retrieve the 'filter' query parameter
	filterParam := ctx.Query("filter")
	if filterParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filter parameter is required"})
		return
	}

	// Authenticate the user and retrieve their claims
	userClaims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	var branches filter.FilterPages[models.Branch]

	// Determine which filter function to use based on account type
	switch userClaims.AccountType {
	case "Owner":
		branches, err = as.modelResource.BranchFilterForOwner(filterParam, userClaims.ID)
	case "Admin":
		branches, err = as.modelResource.BranchFilterForAdmin(filterParam)
	default:
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions."})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Branches not found."})
		return
	}

	data := as.modelResource.BranchToResourceList(branches.Data)
	if data == nil {
		data = []*models.BranchResource{}
	}

	response := gin.H{
		"data":      data,
		"pageIndex": branches.PageIndex,
		"totalPage": branches.TotalPage,
		"pageSize":  branches.PageSize,
		"totalSize": branches.TotalSize,
		"pages":     branches.Pages,
	}

	ctx.JSON(http.StatusOK, response)
}

func (as *BranchService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/branch")
	routes.Use(as.middle.AuthMiddleware())
	{
		routes.GET("/search", as.SearchFilter)
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)
	}
}
