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

type OwnerHandler struct {
	Service *service.OwnerService
}

func NewOwnerHandler(service *service.OwnerService) *OwnerHandler {
	return &OwnerHandler{
		Service: service,
	}
}

// ListOwners handles GET /owners
func (h *OwnerHandler) ListOwners(c *gin.Context) {
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
	listResponse, err := h.Service.ListOwners(listRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, listResponse)
}

// GetOwner handles GET /owners/:id
func (h *OwnerHandler) GetOwner(c *gin.Context) {
	id := c.Param("id")

	ownerResource, err := h.Service.GetOwnerByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Owner not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Write response
	c.JSON(http.StatusOK, ownerResource)
}

// CreateOwner handles POST /owners
func (h *OwnerHandler) CreateOwner(c *gin.Context) {
	var ownerInput models.Owner
	if err := c.ShouldBindJSON(&ownerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ownerResource, err := h.Service.CreateOwner(ownerInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusCreated, ownerResource)
}

// UpdateOwner handles PUT /owners/:id
func (h *OwnerHandler) UpdateOwner(c *gin.Context) {
	id := c.Param("id")

	var ownerInput models.Owner
	if err := c.ShouldBindJSON(&ownerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure ID consistency
	if ownerInput.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID in URL and body do not match"})
		return
	}

	ownerResource, err := h.Service.UpdateOwner(ownerInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, ownerResource)
}

// DeleteOwner handles DELETE /owners/:id
func (h *OwnerHandler) DeleteOwner(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteOwner(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Owner not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// RegisterOwnerRoutes registers the owner routes with the provided router
func RegisterOwnerRoutes(router *gin.RouterGroup, db *gorm.DB) {
	ownerService := service.NewOwnerService(db)
	ownerHandler := NewOwnerHandler(ownerService)

	ownerGroup := router.Group("/owner")
	{
		ownerGroup.GET("", ownerHandler.ListOwners)
		ownerGroup.GET("/:id", ownerHandler.GetOwner)
		ownerGroup.POST("", ownerHandler.CreateOwner)
		ownerGroup.PUT("/:id", ownerHandler.UpdateOwner)
		ownerGroup.DELETE("/:id", ownerHandler.DeleteOwner)
	}
}
