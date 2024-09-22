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

type MediaHandler struct {
	Service *service.MediaService
}

func NewMediaHandler(service *service.MediaService) *MediaHandler {
	return &MediaHandler{
		Service: service,
	}
}

// ListMedias handles GET /medias
func (h *MediaHandler) ListMedias(c *gin.Context) {
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
	listResponse, err := h.Service.ListMedia(listRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, listResponse)
}

// GetMedia handles GET /medias/:id
func (h *MediaHandler) GetMedia(c *gin.Context) {
	id := c.Param("id")

	mediaResource, err := h.Service.GetMediaByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Write response
	c.JSON(http.StatusOK, mediaResource)
}

// CreateMedia handles POST /medias
func (h *MediaHandler) CreateMedia(c *gin.Context) {
	var mediaInput models.Media
	if err := c.ShouldBindJSON(&mediaInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	mediaResource, err := h.Service.CreateMedia(mediaInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusCreated, mediaResource)
}

// UpdateMedia handles PUT /medias/:id
func (h *MediaHandler) UpdateMedia(c *gin.Context) {
	id := c.Param("id")

	var mediaInput models.Media
	if err := c.ShouldBindJSON(&mediaInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure ID consistency
	if mediaInput.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID in URL and body do not match"})
		return
	}

	mediaResource, err := h.Service.UpdateMedia(mediaInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, mediaResource)
}

// DeleteMedia handles DELETE /medias/:id
func (h *MediaHandler) DeleteMedia(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteMedia(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// RegisterMediaRoutes registers the media routes with the provided router
func RegisterMediaRoutes(router *gin.RouterGroup, db *gorm.DB) {
	mediaService := service.NewMediaService(db)
	mediaHandler := NewMediaHandler(mediaService)

	mediaGroup := router.Group("/medias")
	{
		mediaGroup.GET("", mediaHandler.ListMedias)
		mediaGroup.GET("/:id", mediaHandler.GetMedia)
		mediaGroup.POST("", mediaHandler.CreateMedia)
		mediaGroup.PUT("/:id", mediaHandler.UpdateMedia)
		mediaGroup.DELETE("/:id", mediaHandler.DeleteMedia)
	}
}
