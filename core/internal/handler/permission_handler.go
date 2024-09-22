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

type PermissionHandler struct {
	Service *service.PermissionService
}

func NewPermissionHandler(service *service.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		Service: service,
	}
}

// ListPermissions handles GET /permissions
func (h *PermissionHandler) ListPermissions(c *gin.Context) {
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
	listResponse, err := h.Service.ListPermissions(listRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, listResponse)
}

// GetPermission handles GET /permissions/:id
func (h *PermissionHandler) GetPermission(c *gin.Context) {
	id := c.Param("id")

	permissionResource, err := h.Service.GetPermissionByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Write response
	c.JSON(http.StatusOK, permissionResource)
}

// CreatePermission handles POST /permissions
func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	var permissionInput models.Permission
	if err := c.ShouldBindJSON(&permissionInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permissionResource, err := h.Service.CreatePermission(permissionInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusCreated, permissionResource)
}

// UpdatePermission handles PUT /permissions/:id
func (h *PermissionHandler) UpdatePermission(c *gin.Context) {
	id := c.Param("id")

	var permissionInput models.Permission
	if err := c.ShouldBindJSON(&permissionInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure ID consistency
	if permissionInput.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID in URL and body do not match"})
		return
	}

	permissionResource, err := h.Service.UpdatePermission(permissionInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, permissionResource)
}

// DeletePermission handles DELETE /permissions/:id
func (h *PermissionHandler) DeletePermission(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeletePermission(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Permission not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// RegisterPermissionRoutes registers the permission routes with the provided router
func RegisterPermissionRoutes(router *gin.RouterGroup, db *gorm.DB) {
	permissionService := service.NewPermissionService(db)
	permissionHandler := NewPermissionHandler(permissionService)

	permissionGroup := router.Group("/permissions")
	{
		permissionGroup.GET("", permissionHandler.ListPermissions)
		permissionGroup.GET("/:id", permissionHandler.GetPermission)
		permissionGroup.POST("", permissionHandler.CreatePermission)
		permissionGroup.PUT("/:id", permissionHandler.UpdatePermission)
		permissionGroup.DELETE("/:id", permissionHandler.DeletePermission)
	}
}
