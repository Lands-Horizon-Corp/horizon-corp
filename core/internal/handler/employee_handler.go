// handler/employee_handler.go

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

type EmployeeHandler struct {
	Service *service.EmployeeService
}

func NewEmployeeHandler(service *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		Service: service,
	}
}

// ListEmployees handles GET /employees
func (h *EmployeeHandler) ListEmployees(c *gin.Context) {
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
	listResponse, err := h.Service.ListEmployees(listRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, listResponse)
}

// GetEmployee handles GET /employees/:id
func (h *EmployeeHandler) GetEmployee(c *gin.Context) {
	id := c.Param("id")

	employeeResource, err := h.Service.GetEmployeeByID(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Write response
	c.JSON(http.StatusOK, employeeResource)
}

// CreateEmployee handles POST /employees
func (h *EmployeeHandler) CreateEmployee(c *gin.Context) {
	var employeeInput models.Employee
	if err := c.ShouldBindJSON(&employeeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employeeResource, err := h.Service.CreateEmployee(employeeInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusCreated, employeeResource)
}

// UpdateEmployee handles PUT /employees/:id
func (h *EmployeeHandler) UpdateEmployee(c *gin.Context) {
	id := c.Param("id")

	var employeeInput models.Employee
	if err := c.ShouldBindJSON(&employeeInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ensure ID consistency
	if employeeInput.ID != id {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID in URL and body do not match"})
		return
	}

	employeeResource, err := h.Service.UpdateEmployee(employeeInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Write response
	c.JSON(http.StatusOK, employeeResource)
}

// DeleteEmployee handles DELETE /employees/:id
func (h *EmployeeHandler) DeleteEmployee(c *gin.Context) {
	id := c.Param("id")

	err := h.Service.DeleteEmployee(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

// RegisterEmployeeRoutes registers the employee routes with the provided router
func RegisterEmployeeRoutes(router *gin.RouterGroup, db *gorm.DB) {
	employeeService := service.NewEmployeeService(db)
	employeeHandler := NewEmployeeHandler(employeeService)

	employeeGroup := router.Group("/employees")
	{
		employeeGroup.GET("", employeeHandler.ListEmployees)
		employeeGroup.GET("/:id", employeeHandler.GetEmployee)
		employeeGroup.POST("", employeeHandler.CreateEmployee)
		employeeGroup.PUT("/:id", employeeHandler.UpdateEmployee)
		employeeGroup.DELETE("/:id", employeeHandler.DeleteEmployee)
	}
}
