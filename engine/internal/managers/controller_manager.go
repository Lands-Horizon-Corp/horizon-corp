package managers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// T = Model
// V = Validator
// R = Resource
type Controller[T any, V any, R any] struct {
	Repo         *Repository[T]
	Validate     func(request *V) error
	Resource     func(resource *T) *R
	ResourceList func(resources []*T) []*R
}

func NewController[T any, V any, R any](
	repo *Repository[T],
	Validate func(request *V) error,
	Resource func(resource *T) *R,
	ResourceList func(resources []*T) []*R,
) *Controller[T, V, R] {
	return &Controller[T, V, R]{
		Repo:         repo,
		Validate:     Validate,
		Resource:     Resource,
		ResourceList: ResourceList,
	}
}

// Helper function to extract preload parameters from the request
func getPreloads(c *gin.Context) []string {
	preloads := c.QueryArray("preloads")
	return preloads
}

func (h *Controller[T, V, R]) Create(c *gin.Context) {
	var req V
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Validate(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}
	var entity T
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Repo.Create(&entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, h.Resource(&entity))
}

func (h *Controller[T, V, R]) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	preloads := getPreloads(c)
	entity, err := h.Repo.FindByID(uint(id), preloads...)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, h.Resource(entity))
}

func (h *Controller[T, V, R]) GetAll(c *gin.Context) {
	preloads := getPreloads(c)
	entities, err := h.Repo.FindAll(preloads...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	entityPtrs := make([]*T, len(entities))
	for i := range entities {
		entityPtrs[i] = &entities[i]
	}
	c.JSON(http.StatusOK, h.ResourceList(entityPtrs))
}

func (h *Controller[T, V, R]) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Get preloads from the query
	preloads := getPreloads(c)

	// Fetch the entity with preloads
	entity, err := h.Repo.FindByID(uint(id), preloads...)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// Bind the request JSON to the entity
	if err := c.ShouldBindJSON(entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the request
	var req V
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Validate(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	// Perform the update
	if err := h.Repo.Update(entity, preloads); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, h.Resource(entity))
}

func (h *Controller[T, V, R]) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := h.Repo.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Entity deleted successfully"})
}
