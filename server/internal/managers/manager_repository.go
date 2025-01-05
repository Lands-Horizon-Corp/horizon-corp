package managers

import (
	"errors"
	"reflect"
	"regexp"

	"github.com/rotisserie/eris"

	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	DB *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{DB: db}
}

func (r *Repository[T]) Create(entity *T) error {
	if err := r.DB.Create(entity).Error; err != nil {
		return handleDBError(err)
	}
	return nil
}

func (r *Repository[T]) GetByID(id uint, preloads []string) (*T, error) {
	entity := new(T)
	db := preloadDB(r.DB, preloads)
	if err := db.First(entity, id).Error; err != nil {
		return nil, handleDBError(err)
	}
	return entity, nil
}

func (r *Repository[T]) GetAll(preloads []string) ([]*T, error) {
	var entities []*T
	db := preloadDB(r.DB, preloads)
	if err := db.Find(&entities).Error; err != nil {
		return nil, handleDBError(err)
	}
	return entities, nil
}

func (r *Repository[T]) Update(entity *T, preloads []string) error {
	db := preloadDB(r.DB, preloads)
	if err := db.Save(entity).Error; err != nil {
		return handleDBError(err)
	}
	return nil
}

func (r *Repository[T]) Delete(id uint) error {
	entity := new(T)
	if err := r.DB.Delete(entity, id).Error; err != nil {
		return handleDBError(err)
	}
	return nil
}

func (r *Repository[T]) UpdateColumns(id uint, updates T, preloads []string) (*T, error) {
	entity := new(T)
	db := preloadDB(r.DB, preloads)
	if err := db.First(entity, id).Error; err != nil {
		return nil, handleDBError(err)
	}

	updateFields(entity, updates)
	if err := db.Save(entity).Error; err != nil {
		return nil, handleDBError(err)
	}

	if err := db.First(entity, id).Error; err != nil {
		return nil, handleDBError(err)
	}

	return entity, nil
}

var (
	ErrNotFound              = eris.New("we couldn't find the information you were looking for")
	ErrUniqueConstraint      = eris.New("the information you entered needs to be unique, but it already exists")
	ErrInvalidValueType      = eris.New("the type of information you provided is not correct or expected")
	ErrDuplicateEntry        = eris.New("the information you tried to add already exists")
	ErrForeignKeyViolation   = eris.New("there's a link between pieces of information, and the connection you tried to make is invalid")
	ErrTimeout               = eris.New("the action took too long and was stopped")
	ErrDataTooLong           = eris.New("the information you entered is too long for where it's being stored")
	ErrInvalidSyntax         = eris.New("there's an error in the way the request was written")
	ErrCheckConstraintFailed = eris.New("the information you provided doesn't meet the required conditions")
	ErrBadNull               = eris.New("you must provide information; it can't be left blank")
)

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

var (
	reDuplicateEntry   = regexp.MustCompile(`Duplicate entry '.*' for key '(.+)'`)
	reDataTooLong      = regexp.MustCompile(`Data too long for column '(.+)' at row`)
	reBadNull          = regexp.MustCompile(`Column '(.+)' cannot be null`)
	reInvalidValueType = regexp.MustCompile(`Out of range value for column '(.+)' at row`)
)

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

	return eris.Wrap(err, "unexpected database error")
}

func mapMySQLError(mysqlErr *mysql.MySQLError) error {
	switch mysqlErr.Number {
	case MySQLErrorDuplicateEntry:
		return eris.Wrap(ErrDuplicateEntry, extractKeyInfo(mysqlErr.Message, reDuplicateEntry))
	case MySQLErrorDuplicateEntryWithKey:
		return eris.Wrap(ErrUniqueConstraint, extractKeyInfo(mysqlErr.Message, reDuplicateEntry))
	case MySQLErrorNoReferencedRow1, MySQLErrorNoReferencedRow2:
		return ErrForeignKeyViolation
	case MySQLErrorBadNull:
		return eris.Wrap(ErrBadNull, extractColumnInfo(reBadNull, mysqlErr.Message))
	case MySQLErrorDataTooLong:
		return eris.Wrap(ErrDataTooLong, extractColumnInfo(reDataTooLong, mysqlErr.Message))
	case MySQLErrorParseError:
		return ErrInvalidSyntax
	case MySQLErrorCheckConstraintViolated:
		return ErrCheckConstraintFailed
	case MySQLErrorLockDeadlock:
		return ErrTimeout
	case MySQLErrorDataOutOfRange:
		return eris.Wrap(ErrInvalidValueType, extractColumnInfo(reInvalidValueType, mysqlErr.Message))
	default:
		return eris.Errorf("unhandled MySQL error %d: %s", mysqlErr.Number, mysqlErr.Message)
	}
}

func preloadDB(db *gorm.DB, preloads []string) *gorm.DB {
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	return db
}

func updateFields[T any](entity *T, updates T) {
	updatesValue := reflect.ValueOf(updates)
	entityValue := reflect.ValueOf(entity).Elem()
	updatesType := updatesValue.Type()

	for i := 0; i < updatesValue.NumField(); i++ {
		field := updatesType.Field(i)
		updateField := updatesValue.Field(i)
		entityField := entityValue.FieldByName(field.Name)
		if updateField.Kind() == reflect.Bool || (updateField.IsValid() && !updateField.IsZero()) {
			entityField.Set(updateField)
		}
	}
}

func extractKeyInfo(message string, re *regexp.Regexp) string {
	matches := re.FindStringSubmatch(message)
	if len(matches) > 1 {
		return stripTableName(matches[1])
	}
	return ""
}

func extractColumnInfo(re *regexp.Regexp, message string) string {
	matches := re.FindStringSubmatch(message)
	if len(matches) > 1 {
		return stripTableName(matches[1])
	}
	return ""
}

func stripTableName(name string) string {
	if dotIndex := lastDotIndex(name); dotIndex != -1 {
		return name[dotIndex+1:]
	}
	return name
}

func lastDotIndex(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			return i
		}
	}
	return -1
}
