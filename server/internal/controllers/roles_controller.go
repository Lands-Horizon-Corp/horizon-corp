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

type RolesController struct {
	repo *repositories.RolesRepository
}

func NewRolesController(repo *repositories.RolesRepository) *RolesController {
	return &RolesController{repo: repo}
}

func (c *RolesController) Create(ctx *gin.Context) {
	var req requests.RolesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Roles instance
	roles := models.Roles{
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
	if err := c.repo.Create(&roles); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created error details as a resource
	ctx.JSON(http.StatusCreated, resources.ToResourceRoles(roles))
}
func (c *RolesController) GetAll(ctx *gin.Context) {
	errorDetail, err := c.repo.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceListRoles(errorDetail))
}

func (c *RolesController) GetByID(ctx *gin.Context) {
	uid, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}

	roles, err := c.repo.GetByID(uid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Roles not found"})
		return
	}
	ctx.JSON(http.StatusOK, resources.ToResourceRoles(roles))
}

func (c *RolesController) Update(ctx *gin.Context) {
	var req requests.RolesRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create Roles instance
	roles := models.Roles{
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

	if err := c.repo.Update(&roles); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceRoles(roles))
}

func (c *RolesController) Delete(ctx *gin.Context) {
	id, err := helpers.ParseIDParam(ctx, "id")
	if err != nil {
		return
	}
	if err := c.repo.Delete(id); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Roles not found"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func RolesRoutes(router *gin.RouterGroup, controller *RolesController) {
	group := router.Group("/roles")
	{
		group.POST("", controller.Create)
		group.GET("", controller.GetAll)
		group.GET("/:id", controller.GetByID)
		group.PUT("/:id", controller.Update)
		group.DELETE("/:id", controller.Delete)
	}
}
