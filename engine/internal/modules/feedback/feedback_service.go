package feedback

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
)

type FeedbackService struct {
	controller    *managers.Controller[models.Feedback, models.FeedbackRequest, models.FeedbackResource]
	middle        *middleware.Middleware
	db            *providers.DatabaseService
	engine        *providers.EngineService
	models        *models.ModelResource
	tokenProvider *providers.TokenService
	helpers       *helpers.HelpersFunction
	modelResource *models.ModelResource
}

func NewFeedbackService(
	db *providers.DatabaseService,
	middle *middleware.Middleware,
	engine *providers.EngineService,
	models *models.ModelResource,
	tokenProvider *providers.TokenService,
	helpers *helpers.HelpersFunction,
	modelResource *models.ModelResource,
) *FeedbackService {
	controller := managers.NewController(
		models.FeedbackDB,
		models.ValidateFeedbackRequest,
		models.FeedbackToResource,
		models.FeedbackToResourceList,
	)

	return &FeedbackService{
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

func (ts *FeedbackService) getUserClaims(ctx *gin.Context) (*providers.UserClaims, error) {
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

// # Admin only
func (fs *FeedbackService) SearchFilter(ctx *gin.Context) {
	filterParam := ctx.Query("filter")
	if filterParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filter parameter is required"})
		return
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
	feedbacks, err := fs.modelResource.FeedbackFilterForAdmin(filterParam, pageIndex, pageSize)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Feedbacks not found."})
		return
	}
	data := fs.modelResource.FeedbackToResourceList(feedbacks.Data)
	if data == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"data":      []interface{}{},
			"pageIndex": feedbacks.PageIndex,
			"totalPage": feedbacks.TotalPage,
			"pageSize":  feedbacks.PageSize,
			"totalSize": feedbacks.TotalSize,
			"pages":     feedbacks.Pages,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":      data,
		"pageIndex": feedbacks.PageIndex,
		"totalPage": feedbacks.TotalPage,
		"pageSize":  feedbacks.PageSize,
		"totalSize": feedbacks.TotalSize,
		"pages":     feedbacks.Pages,
	})
}
func (fs *FeedbackService) ExportAll(ctx *gin.Context) {
	feedback, err := fs.modelResource.FeedbackDB.FindAll()
	if err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to export CSV: %v", err)
		return
	}
	record, headers := fs.modelResource.FeedbackToRecord(feedback)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("feedback-export-all.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}

func (fs *FeedbackService) ExportAllFiltered(ctx *gin.Context) {
	filterParam := ctx.Query("filter")
	if filterParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filter parameter is required"})
		return
	}
	feedback, err := fs.modelResource.FeedbackFilterForAdminRecord(filterParam)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Companies not found."})
		return
	}
	record, headers := fs.modelResource.FeedbackToRecord(feedback)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("feedback-export-all-filtered.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}
func (fs *FeedbackService) ExportSelected(ctx *gin.Context) {
	ids, err := fs.helpers.ParseIDsFromQuery(ctx, "ids")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	feedbacks, err := fs.modelResource.FeedbackDB.GetAllByIDs(ids)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve companies."})
		return
	}
	record, headers := fs.modelResource.FeedbackToRecord(feedbacks)
	csvManager := managers.NewCSVManager()
	csvManager.SetFileName("feedback-export-all-filtered.csv")
	csvManager.SetHeaders(headers)
	csvManager.AddRecords(record)
	if err := csvManager.WriteCSV(ctx); err != nil {
		ctx.String(http.StatusInternalServerError, "Failed to generate CSV: %v", err)
		return
	}
}

func (as *FeedbackService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/feedback")
	{
		// Public
		routes.POST("/", as.controller.Create)

		// Admin only
		routes.GET("/search", as.middle.AuthMiddlewareAdminOnly(), as.SearchFilter)
		routes.GET("/", as.middle.AuthMiddlewareAdminOnly(), as.controller.GetAll)
		routes.GET("/:id", as.middle.AuthMiddlewareAdminOnly(), as.controller.GetByID)
		routes.DELETE("/:id", as.middle.AuthMiddlewareAdminOnly(), as.controller.Delete)
		routes.DELETE("/bulk-delete", as.middle.AuthMiddlewareAdminOnly(), as.controller.DeleteMany)

		// Export routes
		routes.GET("/export", as.middle.AuthMiddlewareAdminOnly(), as.ExportAll)
		routes.GET("/export-search", as.middle.AuthMiddlewareAdminOnly(), as.ExportAllFiltered)
		routes.GET("/export-selected", as.middle.AuthMiddlewareAdminOnly(), as.ExportSelected)
	}
}
