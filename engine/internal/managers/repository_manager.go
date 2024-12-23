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

func (r *Repository[T]) DeleteMany(ids []uint) error {
	if len(ids) == 0 {
		return fmt.Errorf("no IDs provided for deletion")
	}

	// Perform the deletion using GORM's Delete with a WHERE IN clause
	result := r.DB.Client.Delete(new(T), "id IN ?", ids)
	if result.Error != nil {
		return fmt.Errorf("failed to delete entities: %w", result.Error)
	}

	// Optionally, you can check how many records were deleted
	if result.RowsAffected != int64(len(ids)) {
		return fmt.Errorf("expected to delete %d entities, but deleted %d", len(ids), result.RowsAffected)
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

func (r *Repository[T]) prepareFilteredDB(db *gorm.DB, req string) (filter.PaginatedRequest, *gorm.DB, error) {
	var filterReq filter.PaginatedRequest

	if err := filter.DecodeBase64JSON(req, &filterReq); err != nil {
		return filterReq, nil, fmt.Errorf("failed to decode filter request: %w", err)
	}

	clientDB := r.DB.Client
	if db != nil {
		clientDB = db
	}
	filteredDB := filter.ApplyFilters(clientDB, filterReq)
	filteredDB = filteredDB.Session(&gorm.Session{
		Logger: clientDB.Logger.LogMode(logger.Info),
	})
	stmt := filteredDB.Statement
	if stmt != nil {
		sql := stmt.SQL.String()
		vars := stmt.Vars
		fmt.Printf("Executed SQL: %s\nWith Variables: %v\n", sql, vars)
	}
	return filterReq, filteredDB, nil
}

func (r *Repository[T]) GetFilteredResults(db *gorm.DB, req string) ([]*T, error) {
	var results []*T
	_, filteredDB, err := r.prepareFilteredDB(db, req)
	if err != nil {
		return nil, err
	}
	if err := filteredDB.Find(&results).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve filtered results: %w", err)
	}
	return results, nil
}

func (r *Repository[T]) GetPaginatedResult(db *gorm.DB, req string) (filter.FilterPages[T], error) {
	var results []*T
	var totalSize int64
	filterReq, filteredDB, err := r.prepareFilteredDB(db, req)
	if err != nil {
		return filter.FilterPages[T]{}, err
	}
	if err := filteredDB.Model(new(T)).Count(&totalSize).Error; err != nil {
		return filter.FilterPages[T]{}, err
	}
	totalPages := int(math.Ceil(float64(totalSize) / float64(filterReq.PageSize)))
	if totalPages == 0 {
		totalPages = 1
	}
	paginatedDB := filteredDB.Offset((filterReq.PageIndex - 1) * filterReq.PageSize).Limit(filterReq.PageSize)
	if err := paginatedDB.Find(&results).Error; err != nil {
		return filter.FilterPages[T]{}, fmt.Errorf("failed to retrieve paginated results: %w", err)
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
		PageIndex: filterReq.PageIndex,
		TotalPage: totalPages,
		PageSize:  filterReq.PageSize,
		TotalSize: int(totalSize),
		Pages:     pages,
	}, nil
}
