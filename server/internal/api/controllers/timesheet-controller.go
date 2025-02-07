package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type TimesheetController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewTimesheetController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *TimesheetController {
	return &TimesheetController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

// GET /api/v1/timesheet
func (c *TimesheetController) Index(ctx *gin.Context) {}

// GET /api/v1/current
func (c *TimesheetController) Current(ctx *gin.Context) {}

// GET /api/v1/timein
func (c *TimesheetController) TimeIn(ctx *gin.Context) {}

// GET /api/v1/timout
func (c *TimesheetController) TimeOut(ctx *gin.Context) {}
