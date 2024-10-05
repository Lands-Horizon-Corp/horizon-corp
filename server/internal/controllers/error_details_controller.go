package controllers

import (
	"horizon/server/helpers"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorDetailsController struct {
	repo *repositories.ErrorDetailsRepository
}

func NewErrorDetailsController(repo *repositories.ErrorDetailsRepository) *ErrorDetailsController {
	return &ErrorDetailsController{repo: repo}
}

func (c *ErrorDetailsController) Create(ctx *gin.Context) {
	var req requests.ErrorDetailsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert Stack and Response to JSON string if present
	var stackJSON, responseJSON *string
	if req.Stack != "" {
		stack := helpers.JSONStringify(req.Stack)
		stackJSON = &stack
	}
	if req.Response != "" {
		response := helpers.JSONStringify(req.Response)
		responseJSON = &response
	}

	// Create ErrorDetails instance
	errorDetails := models.ErrorDetails{
		Message:  req.Message,
		Name:     req.Name,
		Stack:    stackJSON,
		Response: responseJSON,
		Status:   &req.Status,
	}

	// Save to the repository
	if err := c.repo.Create(&errorDetails); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created error details as a resource
	ctx.JSON(http.StatusCreated, resources.ToResourceErrorDetails(errorDetails))
}
func (c *ErrorDetailsController) GetAll(ctx *gin.Context) {
	errorDetail, err := c.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceListErrorDetails(errorDetail))
}

func (c *ErrorDetailsController) GetByID(ctx *gin.Context) {
	uid, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	errorDetails, err := c.repo.GetByID(uid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ErrorDetails not found"})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceErrorDetails(errorDetails))
}

func (c *ErrorDetailsController) Update(ctx *gin.Context) {
	var req requests.ErrorDetailsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var stackJSON, responseJSON *string
	if req.Stack != "" {
		stack := helpers.JSONStringify(req.Stack)
		stackJSON = &stack
	}
	if req.Response != "" {
		response := helpers.JSONStringify(req.Response)
		responseJSON = &response
	}

	// Create ErrorDetails instance
	errorDetails := models.ErrorDetails{
		Message:  req.Message,
		Name:     req.Name,
		Stack:    stackJSON,
		Response: responseJSON,
		Status:   &req.Status,
	}

	if err := c.repo.Update(&errorDetails); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceErrorDetails(errorDetails))
}

func (c *ErrorDetailsController) Delete(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}
	if err := c.repo.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ErrorDetails not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func ErrorDetailsRoutes(router *gin.RouterGroup, controller *ErrorDetailsController) {
	group := router.Group("/error-details")
	{
		group.POST("/", controller.Create)
		group.GET("/", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
