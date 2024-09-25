package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"horizon-core/internal/models"
	"horizon-core/internal/repository"
	"horizon-core/internal/service"
)

type RoleHandler struct {
	Service *service.RoleService
}

func NewRoleHandler(service *service.RoleService) *RoleHandler {
	return &RoleHandler{
		Service: service,
	}
}

// ListRoles handles GET /roles
func (h *RoleHandler) ListRoles(c *gin.Context) {
	// Parse query parameters
	limitStr := c.DefaultQuery("limit", "10")
	pageStr := c.DefaultQuery("page", "1")
	sortBy := c.Query("sortBy")
	sortOrder := c.Query("sortOrder")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10 // default limit
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1 // default page
	}

	// Build Filters
	filters := []repository.Filter{}
	for key, values := range c.Request.URL.Query() {
		if key == "limit" || key == "page" || key == "sortBy" || key == "sortOrder" {
			continue
		}
		// For simplicity, we'll assume the operator is "equals"
		filters = append(filters, repository.Filter{
			Field:    key,
			Operator: repository.OpEquals,
			Value:    values[0],
		})
	}

	// Build ListRequest
	listRequest := repository.ListRequest{
		Pagination: repository.Pagination{
			Limit:     limit,
			Page:      page,
			SortBy:    sortBy,
			SortOrder: sortOrder,
		},
		Filters: filters,
	}

	// Call service method
	listResponse, err := h.Service.ListRoles(listRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, listResponse)
}

// GetRole handles GET /roles/:id
func (h *RoleHandler) GetRole(c *gin.Context) {
	id := c.Param("id")

	roleResource, err := h.Service.GetRoleByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Write response
	c.JSON(http.StatusOK, roleResource)
}

// CreateRole handles POST /roles
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var roleInput models.Role
	if err := c.ShouldBindJSON(&roleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roleResource, err := h.Service.CreateRole(roleInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusCreated, roleResource)
}

// UpdateRole handles PUT /roles/:id
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id := c.Param("id")

	var roleInput models.Role
	if err := c.ShouldBindJSON(&roleInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure ID consistency
	if roleInput.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID in URL and body do not match"})
		return
	}

	roleResource, err := h.Service.UpdateRole(roleInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, roleResource)
}

// DeleteRole handles DELETE /roles/:id
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteRole(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// RegisterRoleRoutes registers the role routes with the provided router
func RegisterRoleRoutes(router *gin.RouterGroup, db *gorm.DB) {
	roleService := service.NewRoleService(db)
	roleHandler := NewRoleHandler(roleService)

	roleGroup := router.Group("/role")
	{
		roleGroup.GET("", roleHandler.ListRoles)
		roleGroup.GET("/:id", roleHandler.GetRole)
		roleGroup.POST("", roleHandler.CreateRole)
		roleGroup.PUT("/:id", roleHandler.UpdateRole)
		roleGroup.DELETE("/:id", roleHandler.DeleteRole)
	}
}
