package filter

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// ApplyFilters applies filters and preloads to a GORM query.
func ApplyFilters(db *gorm.DB, request PaginatedRequest) *gorm.DB {
	if err := validatePaginatedRequest(request); err != nil {
		return db
	}

	if strings.ToLower(request.Logic) == "or" {
		// Group OR conditions
		var orConditions *gorm.DB
		for i, filter := range request.Filters {
			if i == 0 {
				orConditions = applyFiltering(db, filter)
			} else {
				orConditions = orConditions.Or(applyFiltering(db, filter))
			}
		}
		if orConditions != nil {
			db = db.Where(orConditions)
		}
	} else {
		fmt.Println("helo 1")
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
		values := convertToTypedSlice(value, dataType)
		convertedValues := convertValues(dataType, values)

		if len(convertedValues) > 0 {
			// Build a dynamic OR condition string
			fieldName := filter.GetField()
			queryParts := []string{}
			args := []interface{}{}

			for _, v := range convertedValues {
				queryParts = append(queryParts, fmt.Sprintf("%s = ?", fieldName))
				args = append(args, v)
			}
			return db.Where(strings.Join(queryParts, " OR "), args...)
		}
	}

	// Single value filtering
	return filtering(db, filter, convertValue(dataType, value))
}

// filtering applies a filter condition based on the mode.
func filtering(db *gorm.DB, filter Filter, value FilterValue) *gorm.DB {
	field := sanitizeField(filter.GetField())
	mode := filter.GetMode()
	dataType := FilterDataType(filter.GetDataType())

	if dataType == DataTypeTime {
		castTime := getTimeCastSyntax(db)
		field = castTime(field)

		switch v := value.(type) {
		case string:
			parsedTime, err := time.Parse("15:04:05", v)
			if err != nil {
				fmt.Printf("Invalid time format for value: %v\n", v)
				return db
			}
			value = parsedTime.Format("15:04:05")
		case time.Time:
			value = v.Format("15:04:05")
		default:
			fmt.Printf("Unsupported type for DataTypeTime: %T\n", v)
			return db
		}
	}

	if dataType == DataTypeText {
		field = fmt.Sprintf("LOWER(%s)", field)
		if strVal, ok := value.(string); ok {
			value = strings.ToLower(strVal)
		}
	}

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
