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

	// Handle DataTypeTime as before
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

	// Handle DataTypeDate to ignore time and search the entire day
	if dataType == DataTypeDate {
		var dateValue time.Time
		switch v := value.(type) {
		case string:
			parsedDate, err := time.Parse("2006-01-02", v)
			if err != nil {
				fmt.Printf("Invalid date format for value: %v\n", v)
				return db
			}
			dateValue = parsedDate
		case time.Time:
			dateValue = v
		default:
			fmt.Printf("Unsupported type for DataTypeDate: %T\n", v)
			return db
		}

		startOfDay := time.Date(dateValue.Year(), dateValue.Month(), dateValue.Day(), 0, 0, 0, 0, dateValue.Location())
		endOfDay := time.Date(dateValue.Year(), dateValue.Month(), dateValue.Day(), 23, 59, 59, 999999999, dateValue.Location())

		if mode == ModeEqual || mode == ModeNotEqual {
			if mode == ModeEqual {
				return db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", field), startOfDay, endOfDay)
			} else {
				return db.Where(fmt.Sprintf("%s NOT BETWEEN ? AND ?", field), startOfDay, endOfDay)
			}
		}

		switch mode {
		case ModeGreaterThan, ModeAfter:
			return db.Where(fmt.Sprintf("%s > ?", field), endOfDay)
		case ModeGreaterEqual:
			return db.Where(fmt.Sprintf("%s >= ?", field), startOfDay)
		case ModeLessThan, ModeBefore:
			return db.Where(fmt.Sprintf("%s < ?", field), startOfDay)
		case ModeLessEqual:
			return db.Where(fmt.Sprintf("%s <= ?", field), endOfDay)
		case ModeRange, ModeBetween:
			rangeVal, ok := value.(RangeValue)
			if !ok {
				return db
			}
			// Assume rangeVal.From and rangeVal.To are dates
			var fromDate, toDate time.Time
			switch fv := rangeVal.From.(type) {
			case string:
				parsedFrom, err := time.Parse("2006-01-02", fv)
				if err != nil {
					fmt.Printf("Invalid date format for range from: %v\n", fv)
					return db
				}
				fromDate = parsedFrom
			case time.Time:
				fromDate = fv
			default:
				fmt.Printf("Unsupported type for range.From in DataTypeDate: %T\n", fv)
				return db
			}
			switch tv := rangeVal.To.(type) {
			case string:
				parsedTo, err := time.Parse("2006-01-02", tv)
				if err != nil {
					fmt.Printf("Invalid date format for range to: %v\n", tv)
					return db
				}
				toDate = parsedTo
			case time.Time:
				toDate = tv
			default:
				fmt.Printf("Unsupported type for range.To in DataTypeDate: %T\n", tv)
				return db
			}

			// Define the full range for the 'to' date
			startOfFromDay := time.Date(fromDate.Year(), fromDate.Month(), fromDate.Day(), 0, 0, 0, 0, fromDate.Location())
			endOfToDay := time.Date(toDate.Year(), toDate.Month(), toDate.Day(), 23, 59, 59, 999999999, toDate.Location())

			return db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", field), startOfFromDay, endOfToDay)
		default:
			// For other modes, you can define behavior as needed
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
