package controllers

import (
	"horizon/server/helpers"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FeedbackController struct {
	repo *repositories.FeedbackRepository
}

func NewFeedbackController(repo *repositories.FeedbackRepository) *FeedbackController {
	return &FeedbackController{repo: repo}
}

func (c *FeedbackController) Create(ctx *gin.Context) {
	var req requests.FeedbackRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feedback := models.Feedback{
		Email:        req.Email,
		Description:  req.Description,
		FeedbackType: req.FeedbackType,
	}

	if err := c.repo.Create(&feedback); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resources.ToResourceFeedback(&feedback))
}

func (c *FeedbackController) GetAll(ctx *gin.Context) {
	feedbacks, err := c.repo.GetAll([]string{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceListFeedback(feedbacks))
}

func (c *FeedbackController) GetByID(ctx *gin.Context) {
	uid, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	feedback, err := c.repo.GetByID(uid, []string{})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceFeedback(feedback))
}

func (c *FeedbackController) Update(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	var req requests.FeedbackRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feedback := &models.Feedback{
		Model:        gorm.Model{ID: id},
		Email:        req.Email,
		Description:  req.Description,
		FeedbackType: req.FeedbackType,
	}

	if err := c.repo.Update(feedback, []string{}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceFeedback(feedback))
}

func (c *FeedbackController) Delete(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}
	if err := c.repo.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Feedback not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func FeedbackRoutes(router *gin.RouterGroup, controller *FeedbackController) {
	group := router.Group("/feedback")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
