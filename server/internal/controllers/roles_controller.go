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

type RoleController struct {
	repo *repositories.RoleRepository
}

func NewRoleController(repo *repositories.RoleRepository) *RoleController {
	return &RoleController{repo: repo}
}

func (c *RoleController) Create(ctx *gin.Context) {
	var req requests.RoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Role instance
	roles := &models.Role{
		Name:               req.Name,
		Description:        req.Description,
		ApiKey:             req.ApiKey,
		Color:              req.Color,
		ReadRole:           req.ReadRole,
		WriteRole:          req.WriteRole,
		UpdateRole:         req.UpdateRole,
		DeleteRole:         req.DeleteRole,
		ReadErrorDetails:   req.ReadErrorDetails,
		WriteErrorDetails:  req.WriteErrorDetails,
		UpdateErrorDetails: req.UpdateErrorDetails,
		DeleteErrorDetails: req.DeleteErrorDetails,
		ReadGender:         req.ReadGender,
		WriteGender:        req.WriteGender,
		UpdateGender:       req.UpdateGender,
		DeleteGender:       req.DeleteGender,
	}

	// Save to the repository
	if err := c.repo.Create(roles); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created error details as a resource
	ctx.JSON(http.StatusCreated, resources.ToResourceRole(roles))
}
func (c *RoleController) GetAll(ctx *gin.Context) {
	roles, err := c.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceListRoles(roles))
}

func (c *RoleController) GetByID(ctx *gin.Context) {
	uid, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	roles, err := c.repo.GetByID(uid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceRole(roles))
}

func (c *RoleController) Update(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	var req requests.RoleRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Role instance
	roles := &models.Role{
		Model:              gorm.Model{ID: id},
		Name:               req.Name,
		Description:        req.Description,
		Color:              req.Color,
		ApiKey:             req.ApiKey,
		ReadRole:           req.ReadRole,
		WriteRole:          req.WriteRole,
		UpdateRole:         req.UpdateRole,
		DeleteRole:         req.DeleteRole,
		ReadErrorDetails:   req.ReadErrorDetails,
		WriteErrorDetails:  req.WriteErrorDetails,
		UpdateErrorDetails: req.UpdateErrorDetails,
		DeleteErrorDetails: req.DeleteErrorDetails,
		ReadGender:         req.ReadGender,
		WriteGender:        req.WriteGender,
		UpdateGender:       req.UpdateGender,
		DeleteGender:       req.DeleteGender,
	}

	if err := c.repo.Update(roles); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceRole(roles))
}

func (c *RoleController) Delete(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}
	if err := c.repo.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func RoleRoutes(router *gin.RouterGroup, controller *RoleController) {
	group := router.Group("/role")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
