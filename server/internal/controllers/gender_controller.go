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

type GenderController struct {
	repo *repositories.GenderRepository
}

func NewGenderController(repo *repositories.GenderRepository) *GenderController {
	return &GenderController{repo: repo}
}

func (c *GenderController) Create(ctx *gin.Context) {
	var req requests.GenderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gender := &models.Gender{
		Name:        req.Name,
		Description: req.Description,
	}

	if err := c.repo.Create(gender); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, resources.ToResourceGender(gender))
}

func (c *GenderController) GetAll(ctx *gin.Context) {
	genders, err := c.repo.GetAll([]string{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceListGender(genders))
}

func (c *GenderController) GetByID(ctx *gin.Context) {
	uid, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	gender, err := c.repo.GetByID(uid, []string{})
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Gender not found"})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceGender(gender))
}

func (c *GenderController) Update(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	var req requests.GenderRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	gender := &models.Gender{
		Model:       gorm.Model{ID: id},
		Name:        req.Name,
		Description: req.Description,
	}

	if err := c.repo.Update(gender); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceGender(gender))
}

func (c *GenderController) Delete(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}
	if err := c.repo.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Gender not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func GenderRoutes(router *gin.RouterGroup, controller *GenderController) {
	group := router.Group("/gender")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
