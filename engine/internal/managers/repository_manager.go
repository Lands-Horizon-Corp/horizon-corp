// managers/managers.go
package managers

import (
	"errors"
	"fmt"
	"math"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers/filter"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers" // Adjust the import path as needed
	"gorm.io/gorm"
)

// Repository provides generic CRUD operations with filtering and pagination.
type Repository[T any] struct {
	DB *providers.DatabaseService
}

// NewRepository initializes a new Repository.
func NewRepository[T any](db *providers.DatabaseService) *Repository[T] {
	return &Repository[T]{DB: db}
}

// Create adds a new entity to the database.
func (r *Repository[T]) Create(entity *T) error {
	if err := r.DB.Client.Create(entity).Error; err != nil {
		return fmt.Errorf("failed to create entity: %w", err)
	}
	return nil
}

// FindByID retrieves an entity by its ID with optional preloads.
func (r *Repository[T]) FindByID(id uint, preloads ...string) (*T, error) {
	var entity T
	query := r.applyPreloads(r.DB.Client, preloads)

	result := query.First(&entity, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, fmt.Errorf("entity with ID %d not found: %w", id, result.Error)
	}
	return &entity, result.Error
}

// FindAll retrieves all entities with optional preloads.
func (r *Repository[T]) FindAll(preloads ...string) ([]*T, error) { // Changed to []*T
	var entities []*T // Changed from []T to []*T
	query := r.applyPreloads(r.DB.Client, preloads)

	result := query.Find(&entities)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve entities: %w", result.Error)
	}
	return entities, nil
}

// Update modifies an existing entity with optional preloads.
func (r *Repository[T]) Update(entity *T, preloads []string) error {
	tx := r.DB.Client

	// Apply preloads
	for _, preload := range preloads {
		tx = tx.Preload(preload)
	}

	if err := tx.Save(entity).Error; err != nil {
		return fmt.Errorf("failed to update entity: %w", err)
	}

	return nil
}

// Delete removes an entity by its ID.
func (r *Repository[T]) Delete(id uint) error {
	if err := r.DB.Client.Delete(new(T), id).Error; err != nil {
		return fmt.Errorf("failed to delete entity with ID %d: %w", id, err)
	}
	return nil
}

// applyPreloads applies preloads to the GORM query.
func (r *Repository[T]) applyPreloads(query *gorm.DB, preloads []string) *gorm.DB {
	for _, preload := range preloads {
		query = query.Preload(preload)
	}
	return query
}

// UpdateColumns performs a partial update on an entity's fields.
func (r *Repository[T]) UpdateColumns(id uint, updates T, preloads []string) (*T, error) {
	entity := new(T)
	db := r.DB.Client
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	if err := db.Model(entity).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}
	if err := db.First(entity, id).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *Repository[T]) GetPaginatedResult(db *gorm.DB, req string) (filter.FilterPages[T], error) {
	var results []*T
	var totalSize int64

	clientDB := r.DB.Client
	if db != nil {
		clientDB = db
	}
	var paginatedReq filter.PaginatedRequest
	if err := filter.DecodeBase64JSON(req, &paginatedReq); err != nil {
		return filter.FilterPages[T]{}, err
	}

	filteredDB := filter.ApplyFilters(clientDB, paginatedReq)

	err := filteredDB.Model(new(T)).Count(&totalSize).Error
	if err != nil {
		return filter.FilterPages[T]{}, err
	}

	totalPages := int(math.Ceil(float64(totalSize) / float64(paginatedReq.PageSize)))
	if totalPages == 0 {
		totalPages = 1
	}

	err = filteredDB.Offset((paginatedReq.PageIndex - 1) * paginatedReq.PageSize).Limit(paginatedReq.PageSize).Find(&results).Error
	if err != nil {
		return filter.FilterPages[T]{}, err
	}

	pages := make([]filter.Page, totalPages)
	for i := 0; i < totalPages; i++ {
		pages[i] = filter.Page{
			Page:      fmt.Sprintf("Page %d", i+1),
			PageIndex: i + 1,
		}
	}

	return filter.FilterPages[T]{
		Data:      results,
		PageIndex: paginatedReq.PageIndex,
		TotalPage: totalPages,
		PageSize:  paginatedReq.PageSize,
		TotalSize: int(totalSize),
		Pages:     pages,
	}, nil
}
