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

type ErrorDetailController struct {
	repo *repositories.ErrorDetailRepository
}

func NewErrorDetailController(repo *repositories.ErrorDetailRepository) *ErrorDetailController {
	return &ErrorDetailController{repo: repo}
}

func (c *ErrorDetailController) Create(ctx *gin.Context) {
	var req requests.ErrorDetailRequest
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

	// Create ErrorDetail instance
	errorDetails := models.ErrorDetail{
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
	ctx.JSON(http.StatusCreated, resources.ToResourceErrorDetail(errorDetails))
}
func (c *ErrorDetailController) GetAll(ctx *gin.Context) {
	errorDetail, err := c.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceListErrorDetail(errorDetail))
}

func (c *ErrorDetailController) GetByID(ctx *gin.Context) {
	uid, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	errorDetails, err := c.repo.GetByID(uid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ErrorDetail not found"})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceErrorDetail(errorDetails))
}

func (c *ErrorDetailController) Update(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	var req requests.ErrorDetailRequest
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

	// Create ErrorDetail instance
	errorDetail := models.ErrorDetail{
		Message:  req.Message,
		Name:     req.Name,
		Stack:    stackJSON,
		Response: responseJSON,
		Status:   &req.Status,
	}

	if err := c.repo.Update(id, &errorDetail); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceErrorDetail(errorDetail))
}

func (c *ErrorDetailController) Delete(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}
	if err := c.repo.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "ErrorDetail not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func ErrorDetailRoutes(router *gin.RouterGroup, controller *ErrorDetailController) {
	group := router.Group("/error-details")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
