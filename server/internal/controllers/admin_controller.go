package controllers

import (
	"horizon/server/config"
	"horizon/server/helpers"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/resources"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	repo *repositories.AdminRepository
}

func NewAdminController(repo *repositories.AdminRepository) *AdminController {
	return &AdminController{repo: repo}
}

func (c *AdminController) Create(ctx *gin.Context) {
	var req requests.AdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := config.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	admin := models.Admin{
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		PermanentAddress:  req.PermanentAddress,
		Description:       req.Description,
		Birthdate:         req.Birthdate,
		Username:          req.Username,
		Email:             req.Email,
		Password:          hashedPassword,
		MediaID:           req.Media.ID,
		IsEmailVerified:   false,
		IsContactVerified: false,
	}
	if err := c.repo.Create(&admin); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, resources.ToResourceAdmin(admin))
}

func (c *AdminController) GetAll(ctx *gin.Context) {
	admins, err := c.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceListAdmin(admins))
}

func (c *AdminController) GetByID(ctx *gin.Context) {
	uid, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	admin, err := c.repo.GetByID(uid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceAdmin(admin))
}

func (c *AdminController) Update(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	var req requests.AdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	admin := models.Admin{
		// what to put here
	}

	if err := c.repo.Update(id, &admin); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceAdmin(admin))
}

func (c *AdminController) Delete(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}
	if err := c.repo.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func AdminRoutes(router *gin.RouterGroup, controller *AdminController) {
	group := router.Group("/admin")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
