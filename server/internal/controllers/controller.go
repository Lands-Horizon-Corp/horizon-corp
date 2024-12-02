package controllers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// BaseController is a generic controller for CRUD operations with preloads
type BaseController[Model any, Request requests.Validatable, Resource any] struct {
	Repo           *repositories.Repository[Model]
	ToModel        func(req Request) *Model
	ToResource     func(model *Model) *Resource
	ToResourceList func(models []*Model) []*Resource
	UpdateModel    func(model *Model, req Request)
}

// NewBaseController creates a new generic BaseController
func NewBaseController[Model any, Request requests.Validatable, Resource any](
	db *gorm.DB,
	toModel func(req Request) *Model,
	toResource func(model *Model) *Resource,
	toResourceList func(models []*Model) []*Resource,
	updateModel func(model *Model, req Request),
) *BaseController[Model, Request, Resource] {
	return &BaseController[Model, Request, Resource]{
		Repo:           repositories.NewRepository[Model](db),
		ToModel:        toModel,
		ToResource:     toResource,
		ToResourceList: toResourceList,
		UpdateModel:    updateModel,
	}
}

// Create handles the creation of a new entity
func (c *BaseController[Model, Request, Resource]) Create(ctx *gin.Context) {
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Call req.Validate() directly
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	model := c.ToModel(req)
	if err := c.Repo.Create(model); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resource := c.ToResource(model)
	ctx.JSON(http.StatusCreated, resource)
}

// GetByID retrieves an entity by ID with optional preloads
func (c *BaseController[Model, Request, Resource]) GetByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	preloads := getPreloadsFromQuery(ctx)
	model, err := c.Repo.GetByID(uint(id), preloads)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	resource := c.ToResource(model)
	ctx.JSON(http.StatusOK, resource)
}

// GetAll retrieves all entities with optional preloads
func (c *BaseController[Model, Request, Resource]) GetAll(ctx *gin.Context) {
	preloads := getPreloadsFromQuery(ctx)
	modelPointers, err := c.Repo.GetAll(preloads)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resources := c.ToResourceList(modelPointers)
	ctx.JSON(http.StatusOK, resources)
}

func (c *BaseController[Model, Request, Resource]) Filter(ctx *gin.Context) {
	filterParam := ctx.Query("filter")
	decodedFilter, err := base64.StdEncoding.DecodeString(filterParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid base64 encoding"})
		return
	}
	fmt.Println(decodedFilter)

	ctx.JSON(http.StatusOK, gin.H{
		"data":      []string{},
		"pageIndex": 1,
		"totalPage": 1,
		"pageSize":  1,
		"totalSize": 1,
		"pages": []string{
			"/api/filter-something-1",
			"/api/filter-something-1",
			"/api/filter-something-1",
			"/api/filter-something-1",
		},
	})
}

// func (c *BaseController[Model, Request, Resource]) Filter(ctx *gin.Context) {
// 	preloads := getPreloadsFromQuery(ctx)
// 	pageIndex, pageSize := getPageParamsFromQuery(ctx)

// 	var filterRequest struct {
// 		Filters []repositories.Filter `json:"filters"`
// 	}
// 	if err := ctx.ShouldBindJSON(&filterRequest); err != nil && err != io.EOF {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	filters := filterRequest.Filters

// 	result, err := c.Repo.ApplyFilters(filters, preloads, pageIndex, pageSize)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	resources := c.ToResourceList(result.Data)
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"data":      resources,
// 		"pageIndex": result.PageIndex,
// 		"totalPage": result.TotalPage,
// 		"pageSize":  result.PageSize,
// 		"totalSize": result.TotalSize,
// 		"pages":     result.Pages,
// 	})
// }

// Update updates an existing entity by ID
func (c *BaseController[Model, Request, Resource]) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	preloads := getPreloadsFromQuery(ctx)
	model, err := c.Repo.GetByID(uint(id), preloads)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Optional: Ensure the ID in the request matches the URL parameter
	if reqWithID, ok := any(req).(interface{ GetID() uint }); ok {
		if reqWithID.GetID() != uint(id) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID in the request body does not match the URL parameter"})
			return
		}
	}

	c.UpdateModel(model, req)
	if err := c.Repo.Update(model, preloads); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resource := c.ToResource(model)
	ctx.JSON(http.StatusOK, resource)
}

// Delete removes an entity by ID
func (c *BaseController[Model, Request, Resource]) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.Repo.Delete(uint(id)); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Resource not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	ctx.JSON(http.StatusNoContent, nil)
}

// Helper function to extract preloads from query parameters
func getPreloadsFromQuery(ctx *gin.Context) []string {
	preloadsParam := ctx.Query("preloads")
	if preloadsParam == "" {
		return []string{}
	}
	return strings.Split(preloadsParam, ",")
}

// Helper function to extract pagination parameters
func getPageParamsFromQuery(ctx *gin.Context) (pageIndex int, pageSize int) {
	pageIndexStr := ctx.DefaultQuery("pageIndex", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")

	pageIndex, err := strconv.Atoi(pageIndexStr)
	if err != nil || pageIndex < 1 {
		pageIndex = 1
	}

	pageSize, err = strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	return pageIndex, pageSize
}

// {
// 	"filter": ["Owishi minecraft"]
// 	"preload": ["Role", "Role.Admin"]
// }
