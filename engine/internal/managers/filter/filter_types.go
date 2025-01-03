package filter

import "strings"

type FilterMode string

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

type FilterDataType string

const (
	DataTypeText     FilterDataType = "text"
	DataTypeNumber   FilterDataType = "number"
	DataTypeFloat    FilterDataType = "float"
	DataTypeBoolean  FilterDataType = "boolean"
	DataTypeDate     FilterDataType = "date"
	DataTypeDatetime FilterDataType = "datetime"
	DataTypeTime     FilterDataType = "time"
)

type FilterValue = any

type RangeValue struct {
	From FilterValue `json:"from"`
	To   FilterValue `json:"to"`
}

type Filter interface {
	GetField() string
	GetMode() FilterMode
	GetDataType() string
	GetValue() FilterValue
	IsMultiple() bool
}

type FilterStruct struct {
	Field    string         `json:"field"`
	Mode     FilterMode     `json:"mode"`
	DataType FilterDataType `json:"dataType"`
	Value    interface{}    `json:"value"`
}

func (f FilterStruct) GetField() string {
	if strings.Contains(f.Field, ".") {
		return f.Field
	}
	return toSnakeCase(f.Field)
}

func (f FilterStruct) GetMode() FilterMode { return f.Mode }

func (f FilterStruct) GetDataType() string { return string(f.DataType) }

func (f FilterStruct) GetValue() FilterValue { return f.Value }

func (f FilterStruct) IsMultiple() bool {
	switch f.Value.(type) {
	case []interface{}, []string, []float64, []int, []bool, []RangeValue:
		return true
	default:
		return false
	}
}

type Page struct {
	Page      string `json:"page"`
	PageIndex int    `json:"pageIndex"`
}

type FilterPages[T any] struct {
	Data      []*T   `json:"data"`
	PageIndex int    `json:"pageIndex"`
	TotalPage int    `json:"totalPage"`
	PageSize  int    `json:"pageSize"`
	TotalSize int    `json:"totalSize"`
	Pages     []Page `json:"pages"`
}

type FilterRequest struct {
	Filters []FilterStruct `json:"filters"`
	Logic   string         `json:"logic"`
}

type PaginatedRequest struct {
	Filters  []FilterStruct `json:"filters"`
	Preloads []string       `json:"preloads"`
	Logic    string         `json:"logic"`
}
