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

type CompanyHandler struct {
	Service *service.CompanyService
}

func NewCompanyHandler(service *service.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		Service: service,
	}
}

// ListCompanys handles GET /companys
func (h *CompanyHandler) ListCompanys(c *gin.Context) {
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
	listResponse, err := h.Service.ListCompanies(listRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, listResponse)
}

// GetCompany handles GET /companys/:id
func (h *CompanyHandler) GetCompany(c *gin.Context) {
	id := c.Param("id")

	companyResource, err := h.Service.GetCompanyByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Write response
	c.JSON(http.StatusOK, companyResource)
}

// CreateCompany handles POST /companys
func (h *CompanyHandler) CreateCompany(c *gin.Context) {
	var companyInput models.Company
	if err := c.ShouldBindJSON(&companyInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyResource, err := h.Service.CreateCompany(companyInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusCreated, companyResource)
}

// UpdateCompany handles PUT /companys/:id
func (h *CompanyHandler) UpdateCompany(c *gin.Context) {
	id := c.Param("id")

	var companyInput models.Company
	if err := c.ShouldBindJSON(&companyInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure ID consistency
	if companyInput.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID in URL and body do not match"})
		return
	}

	companyResource, err := h.Service.UpdateCompany(companyInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, companyResource)
}

// DeleteCompany handles DELETE /companys/:id
func (h *CompanyHandler) DeleteCompany(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteCompany(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// RegisterCompanyRoutes registers the company routes with the provided router
func RegisterCompanyRoutes(router *gin.RouterGroup, db *gorm.DB) {
	companyService := service.NewCompanyService(db)
	companyHandler := NewCompanyHandler(companyService)

	companyGroup := router.Group("/company")
	{
		companyGroup.GET("", companyHandler.ListCompanys)
		companyGroup.GET("/:id", companyHandler.GetCompany)
		companyGroup.POST("", companyHandler.CreateCompany)
		companyGroup.PUT("/:id", companyHandler.UpdateCompany)
		companyGroup.DELETE("/:id", companyHandler.DeleteCompany)
	}
}
