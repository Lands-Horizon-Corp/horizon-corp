package branch

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers/filter"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BranchService struct {
	controller    *managers.Controller[models.Branch, models.BranchRequest, models.BranchResource]
	middle        *middleware.Middleware
	db            *providers.DatabaseService
	engine        *providers.EngineService
	models        *models.ModelResource
	tokenProvider *providers.TokenService
	helpers       *helpers.HelpersFunction
	modelResource *models.ModelResource
}

func NewBranchService(
	db *providers.DatabaseService,
	middle *middleware.Middleware,
	engine *providers.EngineService,
	models *models.ModelResource,
	tokenProvider *providers.TokenService,
	helpers *helpers.HelpersFunction,
	modelResource *models.ModelResource,
) *BranchService {
	controller := managers.NewController(
		modelResource.BranchDB,
		modelResource.ValidateBranchRequest,
		modelResource.BranchToResource,
		modelResource.BranchToResourceList,
	)

	return &BranchService{
		controller:    controller,
		middle:        middle,
		db:            db,
		engine:        engine,
		models:        models,
		tokenProvider: tokenProvider,
		helpers:       helpers,
		modelResource: modelResource,
	}
}

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

func (bs *BranchService) GetAll(ctx *gin.Context) {
	userClaims, err := bs.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	var branches []*models.Branch
	switch userClaims.AccountType {
	case "Owner":
		branches, err = bs.modelResource.BranchGetAllForAdmin()
	case "Admin":
		branches, err = bs.modelResource.BranchGetAllForOwner(userClaims.ID)
	default:
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions."})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Branches not found."})
		return
	}
	ctx.JSON(http.StatusOK, bs.modelResource.BranchToResourceList(branches))
}

func (bs *BranchService) SearchFilter(ctx *gin.Context) {
	filterParam := ctx.Query("filter")
	if filterParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filter parameter is required"})
		return
	}
	userClaims, err := bs.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	var branches filter.FilterPages[models.Branch]
	switch userClaims.AccountType {
	case "Owner":
		branches, err = bs.modelResource.BranchFilterForOwner(filterParam, userClaims.ID)
	case "Admin":
		branches, err = bs.modelResource.BranchFilterForAdmin(filterParam)
	default:
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions."})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Branches not found."})
		return
	}

	data := bs.modelResource.BranchToResourceList(branches.Data)
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

func (as *BranchService) Verify(ctx *gin.Context) {
	userClaims, err := as.getUserClaims(ctx)
	if err != nil {
		as.tokenProvider.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	if userClaims.AccountType != "Admin" {
		as.tokenProvider.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	preloads := ctx.QueryArray("preloads")
	branch := &models.Branch{
		IsAdminVerified: true,
	}
	result, err := as.modelResource.BranchDB.UpdateColumns(uint(id), *branch, preloads)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, as.modelResource.BranchToResource(result))
}

func (bs *BranchService) ExportAll(ctx *gin.Context) {
	userClaims, err := bs.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	var branches []*models.Branch
	switch userClaims.AccountType {
	case "Owner":
		branches, err = bs.modelResource.BranchGetAllForAdmin()
	case "Admin":
		branches, err = bs.modelResource.BranchGetAllForOwner(userClaims.ID)
	default:
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions."})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Branches not found."})
		return
	}
	record, headers := bs.modelResource.BranchToRecord(branches)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("branch-export-all.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}

func (bs *BranchService) ExportAllFiltered(ctx *gin.Context) {
	filterParam := ctx.Query("filter")
	if filterParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filter parameter is required"})
		return
	}
	userClaims, err := bs.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	var branches []*models.Branch
	switch userClaims.AccountType {
	case "Owner":
		branches, err = bs.modelResource.BranchFilterForOwnerRecord(filterParam, userClaims.ID)
	case "Admin":
		branches, err = bs.modelResource.BranchFilterForAdminRecord(filterParam)
	default:
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions."})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Branches not found."})
		return
	}
	record, headers := bs.modelResource.BranchToRecord(branches)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("branches-export-all-filtered.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}
func (bs *BranchService) ExportSelected(ctx *gin.Context) {
	ids, err := bs.helpers.ParseIDsFromQuery(ctx, "ids")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	branch, err := bs.modelResource.BranchDB.GetAllByIDs(ids)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve branches."})
		return
	}
	record, headers := bs.modelResource.BranchToRecord(branch)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("branch-export-all-filtered.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}

func (as *BranchService) ProfilePicture(ctx *gin.Context) {
	var req *models.MediaRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err)})
		return
	}
	if err := as.modelResource.ValidateMediaRequest(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: Validation error: %v", err)})
		return
	}
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	switch claims.AccountType {
	case "Admin":
	case "Owner":
		preloads := ctx.QueryArray("preloads")
		branch := &models.Branch{MediaID: req.ID}
		result, err := as.modelResource.BranchDB.UpdateColumns(uint(id), *branch, preloads)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				ctx.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		ctx.JSON(http.StatusOK, as.modelResource.BranchToResource(result))

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (bs *BranchService) RegisterRoutes() {
	routes := bs.engine.Client.Group("/api/v1/branch")
	routes.Use(bs.middle.AuthMiddleware())
	{
		routes.GET("/search", bs.SearchFilter)
		routes.POST("/", bs.controller.Create)
		routes.GET("/", bs.GetAll)
		routes.GET("/:id", bs.controller.GetByID)
		routes.PUT("/:id", bs.controller.Update)
		routes.DELETE("/:id", bs.controller.Delete)
		routes.DELETE("/bulk-delete", bs.controller.DeleteMany)
		routes.POST("/verify/:id", bs.Verify)
		routes.POST("/profile-picture/:id", bs.ProfilePicture)

		// Export routes
		routes.GET("/export", bs.ExportAll)
		routes.GET("/export-search", bs.ExportAllFiltered)
		routes.GET("/export-selected", bs.ExportSelected)
	}
}
