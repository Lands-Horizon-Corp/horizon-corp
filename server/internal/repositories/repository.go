package repositories

import (
	"errors"
	"fmt"
	"horizon/server/helpers"
	"regexp"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Page struct {
	Page      string `json:"page"`
	PageIndex int    `json:"pageIndex"`
}

type FilterPages[T any] struct {
	Data      []T    `json:"data"`
	PageIndex int    `json:"pageIndex"`
	TotalPage int    `json:"totalPage"`
	PageSize  int    `json:"pageSize"`
	TotalSize int    `json:"totalSize"`
	Pages     []Page `json:"pages"`
}

type PaginatedRequest struct {
	Filters   []Filter `json:"filters"`
	PageIndex int      `json:"pageIndex"`
	PageSize  int      `json:"pageSize"`
}

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
	ModeBefore       FilterMode = "before"
	ModeAfter        FilterMode = "after"
	ModeBooleanTrue  FilterMode = "true"
	ModeBooleanFalse FilterMode = "false"

	DataTypeText     FilterDataType = "text"
	DataTypeNumber   FilterDataType = "number"
	DataTypeFloat    FilterDataType = "float"
	DataTypeBoolean  FilterDataType = "boolean"
	DataTypeDate     FilterDataType = "date"
	DataTypeDatetime FilterDataType = "datetime"
	DataTypeTime     FilterDataType = "time"
	DataTypeEnum     FilterDataType = "enum"
)

// Base filter interface
type Filter interface {
	GetField() string
	GetMode() string
	GetDataType() string
	GetValue() interface{}
}

type FilterMode string
type FilterDataType string

// OneValueFilter, RangeFilter, EnumFilter, etc.
type OneValueFilter struct {
	Field    string         `json:"field"`
	Mode     FilterMode     `json:"mode"`
	DataType FilterDataType `json:"dataType"`
	Value    interface{}    `json:"value"`
}

func (f OneValueFilter) GetField() string {
	value, err := helpers.CamelToSnake(f.Field)
	if err != nil {
		return ""
	}
	return value
}

func (f OneValueFilter) GetMode() string       { return string(f.Mode) }
func (f OneValueFilter) GetDataType() string   { return string(f.DataType) }
func (f OneValueFilter) GetValue() interface{} { return f.Value }

// EnumFilter specifically for multiple values
type EnumFilter struct {
	Field    string         `json:"field"`
	Mode     FilterMode     `json:"mode"`
	DataType FilterDataType `json:"dataType"`
	Value    []string       `json:"value"`
}

func (f EnumFilter) GetField() string {
	value, err := helpers.CamelToSnake(f.Field)
	if err != nil {
		return ""
	}
	return value
}
func (f EnumFilter) GetMode() string       { return string(f.Mode) }
func (f EnumFilter) GetDataType() string   { return string(f.DataType) }
func (f EnumFilter) GetValue() interface{} { return f.Value }

// RangeFilter to handle range types with from and to fields
type RangeFilter struct {
	Field    string         `json:"field"`
	Mode     FilterMode     `json:"mode"`
	DataType FilterDataType `json:"dataType"`
	Value    RangeValue     `json:"value"`
}

func (f RangeFilter) GetField() string {
	value, err := helpers.CamelToSnake(f.Field)
	if err != nil {
		return ""
	}
	return value
}
func (f RangeFilter) GetMode() string       { return string(f.Mode) }
func (f RangeFilter) GetDataType() string   { return string(f.DataType) }
func (f RangeFilter) GetValue() interface{} { return f.Value }

// RangeValue struct for holding from and to values
type RangeValue struct {
	From interface{} `json:"from"`
	To   interface{} `json:"to"`
}

func extractFilters(data []interface{}) []Filter {
	var filters []Filter

	for _, rawFilter := range data {
		if filterMap, ok := rawFilter.(map[string]interface{}); ok {
			field, _ := filterMap["field"].(string)
			mode, _ := filterMap["mode"].(string)
			dataType, _ := filterMap["dataType"].(string)

			filterMode := FilterMode(mode)
			filterDataType := FilterDataType(dataType)
			value := filterMap["value"]

			switch filterDataType {
			case DataTypeText, DataTypeNumber, DataTypeFloat, DataTypeBoolean, DataTypeDate, DataTypeDatetime, DataTypeTime:
				if filterMode == ModeRange {
					if valueMap, ok := value.(map[string]interface{}); ok {
						rangeFilter := RangeFilter{
							Field:    field,
							Mode:     filterMode,
							DataType: filterDataType,
							Value: RangeValue{
								From: valueMap["from"],
								To:   valueMap["to"],
							},
						}
						filters = append(filters, rangeFilter)
					}
				} else {
					oneValueFilter := OneValueFilter{
						Field:    field,
						Mode:     filterMode,
						DataType: filterDataType,
						Value:    value,
					}
					filters = append(filters, oneValueFilter)
				}
			case DataTypeEnum:
				if enumValues, ok := value.([]interface{}); ok {
					var values []string
					for _, v := range enumValues {
						if strVal, ok := v.(string); ok {
							values = append(values, strVal)
						}
					}
					enumFilter := EnumFilter{
						Field:    field,
						Mode:     filterMode,
						DataType: filterDataType,
						Value:    values,
					}
					filters = append(filters, enumFilter)
				}
			}
		}
	}

	return filters
}

func convertFilterMode(mode string) string {
	switch mode {
	case "equal":
		return "="
	case "nequal":
		return "<>"
	case "contains":
		return "LIKE"
	case "ncontains":
		return "NOT LIKE"
	case "gt":
		return ">"
	case "gte":
		return ">="
	case "lt":
		return "<"
	case "lte":
		return "<="
	case "startswith":
		return "LIKE"
	case "endswith":
		return "LIKE"
	case "isempty":
		return "="
	case "isnotempty":
		return "<>"
	case "before":
		return "<"
	case "after":
		return ">"
	case "true":
		return "= TRUE"
	case "false":
		return "= FALSE"
	default:
		return "="
	}
}

func (c *Repository[T]) Filter(filter map[string]interface{}) (*FilterPages[*T], error) {
	var entities []*T
	db := c.DB

	pageIndex := helpers.GetBase64Int(filter, "pageIndex", 1)
	pageSize := helpers.GetBase64Int(filter, "pageSize", 10)

	if filtersData, ok := filter["filters"].([]interface{}); ok {
		filters := extractFilters(filtersData)
		paginatedRequest := PaginatedRequest{
			PageIndex: pageIndex,
			PageSize:  pageSize,
			Filters:   filters,
		}

		filters = paginatedRequest.Filters
		for _, filter := range filters {
			switch filter := filter.(type) {
			case OneValueFilter:
				db = db.Where(fmt.Sprintf("%s %s ?", filter.GetField(), convertFilterMode(filter.GetMode())), filter.GetValue())
			case RangeFilter:
				db = db.Where(fmt.Sprintf("%s >= ? AND %s <= ?", filter.GetField(), filter.GetField()), filter.Value.From, filter.Value.To)
			case EnumFilter:
				db = db.Where(fmt.Sprintf("%s IN (?)", filter.GetField()), filter.GetValue())
			}
		}

		err := db.Offset((paginatedRequest.PageIndex - 1) * paginatedRequest.PageSize).Limit(paginatedRequest.PageSize).Find(&entities).Error
		if err != nil {
			return &FilterPages[*T]{}, fmt.Errorf("failed to retrieve records: %w", err)
		}

		var totalSize int64
		db.Model(&entities).Count(&totalSize)

		totalPage := (int(totalSize) + paginatedRequest.PageSize - 1) / paginatedRequest.PageSize

		pages := make([]Page, totalPage)
		for i := 0; i < totalPage; i++ {
			pages[i] = Page{
				Page:      fmt.Sprintf("/api/page/%d", i+1),
				PageIndex: i + 1,
			}
		}
		return &FilterPages[*T]{
			Data:      entities,
			PageIndex: paginatedRequest.PageIndex,
			TotalPage: totalPage,
			PageSize:  paginatedRequest.PageSize,
			TotalSize: int(totalSize),
			Pages:     pages,
		}, nil
	}
	return nil, fmt.Errorf("%s", "something wrong extracting data")
}

// Repository is a generic repository for managing entities in the database
type Repository[T any] struct {
	DB *gorm.DB
}

// NewRepository creates a new instance of Repository
func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{DB: db}
}

// Create adds a new entity to the database
func (r *Repository[T]) Create(entity *T) error {
	return handleDBError(r.DB.Create(entity).Error)
}

// GetByID retrieves an entity by its ID, optionally preloading related data
func (r *Repository[T]) GetByID(id uint, preloads []string) (*T, error) {
	entity := new(T)
	db := r.DB
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	if err := db.First(entity, id).Error; err != nil {
		return nil, handleDBError(err)
	}
	return entity, nil
}

// GetAll retrieves all entities, optionally preloading related data
func (r *Repository[T]) GetAll(preloads []string) ([]*T, error) {
	var entities []*T
	db := r.DB
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	return entities, handleDBError(db.Find(&entities).Error)
}

// Update modifies an existing entity
func (r *Repository[T]) Update(entity *T, preloads []string) error {
	db := r.DB
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	return handleDBError(db.Save(entity).Error)
}

// Delete removes an entity by its ID
func (r *Repository[T]) Delete(id uint) error {
	entity := new(T)
	return handleDBError(r.DB.Delete(entity, id).Error)
}

// UpdateColumns performs a partial update on an entity's fields
func (r *Repository[T]) UpdateColumns(id uint, updates T, preloads []string) (*T, error) {
	entity := new(T)
	db := r.DB
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	if err := db.Model(entity).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, handleDBError(err)
	}
	if err := db.First(entity, id).Error; err != nil {
		return nil, handleDBError(err)
	}
	return entity, nil
}

// Custom error definitions
var (
	ErrNotFound              = errors.New("record not found")
	ErrUniqueConstraint      = errors.New("unique constraint violation")
	ErrInvalidValueType      = errors.New("invalid value type")
	ErrDuplicateEntry        = errors.New("duplicate entry")
	ErrForeignKeyViolation   = errors.New("foreign key constraint violation")
	ErrTimeout               = errors.New("operation timed out")
	ErrDataTooLong           = errors.New("data too long for column")
	ErrInvalidSyntax         = errors.New("syntax error")
	ErrCheckConstraintFailed = errors.New("check constraint violation")
	ErrBadNull               = errors.New("null value not allowed")
)

// MySQL error codes
const (
	MySQLErrorDuplicateEntry          = 1062
	MySQLErrorNoReferencedRow1        = 1451
	MySQLErrorNoReferencedRow2        = 1452
	MySQLErrorBadNull                 = 1048
	MySQLErrorDataTooLong             = 1406
	MySQLErrorParseError              = 1064
	MySQLErrorCheckConstraintViolated = 3819
	MySQLErrorLockDeadlock            = 1213
	MySQLErrorDataOutOfRange          = 1264
	MySQLErrorDuplicateEntryWithKey   = 1586
)

// Regex patterns for error parsing
var (
	reDuplicateEntry   = regexp.MustCompile(`Duplicate entry '.*' for key '(.+)'`)
	reDataTooLong      = regexp.MustCompile(`Data too long for column '(.+)' at row`)
	reBadNull          = regexp.MustCompile(`Column '(.+)' cannot be null`)
	reInvalidValueType = regexp.MustCompile(`Out of range value for column '(.+)' at row`)
)

// handleDBError maps database errors to custom error types
func handleDBError(err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mapMySQLError(mysqlErr)
	}
	return err
}

// mapMySQLError maps MySQL-specific errors to custom error types
func mapMySQLError(mysqlErr *mysql.MySQLError) error {
	switch mysqlErr.Number {
	case MySQLErrorDuplicateEntry:
		return formatError(ErrDuplicateEntry, mysqlErr.Message, reDuplicateEntry)
	case MySQLErrorDuplicateEntryWithKey:
		return formatError(ErrUniqueConstraint, mysqlErr.Message, reDuplicateEntry)
	case MySQLErrorNoReferencedRow1, MySQLErrorNoReferencedRow2:
		return ErrForeignKeyViolation
	case MySQLErrorBadNull:
		return formatError(ErrBadNull, mysqlErr.Message, reBadNull)
	case MySQLErrorDataTooLong:
		return formatError(ErrDataTooLong, mysqlErr.Message, reDataTooLong)
	case MySQLErrorParseError:
		return ErrInvalidSyntax
	case MySQLErrorCheckConstraintViolated:
		return ErrCheckConstraintFailed
	case MySQLErrorLockDeadlock:
		return ErrTimeout
	case MySQLErrorDataOutOfRange:
		return formatError(ErrInvalidValueType, mysqlErr.Message, reInvalidValueType)
	default:
		return fmt.Errorf("unhandled MySQL error %d: %s", mysqlErr.Number, mysqlErr.Message)
	}
}

// formatError extracts key or column information from an error message and formats it
func formatError(baseErr error, message string, re *regexp.Regexp) error {
	info := extractColumnInfo(re, message)
	if info != "" {
		return fmt.Errorf("%w: %s", baseErr, info)
	}
	return baseErr
}

// extractColumnInfo extracts the column or key name from an error message
func extractColumnInfo(re *regexp.Regexp, message string) string {
	matches := re.FindStringSubmatch(message)
	if len(matches) > 1 {
		return stripTableName(matches[1])
	}
	return ""
}

// stripTableName removes the table prefix from a key or column name
func stripTableName(name string) string {
	if dotIndex := lastDotIndex(name); dotIndex != -1 {
		return name[dotIndex+1:]
	}
	return name
}

// lastDotIndex returns the last dot index in a string or -1 if not found
func lastDotIndex(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			return i
		}
	}
	return -1
}
