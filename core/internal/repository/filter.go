package repository

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (r *ModelRepository[T]) List(req ListRequest, eagerLoads []string) (ListResponse[T], error) {
	var entities []T
	var total int64

	// Start building the query with preloads
	query := r.preload(r.db, eagerLoads)

	// Apply filters
	for _, filter := range req.Filters {
		var err error
		query, err = applyFilter(query, filter)
		if err != nil {
			return ListResponse[T]{}, err
		}
	}

	// Count total records before pagination
	if err := query.Count(&total).Error; err != nil {
		return ListResponse[T]{}, err
	}

	// Apply pagination
	limit := req.Pagination.Limit
	if limit <= 0 {
		limit = 10 // default limit
	}
	page := req.Pagination.Page
	if page <= 0 {
		page = 1 // default page
	}
	offset := (page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	// Apply sorting with validation
	if req.Pagination.SortBy != "" {
		if !isValidField(req.Pagination.SortBy) {
			return ListResponse[T]{}, fmt.Errorf("invalid sort field: %s", req.Pagination.SortBy)
		}
		sortOrder := strings.ToUpper(req.Pagination.SortOrder)
		if sortOrder != "ASC" && sortOrder != "DESC" {
			sortOrder = "ASC" // default sort order
		}
		orderClause := fmt.Sprintf("%s %s", req.Pagination.SortBy, sortOrder)
		query = query.Order(orderClause)
	}

	// Execute the query
	if err := query.Find(&entities).Error; err != nil {
		return ListResponse[T]{}, err
	}

	// Calculate pagination details
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	prevPage := page - 1
	if prevPage < 1 {
		prevPage = 1
	}
	nextPage := page + 1
	if nextPage > totalPages {
		nextPage = totalPages
	}

	pagination := req.Pagination
	pagination.Total = total
	pagination.TotalPages = totalPages
	pagination.PrevPage = prevPage
	pagination.NextPage = nextPage

	return ListResponse[T]{
		Data:       entities,
		Pagination: pagination,
	}, nil
}

// applyFilter applies the given filter to the Gorm DB query
func applyFilter(query *gorm.DB, filter Filter) (*gorm.DB, error) {
	// Validate the operator
	if !filter.Operator.IsValid() {
		return nil, fmt.Errorf("unsupported operator: %s", filter.Operator)
	}

	// Retrieve the handler based on the operator
	handler, exists := OperatorHandlers[filter.Operator]
	if !exists {
		return nil, fmt.Errorf("operator handler not found for operator: %s", filter.Operator)
	}

	// Validate the field name to prevent SQL injection
	if !isValidField(filter.Field) {
		return nil, fmt.Errorf("invalid field name: %s", filter.Field)
	}

	// Execute the handler
	newQuery, err := handler(query, filter)
	if err != nil {
		return nil, err
	}

	return newQuery, nil
}
