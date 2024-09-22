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

type BranchHandler struct {
	Service *service.BranchService
}

func NewBranchHandler(service *service.BranchService) *BranchHandler {
	return &BranchHandler{
		Service: service,
	}
}

// ListBranchs handles GET /branchs
func (h *BranchHandler) ListBranchs(c *gin.Context) {
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
	listResponse, err := h.Service.ListBranches(listRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, listResponse)
}

// GetBranch handles GET /branchs/:id
func (h *BranchHandler) GetBranch(c *gin.Context) {
	id := c.Param("id")

	branchResource, err := h.Service.GetBranchByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Write response
	c.JSON(http.StatusOK, branchResource)
}

// CreateBranch handles POST /branchs
func (h *BranchHandler) CreateBranch(c *gin.Context) {
	var branchInput models.Branch
	if err := c.ShouldBindJSON(&branchInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branchResource, err := h.Service.CreateBranch(branchInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusCreated, branchResource)
}

// UpdateBranch handles PUT /branchs/:id
func (h *BranchHandler) UpdateBranch(c *gin.Context) {
	id := c.Param("id")

	var branchInput models.Branch
	if err := c.ShouldBindJSON(&branchInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure ID consistency
	if branchInput.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID in URL and body do not match"})
		return
	}

	branchResource, err := h.Service.UpdateBranch(branchInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, branchResource)
}

// DeleteBranch handles DELETE /branchs/:id
func (h *BranchHandler) DeleteBranch(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteBranch(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Branch not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// RegisterBranchRoutes registers the branch routes with the provided router
func RegisterBranchRoutes(router *gin.RouterGroup, db *gorm.DB) {
	branchService := service.NewBranchService(db)
	branchHandler := NewBranchHandler(branchService)

	branchGroup := router.Group("/branch")
	{
		branchGroup.GET("", branchHandler.ListBranchs)
		branchGroup.GET("/:id", branchHandler.GetBranch)
		branchGroup.POST("", branchHandler.CreateBranch)
		branchGroup.PUT("/:id", branchHandler.UpdateBranch)
		branchGroup.DELETE("/:id", branchHandler.DeleteBranch)
	}
}
