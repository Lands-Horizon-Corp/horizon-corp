package managers

import (
	"errors"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DeleteManyRequest struct {
	IDs []uint `json:"ids" binding:"required,min=1"`
}

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

	entity, err := h.mapToEntity(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to map create data"})
		return
	}

	preloads := getPreloads(c)

	if err := h.Repo.Create(entity, preloads...); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, h.Resource(entity))
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
	c.JSON(http.StatusOK, h.ResourceList(entityPtrs))
}

func (h *Controller[T, V, R]) Update(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var req V
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Validate(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}
	updates, err := h.mapToEntity(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to map update data"})
		return
	}

	preloads := getPreloads(c)

	updatedEntity, err := h.Repo.UpdateByID(uint(id), updates, preloads...)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, h.Resource(updatedEntity))
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

func (h *Controller[T, V, R]) mapToEntity(validator *V) (*T, error) {

	var entity T
	vValue := reflect.ValueOf(validator).Elem()
	tValue := reflect.ValueOf(&entity).Elem()

	for i := 0; i < vValue.NumField(); i++ {
		tField := tValue.FieldByName(vValue.Type().Field(i).Name)
		if tField.IsValid() && tField.CanSet() {
			tField.Set(vValue.Field(i))
		}
	}

	return &entity, nil
}

func (h *Controller[T, V, R]) DeleteMany(c *gin.Context) {
	var req DeleteManyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Repo.DeleteMany(req.IDs); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "One or more entities not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Entities deleted successfully"})
}
