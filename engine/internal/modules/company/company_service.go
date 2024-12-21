package company

import (
	"fmt"
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
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

func (as *CompanyService) ExportAll(ctx *gin.Context) {
	fmt.Println("Export all")
}

func (as *CompanyService) ExportAllFiltered(ctx *gin.Context) {
	fmt.Println("Export all filtered")
}

func (as *CompanyService) ExportSelected(ctx *gin.Context) {
	fmt.Println("Export all selected")
}

func (as *CompanyService) ExportCurrentPage(ctx *gin.Context) {
	fmt.Println("Export all current page")
}

func (as *CompanyService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/company")
	routes.Use(as.middle.AuthMiddleware())
	{
		routes.GET("/search", as.SearchFilter)
		routes.POST("/", as.controller.Create)
		routes.GET("/", as.controller.GetAll)
		routes.GET("/:id", as.controller.GetByID)
		routes.PUT("/:id", as.controller.Update)
		routes.DELETE("/:id", as.controller.Delete)

		// Export routes
		routes.GET("/export", as.ExportAll)
		routes.GET("/export-search", as.ExportAllFiltered)
		routes.POST("/export-selected", as.ExportSelected)
		routes.GET("/export-current-page/:page", as.ExportCurrentPage)
	}
}
