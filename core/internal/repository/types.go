package repository

import (
	"gorm.io/gorm"
)

type Operator string

type OperatorHandler func(query *gorm.DB, filter Filter) (*gorm.DB, error)

const (
	// Existing operators
	OpEquals             Operator = "equals"
	OpNotEquals          Operator = "not_equals"
	OpGreaterThan        Operator = "greater_than"
	OpGreaterThanOrEqual Operator = "greater_than_or_equal"
	OpLessThan           Operator = "less_than"
	OpLessThanOrEqual    Operator = "less_than_or_equal"
	OpRange              Operator = "range"
	OpStartsWith         Operator = "starts_with"
	OpEndsWith           Operator = "ends_with"
	OpIncludes           Operator = "includes"
	OpBooleanEquals      Operator = "boolean_equals"

	// Newly added operators
	OpIn        Operator = "in"
	OpNotIn     Operator = "not_in"
	OpIsNull    Operator = "is_null"
	OpIsNotNull Operator = "is_not_null"
)

type Filter struct {
	Field    string      `json:"field"`
	Operator Operator    `json:"operator"`
	Value    interface{} `json:"value"`
}

// Pagination represents pagination and sorting information.
type Pagination struct {
	Limit      int    `json:"limit"`
	Page       int    `json:"page"`
	SortBy     string `json:"sortBy,omitempty"`
	SortOrder  string `json:"sortOrder,omitempty"` // "ASC" or "DESC"
	Total      int64  `json:"total,omitempty"`
	TotalPages int    `json:"totalPages,omitempty"`
	PrevPage   int    `json:"prevPage,omitempty"`
	NextPage   int    `json:"nextPage,omitempty"`
}

// ListRequest encapsulates filters and pagination for listing entities.
type ListRequest struct {
	Filters    []Filter   `json:"filters"`
	Pagination Pagination `json:"pagination"`
}

// ListResponse represents the response for a list operation.
type ListResponse[T any] struct {
	Data       []T        `json:"data"`
	Pagination Pagination `json:"pagination"`
}
