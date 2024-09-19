package helpers

import (
	"fmt"
	"math"

	"github.com/jinzhu/gorm"
)

// querybuilder/filter.go
type Filter struct {
	Field    string      `json:"field"`
	Operator string      `json:"operator"`
	Value    interface{} `json:"value"`
}

// querybuilder/pagination.go
type Pagination struct {
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	SortBy     string `json:"sort_by"`
	SortOrder  string `json:"sort_order"` // "asc" or "desc"
	Total      int64  `json:"total"`
	TotalPages int    `json:"total_pages"`
	PrevPage   int    `json:"prev_page"`
	NextPage   int    `json:"next_page"`
}

type ListRequest struct {
	Filters    []Filter    `json:"filters"`
	Pagination *Pagination `json:"pagination"`
}

// ApplyFilters applies the given filters to the GORM DB query
func ApplyFilters(db *gorm.DB, filters []Filter) *gorm.DB {
	for _, filter := range filters {
		switch filter.Operator {
		case "eq":
			db = db.Where(fmt.Sprintf("%s = ?", filter.Field), filter.Value)
		case "gt":
			db = db.Where(fmt.Sprintf("%s > ?", filter.Field), filter.Value)
		case "gte":
			db = db.Where(fmt.Sprintf("%s >= ?", filter.Field), filter.Value)
		case "lt":
			db = db.Where(fmt.Sprintf("%s < ?", filter.Field), filter.Value)
		case "lte":
			db = db.Where(fmt.Sprintf("%s <= ?", filter.Field), filter.Value)
		case "contains":
			db = db.Where(fmt.Sprintf("%s LIKE ?", filter.Field), "%"+filter.Value.(string)+"%")
		case "starts_with":
			db = db.Where(fmt.Sprintf("%s LIKE ?", filter.Field), filter.Value.(string)+"%")
		case "ends_with":
			db = db.Where(fmt.Sprintf("%s LIKE ?", filter.Field), "%"+filter.Value.(string))
		case "range":
			values := filter.Value.([]interface{})
			db = db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", filter.Field), values[0], values[1])
		case "in":
			db = db.Where(fmt.Sprintf("%s IN (?)", filter.Field), filter.Value)
		default:
			// Handle unsupported operator if necessary
		}
	}
	return db
}

// ApplyPagination applies pagination and sorting to the GORM DB query
func ApplyPagination(db *gorm.DB, pagination *Pagination) *gorm.DB {
	if pagination.Limit <= 0 {
		pagination.Limit = 10 // Default limit
	}
	if pagination.Page <= 0 {
		pagination.Page = 1 // Default page
	}
	offset := (pagination.Page - 1) * pagination.Limit
	db = db.Offset(offset).Limit(pagination.Limit)

	if pagination.SortBy != "" {
		order := pagination.SortBy
		if pagination.SortOrder != "" {
			order += " " + pagination.SortOrder
		}
		db = db.Order(order)
	}
	return db
}

// Paginate executes the query with filters and pagination, and populates the result
func Paginate(db *gorm.DB, model interface{}, filters []Filter, pagination *Pagination) error {
	// Apply filters
	db = ApplyFilters(db, filters)

	// Count total records
	var total int64
	if err := db.Model(model).Count(&total).Error; err != nil {
		return err
	}
	pagination.Total = total

	// Calculate total pages
	pagination.TotalPages = int(math.Ceil(float64(total) / float64(pagination.Limit)))

	// Calculate previous and next pages
	if pagination.Page > 1 {
		pagination.PrevPage = pagination.Page - 1
	}
	if pagination.Page < pagination.TotalPages {
		pagination.NextPage = pagination.Page + 1
	}

	// Apply pagination and sorting
	db = ApplyPagination(db, pagination)

	// Execute query and populate result
	if err := db.Find(model).Error; err != nil {
		return err
	}

	return nil
}
