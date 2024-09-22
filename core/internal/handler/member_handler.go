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

type MemberHandler struct {
	Service *service.MemberService
}

func NewMemberHandler(service *service.MemberService) *MemberHandler {
	return &MemberHandler{
		Service: service,
	}
}

// ListMembers handles GET /members
func (h *MemberHandler) ListMembers(c *gin.Context) {
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
	listResponse, err := h.Service.ListMembers(listRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, listResponse)
}

// GetMember handles GET /members/:id
func (h *MemberHandler) GetMember(c *gin.Context) {
	id := c.Param("id")

	memberResource, err := h.Service.GetMemberByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Write response
	c.JSON(http.StatusOK, memberResource)
}

// CreateMember handles POST /members
func (h *MemberHandler) CreateMember(c *gin.Context) {
	var memberInput models.Member
	if err := c.ShouldBindJSON(&memberInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	memberResource, err := h.Service.CreateMember(memberInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusCreated, memberResource)
}

// UpdateMember handles PUT /members/:id
func (h *MemberHandler) UpdateMember(c *gin.Context) {
	id := c.Param("id")

	var memberInput models.Member
	if err := c.ShouldBindJSON(&memberInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure ID consistency
	if memberInput.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID in URL and body do not match"})
		return
	}

	memberResource, err := h.Service.UpdateMember(memberInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, memberResource)
}

// DeleteMember handles DELETE /members/:id
func (h *MemberHandler) DeleteMember(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteMember(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// RegisterMemberRoutes registers the member routes with the provided router
func RegisterMemberRoutes(router *gin.RouterGroup, db *gorm.DB) {
	memberService := service.NewMemberService(db)
	memberHandler := NewMemberHandler(memberService)

	memberGroup := router.Group("/members")
	{
		memberGroup.GET("", memberHandler.ListMembers)
		memberGroup.GET("/:id", memberHandler.GetMember)
		memberGroup.POST("", memberHandler.CreateMember)
		memberGroup.PUT("/:id", memberHandler.UpdateMember)
		memberGroup.DELETE("/:id", memberHandler.DeleteMember)
	}
}
