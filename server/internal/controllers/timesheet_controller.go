package controllers

import (
	"fmt"
	"horizon/server/internal/auth"
	"horizon/server/internal/middleware"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TimesheetController struct {
	timesheetRepo *repositories.TimesheetRepository
	tokenService  auth.TokenService
}

func NewTimesheetController(
	timesheetRepo *repositories.TimesheetRepository,
	tokenService auth.TokenService,
) *TimesheetController {
	return &TimesheetController{
		timesheetRepo: timesheetRepo,
		tokenService:  tokenService,
	}
}

func (c *TimesheetController) getUserClaims(ctx *gin.Context) (*auth.UserClaims, error) {
	claims, exists := ctx.Get("claims")
	if !exists {
		c.tokenService.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("claims not found in context")
	}

	userClaims, ok := claims.(*auth.UserClaims)
	if !ok || userClaims.AccountType != "Employee" {
		c.tokenService.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("unauthorized user")
	}
	return userClaims, nil
}

func (c *TimesheetController) CurrentEmployeeTime(ctx *gin.Context) {
	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	employeeID := userClaims.ID
	timesheet, err := c.timesheetRepo.GetLatestForEmployee(employeeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve timesheet"})
		return
	}

	if timesheet.ID == 0 {
		ctx.JSON(http.StatusOK, gin.H{"timesheet": nil})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceTimesheet(&timesheet))
}

func (c *TimesheetController) TimeIn(ctx *gin.Context) {
	var timeInReq requests.TimeInRequest
	if err := ctx.ShouldBindJSON(&timeInReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := timeInReq.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	employeeID := userClaims.ID
	latestTimesheet, err := c.timesheetRepo.GetLatestForEmployee(employeeID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check latest timesheet"})
		return
	}

	if latestTimesheet.TimeOut == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Already timed in, please time out first"})
		return
	}

	timesheet := models.Timesheet{
		EmployeeID: employeeID,
		TimeIn:     &timeInReq.TimeIn,
		MediaInID:  &timeInReq.MediaIn.ID,
	}

	if err := c.timesheetRepo.Create(&timesheet); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record time in"})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceTimesheet(&timesheet))
}

func (c *TimesheetController) TimeOut(ctx *gin.Context) {
	var timeOutReq requests.TimeOutRequest
	if err := ctx.ShouldBindJSON(&timeOutReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := timeOutReq.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	employeeID := userClaims.ID
	timesheet, err := c.timesheetRepo.GetLatestForEmployee(employeeID)
	if err != nil || timesheet.TimeOut != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No active time in record found"})
		return
	}

	timesheet.TimeOut = &timeOutReq.TimeOut
	timesheet.MediaOutID = &timeOutReq.MediaOut.ID

	if err := c.timesheetRepo.Update(&timesheet); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to record time out"})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceTimesheet(&timesheet))
}

func TimesheetRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware, controller *TimesheetController) {
	authGroup := router.Group("/timesheet")
	authGroup.Use(mw.Middleware())
	{
		authGroup.GET("/current", controller.CurrentEmployeeTime)
		authGroup.POST("/time-in", controller.TimeIn)
		authGroup.POST("/time-out", controller.TimeOut)
	}
}
