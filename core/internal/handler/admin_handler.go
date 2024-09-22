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

type AdminHandler struct {
	Service *service.AdminService
}

func NewAdminHandler(service *service.AdminService) *AdminHandler {
	return &AdminHandler{
		Service: service,
	}
}

// ListAdmins handles GET /admins
func (h *AdminHandler) ListAdmins(c *gin.Context) {
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
	listResponse, err := h.Service.ListAdmins(listRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, listResponse)
}

// GetAdmin handles GET /admins/:id
func (h *AdminHandler) GetAdmin(c *gin.Context) {
	id := c.Param("id")

	adminResource, err := h.Service.GetAdminByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Write response
	c.JSON(http.StatusOK, adminResource)
}

// CreateAdmin handles POST /admins
func (h *AdminHandler) CreateAdmin(c *gin.Context) {
	var adminInput models.Admin
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	adminResource, err := h.Service.CreateAdmin(adminInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusCreated, adminResource)
}

// UpdateAdmin handles PUT /admins/:id
func (h *AdminHandler) UpdateAdmin(c *gin.Context) {
	id := c.Param("id")

	var adminInput models.Admin
	if err := c.ShouldBindJSON(&adminInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure ID consistency
	if adminInput.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID in URL and body do not match"})
		return
	}

	adminResource, err := h.Service.UpdateAdmin(adminInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, adminResource)
}

// DeleteAdmin handles DELETE /admins/:id
func (h *AdminHandler) DeleteAdmin(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteAdmin(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// RegisterAdminRoutes registers the admin routes with the provided router
func RegisterAdminRoutes(router *gin.RouterGroup, db *gorm.DB) {
	adminService := service.NewAdminService(db)
	adminHandler := NewAdminHandler(adminService)

	adminGroup := router.Group("/admins")
	{
		adminGroup.GET("", adminHandler.ListAdmins)
		adminGroup.GET("/:id", adminHandler.GetAdmin)
		adminGroup.POST("", adminHandler.CreateAdmin)
		adminGroup.PUT("/:id", adminHandler.UpdateAdmin)
		adminGroup.DELETE("/:id", adminHandler.DeleteAdmin)
	}
}
