// managers/managers.go
package managers

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers" // Adjust the import path as needed
	"gorm.io/gorm"
)

// --------------------
// Filter Definitions
// --------------------

// FilterMode represents the mode of filtering.
type FilterMode string

// Filter modes
const (
	ModeEqual        FilterMode = "equal"
	ModeNotEqual     FilterMode = "nequal"
	ModeContains     FilterMode = "contains"
	ModeNotContains  FilterMode = "ncontains"
	ModeStartsWith   FilterMode = "startswith"
	ModeEndsWith     FilterMode = "endswith"
	ModeIsEmpty      FilterMode = "isempty"
	ModeIsNotEmpty   FilterMode = "isnotempty"
	ModeGreaterThan  FilterMode = "gt"
	ModeGreaterEqual FilterMode = "gte"
	ModeLessThan     FilterMode = "lt"
	ModeLessEqual    FilterMode = "lte"
	ModeRange        FilterMode = "range"
	ModeBetween      FilterMode = "between"
	ModeBefore       FilterMode = "before"
	ModeAfter        FilterMode = "after"
	ModeBooleanTrue  FilterMode = "true"
	ModeBooleanFalse FilterMode = "false"
)

// FilterDataType represents the data type of the filter value.
type FilterDataType string

// Data types
const (
	DataTypeText     FilterDataType = "text"
	DataTypeNumber   FilterDataType = "number"
	DataTypeFloat    FilterDataType = "float"
	DataTypeBoolean  FilterDataType = "boolean"
	DataTypeDate     FilterDataType = "date"
	DataTypeDatetime FilterDataType = "datetime"
	DataTypeTime     FilterDataType = "time"
)

// FilterValue is a generic type for filter values.
type FilterValue = any

// RangeValue represents a from-to structure for range queries.
type RangeValue struct {
	From FilterValue `json:"from"`
	To   FilterValue `json:"to"`
}

// Filter interface for different filter implementations
type Filter interface {
	GetField() string
	GetMode() FilterMode
	GetDataType() string
	GetValue() FilterValue
	IsMultiple() bool
}

// FilterStruct is a unified filter structure that can represent both single and range values.
type FilterStruct struct {
	Field    string         `json:"field"`
	Mode     FilterMode     `json:"mode"`
	DataType FilterDataType `json:"dataType"`
	Value    interface{}    `json:"value"`
}

// GetField returns the field name of the filter.
func (f FilterStruct) GetField() string { return f.Field }

// GetMode returns the filter mode.
func (f FilterStruct) GetMode() FilterMode { return f.Mode }

// GetDataType returns the data type as a string.
func (f FilterStruct) GetDataType() string { return string(f.DataType) }

// GetValue returns the filter value.
func (f FilterStruct) GetValue() FilterValue { return f.Value }

// IsMultiple determines if the filter has multiple values.
func (f FilterStruct) IsMultiple() bool {
	switch f.Mode {
	case ModeContains, ModeNotContains, ModeBetween, ModeRange, ModeStartsWith, ModeEndsWith:
		switch f.Value.(type) {
		case []interface{}, []string, []float64, []int, []bool, []RangeValue:
			return true
		}
	}
	return false
}

// --------------------
// Pagination and Request Structures
// --------------------

// Page represents pagination information.
type Page struct {
	Page      string `json:"page"`
	PageIndex int    `json:"pageIndex"`
}

// FilterPages encapsulates paginated data and metadata.
type FilterPages[T any] struct {
	Data      []*T   `json:"data"` // Changed from []T to []*T
	PageIndex int    `json:"pageIndex"`
	TotalPage int    `json:"totalPage"`
	PageSize  int    `json:"pageSize"`
	TotalSize int    `json:"totalSize"`
	Pages     []Page `json:"pages"`
}

// PaginatedRequest is a generic request structure for pagination and filtering.
type PaginatedRequest struct {
	Filters   []Filter `json:"filters"`
	Preloads  []string `json:"preloads"`
	PageIndex int      `json:"pageIndex"`
	PageSize  int      `json:"pageSize"`
	Logic     string   `json:"logic"` // "and" or "or"
}

// --------------------
// Repository Definition
// --------------------

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

// --------------------
// Filtering and Pagination
// --------------------

// ApplyFilters applies the filters from PaginatedRequest to the GORM DB instance.
func (r *Repository[T]) ApplyFilters(db *gorm.DB, request PaginatedRequest) *gorm.DB {
	if strings.ToLower(request.Logic) == "or" {
		db = db.Where(func(tx *gorm.DB) *gorm.DB {
			for i, filter := range request.Filters {
				if i == 0 {
					tx = r.applySingleFilter(tx, filter)
				} else {
					tx = tx.Or(r.applySingleFilter(&gorm.DB{}, filter))
				}
			}
			return tx
		})
	} else {
		for _, filter := range request.Filters {
			db = r.applySingleFilter(db, filter)
		}
	}
	for _, preload := range request.Preloads {
		db = db.Preload(preload)
	}
	return db
}

// applySingleFilter applies a single filter to the GORM DB instance.
func (r *Repository[T]) applySingleFilter(db *gorm.DB, filter Filter) *gorm.DB {
	field := filter.GetField()
	mode := filter.GetMode()
	value := filter.GetValue()

	// Check if the filter has multiple values
	if filter.IsMultiple() {
		// Handle multiple values based on the mode
		switch mode {
		case ModeContains, ModeNotContains, ModeStartsWith, ModeEndsWith:
			// For string operations with multiple values, use OR conditions
			switch v := value.(type) {
			case []string:
				conditions := make([]string, 0, len(v))
				args := make([]interface{}, 0, len(v))
				for _, val := range v {
					var condition string
					switch mode {
					case ModeContains:
						condition = fmt.Sprintf("%s LIKE ?", field)
						args = append(args, fmt.Sprintf("%%%s%%", val))
					case ModeNotContains:
						condition = fmt.Sprintf("%s NOT LIKE ?", field)
						args = append(args, fmt.Sprintf("%%%s%%", val))
					case ModeStartsWith:
						condition = fmt.Sprintf("%s LIKE ?", field)
						args = append(args, fmt.Sprintf("%s%%", val))
					case ModeEndsWith:
						condition = fmt.Sprintf("%s LIKE ?", field)
						args = append(args, fmt.Sprintf("%%%s", val))
					}
					conditions = append(conditions, condition)
				}
				db = db.Where(strings.Join(conditions, " OR "), args...)
			default:
				// Handle other slice types if necessary
			}
		case ModeBetween, ModeRange:
			// For range operations with multiple RangeValues, use OR conditions
			switch v := value.(type) {
			case []RangeValue:
				for _, rv := range v {
					db = db.Or(fmt.Sprintf("%s BETWEEN ? AND ?", field), rv.From, rv.To)
				}
			}
		default:
		}
	} else {
		switch mode {
		case ModeEqual:
			return db.Where(fmt.Sprintf("%s = ?", field), value)
		case ModeNotEqual:
			return db.Where(fmt.Sprintf("%s != ?", field), value)
		case ModeContains:
			return db.Where(fmt.Sprintf("%s LIKE ?", field), fmt.Sprintf("%%%v%%", value))
		case ModeNotContains:
			return db.Where(fmt.Sprintf("%s NOT LIKE ?", field), fmt.Sprintf("%%%v%%", value))
		case ModeStartsWith:
			return db.Where(fmt.Sprintf("%s LIKE ?", field), fmt.Sprintf("%v%%", value))
		case ModeEndsWith:
			return db.Where(fmt.Sprintf("%s LIKE ?", field), fmt.Sprintf("%%%v", value))
		case ModeGreaterThan:
			return db.Where(fmt.Sprintf("%s > ?", field), value)
		case ModeGreaterEqual:
			return db.Where(fmt.Sprintf("%s >= ?", field), value)
		case ModeLessThan:
			return db.Where(fmt.Sprintf("%s < ?", field), value)
		case ModeLessEqual:
			return db.Where(fmt.Sprintf("%s <= ?", field), value)
		case ModeRange, ModeBetween:
			// Assuming value is RangeValue
			switch v := value.(type) {
			case RangeValue:
				return db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", field), v.From, v.To)
			}
		case ModeBefore:
			return db.Where(fmt.Sprintf("%s < ?", field), value)
		case ModeAfter:
			return db.Where(fmt.Sprintf("%s > ?", field), value)
		case ModeBooleanTrue:
			return db.Where(fmt.Sprintf("%s = ?", field), true)
		case ModeBooleanFalse:
			return db.Where(fmt.Sprintf("%s = ?", field), false)
		default:
			return db
		}
	}

	return db
}

func (r *Repository[T]) GetPaginatedResult(db *gorm.DB, request PaginatedRequest) (FilterPages[T], error) {
	var results []*T // []*T where T is models.Company
	var totalSize int64

	clientDB := r.DB.Client
	if db != nil {
		clientDB = db
	}

	// Validate the request before applying filters
	if err := r.ValidatePaginatedRequest(request); err != nil {
		return FilterPages[T]{}, err
	}

	// Apply filters and preloads
	filteredDB := r.ApplyFilters(clientDB, request)

	// Count total records
	err := filteredDB.Model(new(T)).Count(&totalSize).Error
	if err != nil {
		return FilterPages[T]{}, err
	}

	// Calculate total pages
	totalPages := int(math.Ceil(float64(totalSize) / float64(request.PageSize)))
	if totalPages == 0 {
		totalPages = 1 // At least one page
	}

	// Fetch paginated data
	err = filteredDB.Offset((request.PageIndex - 1) * request.PageSize).
		Limit(request.PageSize).
		Find(&results).Error
	if err != nil {
		return FilterPages[T]{}, err
	}

	// Construct pages metadata
	pages := make([]Page, totalPages)
	for i := 0; i < totalPages; i++ {
		pages[i] = Page{
			Page:      fmt.Sprintf("Page %d", i+1),
			PageIndex: i + 1,
		}
	}

	return FilterPages[T]{
		Data:      results,
		PageIndex: request.PageIndex,
		TotalPage: totalPages,
		PageSize:  request.PageSize,
		TotalSize: int(totalSize),
		Pages:     pages,
	}, nil
}

// --------------------
// Validation
// --------------------

// ValidatePaginatedRequest validates the PaginatedRequest for consistency and correctness.
func (r *Repository[T]) ValidatePaginatedRequest(request PaginatedRequest) error {
	for _, filter := range request.Filters {
		if filter.GetField() == "" {
			return errors.New("filter field cannot be empty")
		}
		if filter.GetMode() == "" {
			return errors.New("filter mode cannot be empty")
		}
		if filter.GetDataType() == "" {
			return errors.New("filter dataType cannot be empty")
		}
		// Additional validations based on mode and dataType can be added here
	}
	return nil
}

// --------------------
// Sample Data Construction
// --------------------

// Sample constructs a sample PaginatedRequest with various filters.
func Sample() PaginatedRequest {
	// Single-value filters
	textFilter := FilterStruct{
		Field:    "username",
		Mode:     ModeContains,
		DataType: DataTypeText,
		Value:    "John",
	}

	numberFilter := FilterStruct{
		Field:    "age",
		Mode:     ModeGreaterThan,
		DataType: DataTypeNumber,
		Value:    30,
	}

	floatFilter := FilterStruct{
		Field:    "salary",
		Mode:     ModeGreaterEqual,
		DataType: DataTypeFloat,
		Value:    50000.5,
	}

	booleanFilter := FilterStruct{
		Field:    "isActive",
		Mode:     ModeBooleanTrue,
		DataType: DataTypeBoolean,
		Value:    true,
	}

	dateFilter := FilterStruct{
		Field:    "creationDate",
		Mode:     ModeAfter,
		DataType: DataTypeDate,
		Value:    "2024-01-01",
	}

	dateTimeFilter := FilterStruct{
		Field:    "lastLogin",
		Mode:     ModeBefore,
		DataType: DataTypeDatetime,
		Value:    "2024-12-31T23:59:59Z",
	}

	singleTimeFilter := FilterStruct{
		Field:    "scheduledTime",
		Mode:     ModeEqual,
		DataType: DataTypeTime,
		Value:    "15:30:00",
	}

	// Array-value filters
	textArrayFilter := FilterStruct{
		Field:    "roles",
		Mode:     ModeContains,
		DataType: DataTypeText,
		Value:    []string{"admin", "editor", "viewer"},
	}

	numberArrayFilter := FilterStruct{
		Field:    "scores",
		Mode:     ModeBetween,
		DataType: DataTypeNumber,
		Value:    []float64{50, 75, 100},
	}

	booleanArrayFilter := FilterStruct{
		Field:    "flags",
		Mode:     ModeNotContains,
		DataType: DataTypeBoolean,
		Value:    []bool{true, false},
	}

	// Range filters
	dateRangeFilter := FilterStruct{
		Field:    "dateRange",
		Mode:     ModeRange,
		DataType: DataTypeDate,
		Value: RangeValue{
			From: "2024-01-01",
			To:   "2024-12-31",
		},
	}

	numberRangeFilter := FilterStruct{
		Field:    "betweenAges",
		Mode:     ModeBetween,
		DataType: DataTypeNumber,
		Value: RangeValue{
			From: 18,
			To:   30,
		},
	}

	timeRangeFilter := FilterStruct{
		Field:    "workingHours",
		Mode:     ModeRange,
		DataType: DataTypeTime,
		Value: RangeValue{
			From: "09:00:00",
			To:   "12:00:00",
		},
	}

	// Multiple ranges in one filter (slice of RangeValue)
	availableSlotsFilter := FilterStruct{
		Field:    "availableSlots",
		Mode:     ModeRange,
		DataType: DataTypeTime,
		Value: []RangeValue{
			{From: "09:00:00", To: "12:00:00"},
			{From: "13:00:00", To: "17:00:00"},
		},
	}

	// Nested filters
	ordersStatusFilter := FilterStruct{
		Field:    "Orders.status",
		Mode:     ModeEqual,
		DataType: DataTypeText,
		Value:    "completed",
	}

	ordersAmountFilter := FilterStruct{
		Field:    "Orders.totalAmount",
		Mode:     ModeGreaterThan,
		DataType: DataTypeFloat,
		Value:    100.0,
	}

	profileCountryFilter := FilterStruct{
		Field:    "Profile.country",
		Mode:     ModeEqual,
		DataType: DataTypeText,
		Value:    "USA",
	}

	profileAgeFilter := FilterStruct{
		Field:    "Profile.age",
		Mode:     ModeBetween,
		DataType: DataTypeNumber,
		Value: []RangeValue{
			{From: 10, To: 20},
			{From: 30, To: 40},
		},
	}

	profileNameFilter := FilterStruct{
		Field:    "Profile.name",
		Mode:     ModeStartsWith,
		DataType: DataTypeText,
		Value:    []string{"sample", "sample1", "sample2"},
	}

	filters := []Filter{
		textFilter,
		numberFilter,
		floatFilter,
		booleanFilter,
		dateFilter,
		dateTimeFilter,
		singleTimeFilter,
		textArrayFilter,
		numberArrayFilter,
		booleanArrayFilter,
		dateRangeFilter,
		numberRangeFilter,
		timeRangeFilter,
		availableSlotsFilter,
		ordersStatusFilter,
		ordersAmountFilter,
		profileCountryFilter,
		profileAgeFilter,
		profileNameFilter,
	}

	request := PaginatedRequest{
		Filters:   filters,
		Preloads:  []string{"Orders", "Profile"},
		PageIndex: 1,
		PageSize:  20,
		Logic:     "and",
	}

	return request
}
