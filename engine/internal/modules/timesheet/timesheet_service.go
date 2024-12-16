package timesheet

import (
	"fmt"
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers/filter"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TimesheetService struct {
	db            *providers.DatabaseService
	engine        *providers.EngineService
	models        *models.ModelResource
	middle        *middleware.Middleware
	tokenProvider *providers.TokenService
	modelResource *models.ModelResource
	helpers       *helpers.HelpersFunction
}

func NewTimesheetService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
	middle *middleware.Middleware,
	tokenProvider *providers.TokenService,
	modelResource *models.ModelResource,
	helpers *helpers.HelpersFunction,
) *TimesheetService {

	return &TimesheetService{
		db:            db,
		engine:        engine,
		models:        models,
		middle:        middle,
		tokenProvider: tokenProvider,
		modelResource: modelResource,
		helpers:       helpers,
	}
}

func (ts *TimesheetService) getUserClaims(ctx *gin.Context) (*providers.UserClaims, error) {
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

func (ts *TimesheetService) findall(ctx *gin.Context) {

	filterParam := ctx.Query("filter")
	if filterParam == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "filter parameter is required"})
		return
	}

	var paginatedReq filter.PaginatedRequest
	if err := ts.helpers.DecodeBase64JSON(filterParam, &paginatedReq); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid filter parameter"})
		return
	}

	if err := ctx.ShouldBind(&paginatedReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	userClaims, err := ts.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Companies not authenticated."})
		return
	}
	timesheet, err := ts.modelResource.TimesheetFindallForEmployee(userClaims.ID, paginatedReq)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Companies not found."})
		return
	}
	data := ts.modelResource.TimesheetToResourceList(timesheet.Data)
	ctx.JSON(http.StatusOK, gin.H{
		"data":      data,
		"pageIndex": timesheet.PageIndex,
		"totalPage": timesheet.TotalPage,
		"pageSize":  timesheet.PageSize,
		"totalSize": timesheet.TotalSize,
		"pages":     timesheet.Pages,
	})
}

func (ts *TimesheetService) current(ctx *gin.Context) {
	userClaims, err := ts.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	timesheet, err := ts.modelResource.TimesheetCurrentForEmployee(userClaims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found."})
		return
	}
	ctx.JSON(http.StatusOK, ts.modelResource.TimesheetToResource(timesheet))
}
func (ts *TimesheetService) timein(ctx *gin.Context) {
	userClaims, err := ts.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	var req models.TimeInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if err := ts.models.ValidateTimeInRequest(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingOpen models.Timesheet
	err = ts.modelResource.EmployeeDB.DB.Client.
		Where("employee_id = ? AND time_out IS NULL", userClaims.ID).
		First(&existingOpen).Error
	if err == nil {
		ctx.JSON(http.StatusConflict, gin.H{"error": "You already have an open timesheet"})
		return
	}
	if err != nil && err != gorm.ErrRecordNotFound {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not check existing timesheets"})
		return
	}

	newTimesheet := &models.Timesheet{
		EmployeeID: userClaims.ID,
		TimeIn:     &req.TimeIn,
		MediaInID:  req.MediaIn.ID,
	}

	if err := ts.modelResource.TimesheetDB.Create(newTimesheet).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not create timesheet"})
		return
	}

	ctx.JSON(http.StatusCreated, ts.modelResource.TimesheetToResource(newTimesheet))
}

func (ts *TimesheetService) timeout(ctx *gin.Context) {
	userClaims, err := ts.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	var req models.TimeOutRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	if err := ts.models.ValidateTimeOutRequest(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var currentTS models.Timesheet
	err = ts.modelResource.EmployeeDB.DB.Client.
		Where("employee_id = ? AND time_out IS NULL", userClaims.ID).
		First(&currentTS).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "no open timesheet to close"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve current timesheet"})
		return
	}

	currentTS.TimeOut = &req.TimeOut
	currentTS.MediaOutID = req.MediaOut.ID

	if err := ts.modelResource.TimesheetDB.Update(&currentTS, []string{}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not update timesheet"})
		return
	}
	ctx.JSON(http.StatusCreated, ts.modelResource.TimesheetToResource(&currentTS))

}

func (ts *TimesheetService) RegisterRoutes() {
	routes := ts.engine.Client.Group("/api/v1/timesheet")
	{
		routes.Use(ts.middle.AuthMiddleware())
		routes.GET("/", ts.findall)
		routes.GET("/current", ts.current)
		routes.POST("/time-in", ts.timein)
		routes.POST("/time-out", ts.timeout)

	}
}
