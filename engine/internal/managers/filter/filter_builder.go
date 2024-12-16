package filter

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func ApplyFilters(db *gorm.DB, request PaginatedRequest) *gorm.DB {
	if strings.ToLower(request.Logic) == "or" {
		db = db.Where(func(tx *gorm.DB) *gorm.DB {
			for i, filter := range request.Filters {
				if i == 0 {
					tx = applyFiltering(tx, filter)
				} else {
					tx = tx.Or(func(orTx *gorm.DB) *gorm.DB {
						return applyFiltering(orTx, filter)
					})
				}
			}
			return tx
		})
	} else {
		for _, filter := range request.Filters {
			db = applyFiltering(db, filter)
		}
	}
	for _, preload := range request.Preloads {
		db = db.Preload(preload)
	}
	return db
}

func applyFiltering(db *gorm.DB, filter Filter) *gorm.DB {
	value := filter.GetValue()
	dataType := FilterDataType(filter.GetDataType())
	if filter.IsMultiple() {
		values, ok := toSlice(value)
		if !ok {
			return db
		}

		convertedValues := convertValues(dataType, values)
		return db.Where(func(db *gorm.DB) *gorm.DB {
			for i, v := range convertedValues {
				if i == 0 {
					db = filtering(db, filter, v)
				} else {
					db = db.Or(func(db *gorm.DB) *gorm.DB {
						return filtering(db, filter, v)
					})
				}
			}
			return db
		})
	}
	return filtering(db, filter, convertValue(dataType, value))
}

func filtering(db *gorm.DB, filter Filter, value FilterValue) *gorm.DB {
	field := sanitizeField(filter.GetField())
	mode := filter.GetMode()
	dataType := FilterDataType(filter.GetDataType())

	switch mode {
	case ModeEqual:
		return db.Where(fmt.Sprintf("%s = ?", field), value)
	case ModeNotEqual:
		return db.Where(fmt.Sprintf("%s != ?", field), value)
	case ModeContains:
		if dataType != DataTypeText {
			return db
		}
		return db.Where(fmt.Sprintf("%s LIKE ?", field), fmt.Sprintf("%%%v%%", value))
	case ModeNotContains:
		if dataType != DataTypeText {
			return db
		}
		return db.Where(fmt.Sprintf("%s NOT LIKE ?", field), fmt.Sprintf("%%%v%%", value))
	case ModeStartsWith:
		if dataType != DataTypeText {
			return db
		}
		return db.Where(fmt.Sprintf("%s LIKE ?", field), fmt.Sprintf("%v%%", value))
	case ModeEndsWith:
		if dataType != DataTypeText {
			return db
		}
		return db.Where(fmt.Sprintf("%s LIKE ?", field), fmt.Sprintf("%%%v", value))
	case ModeGreaterThan:
		if !isComparable(dataType) {
			return db
		}
		return db.Where(fmt.Sprintf("%s > ?", field), value)
	case ModeGreaterEqual:
		if !isComparable(dataType) {
			return db
		}
		return db.Where(fmt.Sprintf("%s >= ?", field), value)
	case ModeLessThan:
		if !isComparable(dataType) {
			return db
		}
		return db.Where(fmt.Sprintf("%s < ?", field), value)
	case ModeLessEqual:
		if !isComparable(dataType) {
			return db
		}
		return db.Where(fmt.Sprintf("%s <= ?", field), value)
	case ModeIsEmpty:
		return db.Where(fmt.Sprintf("%s IS NULL OR %s = ?", field, field), "")
	case ModeIsNotEmpty:
		return db.Where(fmt.Sprintf("%s IS NOT NULL AND %s != ?", field, field), "")
	case ModeRange, ModeBetween:
		rangeVal, ok := value.(RangeValue)
		if !ok {
			return db
		}
		return db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", field), rangeVal.From, rangeVal.To)
	case ModeBefore:
		if !isComparable(dataType) {
			return db
		}
		return db.Where(fmt.Sprintf("%s < ?", field), value)
	case ModeAfter:
		if !isComparable(dataType) {
			return db
		}
		return db.Where(fmt.Sprintf("%s > ?", field), value)
	case ModeBooleanTrue:
		if dataType != DataTypeBoolean {
			return db
		}
		return db.Where(fmt.Sprintf("%s = ?", field), true)
	case ModeBooleanFalse:
		if dataType != DataTypeBoolean {
			return db
		}
		return db.Where(fmt.Sprintf("%s = ?", field), false)
	default:
		return db
	}
}
