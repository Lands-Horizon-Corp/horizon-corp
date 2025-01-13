package controllers

import "github.com/gin-gonic/gin"

type TimesheetController struct{}

func NewTimesheetController() *TimesheetController {
	return &TimesheetController{}
}

// GET /api/v1/timesheet
func (c *TimesheetController) Index(ctx *gin.Context) {}

// GET /api/v1/current
func (c *TimesheetController) Current(ctx *gin.Context) {}

// GET /api/v1/timein
func (c *TimesheetController) TimeIn(ctx *gin.Context) {}

// GET /api/v1/timout
func (c *TimesheetController) TimeOut(ctx *gin.Context) {}
