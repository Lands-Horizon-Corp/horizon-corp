package company

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CompanyService struct {
	controller    *managers.Controller[models.Company, models.CompanyRequest, models.CompanyResource]
	middle        *middleware.Middleware
	db            *providers.DatabaseService
	engine        *providers.EngineService
	models        *models.ModelResource
	tokenProvider *providers.TokenService
	helpers       *helpers.HelpersFunction
	modelResource *models.ModelResource
}

func NewCompanyService(
	db *providers.DatabaseService,
	middle *middleware.Middleware,
	engine *providers.EngineService,
	models *models.ModelResource,
	tokenProvider *providers.TokenService,
	helpers *helpers.HelpersFunction,
	modelResource *models.ModelResource,
) *CompanyService {
	controller := managers.NewController(
		models.CompanyDB,
		models.ValidateCompanyRequest,
		models.CompanyToResource,
		models.CompanyToResourceList,
	)

	return &CompanyService{
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

func (ts *CompanyService) getUserClaims(ctx *gin.Context) (*providers.UserClaims, error) {
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

func (as *CompanyService) SearchFilter(ctx *gin.Context) {
	filterParam := ctx.Query("filter")
	if filterParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filter parameter is required"})
		return
	}

	companies, err := as.modelResource.CompanyFilterForAdmin(filterParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Companies not found."})
		return
	}
	data := as.modelResource.CompanyToResourceList(companies.Data)
	if data == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data":      []interface{}{},
			"pageIndex": companies.PageIndex,
			"totalPage": companies.TotalPage,
			"pageSize":  companies.PageSize,
			"totalSize": companies.TotalSize,
			"pages":     companies.Pages,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":      data,
		"pageIndex": companies.PageIndex,
		"totalPage": companies.TotalPage,
		"pageSize":  companies.PageSize,
		"totalSize": companies.TotalSize,
		"pages":     companies.Pages,
	})
}

func (as *CompanyService) Verify(ctx *gin.Context) {

	userClaims, err := as.getUserClaims(ctx)
	if err != nil {
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
	company := &models.Company{
		IsAdminVerified: true,
	}
	result, err := as.modelResource.CompanyDB.UpdateColumns(uint(id), *company, preloads)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusOK, as.modelResource.CompanyToResource(result))
}

func (as *CompanyService) ExportAll(ctx *gin.Context) {
	userClaims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	if userClaims.AccountType != "Admin" {
		as.tokenProvider.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	company, err := as.modelResource.CompanyDB.FindAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to export CSV: %v", err)
		return
	}
	record, headers := as.modelResource.CompanyToRecord(company)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("company-export-all.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}

func (as *CompanyService) ExportAllFiltered(ctx *gin.Context) {
	userClaims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	if userClaims.AccountType != "Admin" {
		as.tokenProvider.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	filterParam := ctx.Query("filter")
	if filterParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filter parameter is required"})
		return
	}
	company, err := as.modelResource.CompanyFilterForAdminRecord(filterParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Companies not found."})
		return
	}
	record, headers := as.modelResource.CompanyToRecord(company)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("company-export-all-filtered.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}

func (as *CompanyService) ExportSelected(ctx *gin.Context) {
	userClaims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	if userClaims.AccountType != "Admin" {
		as.tokenProvider.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	ids, err := as.helpers.ParseIDsFromQuery(ctx, "ids")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	company, err := as.modelResource.CompanyDB.GetAllByIDs(ids)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve companies."})
		return
	}
	record, headers := as.modelResource.CompanyToRecord(company)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("company-export-all-filtered.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}

func (as *CompanyService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/company")
	routes.Use(as.middle.AuthMiddleware())
	{
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)

		routes.POST("/verify/:id", as.middle.AuthMiddlewareAdminOnly(), as.Verify)
		routes.GET("/search", as.middle.AuthMiddlewareAdminOnly(), as.SearchFilter)
		routes.DELETE("/bulk-delete", as.middle.AuthMiddlewareAdminOnly(), as.controller.DeleteMany)
		routes.DELETE("/:id", as.middle.AuthMiddlewareAdminOnly(), as.controller.Delete)
		// Export routes
		routes.GET("/export", as.ExportAll)
		routes.GET("/export-search", as.ExportAllFiltered)
		routes.GET("/export-selected", as.ExportSelected)
	}
}
