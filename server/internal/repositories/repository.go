package repositories

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

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
