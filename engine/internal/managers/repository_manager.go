// managers/managers.go
package managers

import (
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/managers/filter"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers" // Adjust the import path as needed
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Repository provides generic CRUD operations with filtering and pagination.
type Repository[T any] struct {
	DB *providers.DatabaseService
}

// NewRepository initializes a new Repository.
func NewRepository[T any](db *providers.DatabaseService) *Repository[T] {
	return &Repository[T]{DB: db}
}

func (r *Repository[T]) Create(entity *T, preloads ...string) error {
	// Create the entity in the database
	if err := r.DB.Client.Create(entity).Error; err != nil {
		return fmt.Errorf("failed to create entity: %w", err)
	}

	// If no preloads are specified, return early
	if len(preloads) == 0 {
		return nil
	}

	// Use reflection to retrieve the primary key (assumed to be "ID" of type uint)
	v := reflect.ValueOf(entity).Elem()
	idField := v.FieldByName("ID")
	if !idField.IsValid() {
		return fmt.Errorf("entity does not have an 'ID' field for preloading")
	}

	if idField.Kind() != reflect.Uint && idField.Kind() != reflect.Uint32 && idField.Kind() != reflect.Uint64 {
		return fmt.Errorf("entity 'ID' field is not of type uint")
	}

	id := idField.Uint()

	// Apply preloads using the existing applyPreloads method
	query := r.applyPreloads(r.DB.Client, preloads)

	// Retrieve the created entity with preloads and update the original entity in place
	if err := query.First(entity, id).Error; err != nil {
		return fmt.Errorf("failed to preload associations: %w", err)
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

func (r *Repository[T]) UpdateByID(id uint, updates *T, preloads ...string) (*T, error) {
	var entity T
	db := r.DB.Client

	// Apply preloads
	for _, preload := range preloads {
		db = db.Preload(preload)
	}

	// Find the existing entity by ID
	if err := db.First(&entity, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("entity with ID %d not found: %w", id, err)
		}
		return nil, fmt.Errorf("failed to retrieve entity: %w", err)
	}

	// Update the entity with the provided updates
	if err := db.Model(&entity).Updates(updates).Error; err != nil {
		return nil, fmt.Errorf("failed to update entity: %w", err)
	}

	// Optionally, reload the entity with preloads
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	if err := db.First(&entity, id).Error; err != nil {
		return nil, fmt.Errorf("failed to reload updated entity: %w", err)
	}

	return &entity, nil
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

	// Apply filters and enable SQL query logging with Debug()
	filteredDB := filter.ApplyFilters(clientDB, paginatedReq).Session(&gorm.Session{Logger: clientDB.Logger.LogMode(logger.Info)})

	// Count total records
	err := filteredDB.Model(new(T)).Count(&totalSize).Error
	if err != nil {
		return filter.FilterPages[T]{}, err
	}

	// Log the SQL query manually
	stmt := filteredDB.Statement
	if stmt != nil {
		sql := stmt.SQL.String()
		vars := stmt.Vars
		fmt.Printf("SQL: %s\nVars: %v\n", sql, vars)
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalSize) / float64(paginatedReq.PageSize)))
	if totalPages == 0 {
		totalPages = 1
	}

	// Fetch paginated results
	err = filteredDB.Offset((paginatedReq.PageIndex - 1) * paginatedReq.PageSize).Limit(paginatedReq.PageSize).Find(&results).Error
	if err != nil {
		return filter.FilterPages[T]{}, err
	}

	// Create pagination metadata
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
