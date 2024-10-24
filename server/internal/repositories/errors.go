package repositories

import (
	"errors"
	"fmt"
	"regexp"

	"gorm.io/gorm"

	"github.com/go-sql-driver/mysql"
)

// Custom application errors
var (
	ErrNotFound              = errors.New("we couldn't find the information you were looking for")
	ErrUniqueConstraint      = errors.New("the information you entered needs to be unique, but it already exists")
	ErrInvalidValueType      = errors.New("the type of information you provided is not correct or expected")
	ErrDuplicateEntry        = errors.New("the information you tried to add already exists")
	ErrForeignKeyViolation   = errors.New("there's a link between pieces of information, and the connection you tried to make is invalid")
	ErrTimeout               = errors.New("the action took too long and was stopped")
	ErrDataTooLong           = errors.New("the information you entered is too long for where it's being stored")
	ErrInvalidSyntax         = errors.New("there's an error in the way the request was written")
	ErrCheckConstraintFailed = errors.New("the information you provided doesn't meet the required conditions")
	ErrBadNull               = errors.New("you must provide information; it can't be left blank")
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

	return err
}

func mapMySQLError(mysqlErr *mysql.MySQLError) error {
	switch mysqlErr.Number {
	case MySQLErrorDuplicateEntry:
		keyInfo := extractKeyInfo(mysqlErr.Message, reDuplicateEntry)
		if keyInfo != "" {
			return fmt.Errorf("%w: %s", ErrDuplicateEntry, keyInfo)
		}
		return ErrDuplicateEntry

	case MySQLErrorDuplicateEntryWithKey:
		keyInfo := extractKeyInfo(mysqlErr.Message, reDuplicateEntry)
		if keyInfo != "" {
			return fmt.Errorf("%w: %s", ErrUniqueConstraint, keyInfo)
		}
		return ErrUniqueConstraint

	case MySQLErrorNoReferencedRow1, MySQLErrorNoReferencedRow2:
		return ErrForeignKeyViolation

	case MySQLErrorBadNull:
		columnInfo := extractColumnInfo(reBadNull, mysqlErr.Message)
		if columnInfo != "" {
			return fmt.Errorf("%w: %s", ErrBadNull, columnInfo)
		}
		return ErrBadNull

	case MySQLErrorDataTooLong:
		columnInfo := extractColumnInfo(reDataTooLong, mysqlErr.Message)
		if columnInfo != "" {
			return fmt.Errorf("%w: %s", ErrDataTooLong, columnInfo)
		}
		return ErrDataTooLong

	case MySQLErrorParseError:
		return ErrInvalidSyntax

	case MySQLErrorCheckConstraintViolated:
		return ErrCheckConstraintFailed

	case MySQLErrorLockDeadlock:
		return ErrTimeout

	case MySQLErrorDataOutOfRange:
		columnInfo := extractColumnInfo(reInvalidValueType, mysqlErr.Message)
		if columnInfo != "" {
			return fmt.Errorf("%w: %s", ErrInvalidValueType, columnInfo)
		}
		return ErrInvalidValueType

	default:
		return fmt.Errorf("unhandled MySQL error %d: %s", mysqlErr.Number, mysqlErr.Message)
	}
}

func extractKeyInfo(message string, re *regexp.Regexp) string {
	matches := re.FindStringSubmatch(message)
	if len(matches) > 1 {
		return stripTableName(matches[1])
	}
	return ""
}

// extractColumnInfo extracts the column name from the error message using the provided regex
func extractColumnInfo(re *regexp.Regexp, message string) string {
	matches := re.FindStringSubmatch(message)
	if len(matches) > 1 {
		return stripTableName(matches[1])
	}
	return ""
}

// stripTableName removes the table name from the key or column name if present
// For example, "users.email_unique" becomes "email_unique"
func stripTableName(name string) string {
	if dotIndex := lastDotIndex(name); dotIndex != -1 {
		return name[dotIndex+1:]
	}
	return name
}

// lastDotIndex returns the index of the last dot in the string or -1 if not found
func lastDotIndex(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			return i
		}
	}
	return -1
}
