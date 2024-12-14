package timesheet

import (
	"fmt"
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
)

type TimesheetService struct {
	db            *providers.DatabaseService
	engine        *providers.EngineService
	models        *models.ModelResource
	middle        *middleware.Middleware
	tokenProvider *providers.TokenService
	modelResource *models.ModelResource
}

func NewTimesheetService(
	db *providers.DatabaseService,
	engine *providers.EngineService,
	models *models.ModelResource,
	middle *middleware.Middleware,
	tokenProvider *providers.TokenService,
	modelResource *models.ModelResource,
) *TimesheetService {

	return &TimesheetService{
		db:            db,
		engine:        engine,
		models:        models,
		middle:        middle,
		tokenProvider: tokenProvider,
		modelResource: modelResource,
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
	userClaims, err := ts.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	timesheet, err := ts.modelResource.TimesheetCurrent(userClaims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found."})
		return
	}
	ctx.JSON(http.StatusOK, ts.modelResource.TimesheetToResource(timesheet))
}
func (ts *TimesheetService) current(ctx *gin.Context) {
	userClaims, err := ts.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	timesheet, err := ts.modelResource.TimesheetCurrent(userClaims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found."})
		return
	}
	ctx.JSON(http.StatusOK, ts.modelResource.TimesheetToResource(timesheet))
}
func (ts *TimesheetService) timein(ctx *gin.Context) {
	// userClaims, err := ts.getUserClaims(ctx)
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
	// 	return
	// }
}
func (ts *TimesheetService) timeout(ctx *gin.Context) {
	// userClaims, err := ts.getUserClaims(ctx)
	// if err != nil {
	// 	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
	// 	return
	// }
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
