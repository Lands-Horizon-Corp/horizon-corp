package filter

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// ----------------------
// Helper Functions
// ----------------------

// sanitizeField ensures that the field name is safe to use in queries.
// It allows only alphanumeric characters and underscores.
func sanitizeField(field string) string {
	allowed := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	var sanitized strings.Builder
	for _, char := range field {
		if strings.ContainsRune(allowed, char) {
			sanitized.WriteRune(char)
		}
	}
	return sanitized.String()
}

// toSlice attempts to convert an interface{} to a slice of FilterValue.
// Returns the slice and true if successful, otherwise nil and false.
func toSlice(value interface{}) ([]FilterValue, bool) {
	switch v := value.(type) {
	case []FilterValue:
		return v, true
	case []string:
		slice := make([]FilterValue, len(v))
		for i, val := range v {
			slice[i] = val
		}
		return slice, true
	case []int:
		slice := make([]FilterValue, len(v))
		for i, val := range v {
			slice[i] = val
		}
		return slice, true
	case []float64:
		slice := make([]FilterValue, len(v))
		for i, val := range v {
			slice[i] = val
		}
		return slice, true
	case []bool:
		slice := make([]FilterValue, len(v))
		for i, val := range v {
			slice[i] = val
		}
		return slice, true
	case []RangeValue:
		slice := make([]FilterValue, len(v))
		for i, val := range v {
			slice[i] = val
		}
		return slice, true
	default:
		return nil, false
	}
}

// convertValue converts the FilterValue to the appropriate type based on FilterDataType.
func convertValue(dataType FilterDataType, value FilterValue) FilterValue {
	switch dataType {
	case DataTypeNumber:
		return convertToInt(value)
	case DataTypeFloat:
		return convertToFloat(value)
	case DataTypeBoolean:
		return convertToBool(value)
	case DataTypeDate, DataTypeDatetime, DataTypeTime:
		return convertToTime(dataType, value)
	case DataTypeText:
		return fmt.Sprintf("%v", value)
	default:
		return value
	}
}

// convertValues converts a slice of FilterValue to their appropriate types based on FilterDataType.
func convertValues(dataType FilterDataType, values []FilterValue) []FilterValue {
	converted := make([]FilterValue, len(values))
	for i, v := range values {
		converted[i] = convertValue(dataType, v)
	}
	return converted
}

// convertToInt converts a FilterValue to int if possible.
func convertToInt(value FilterValue) FilterValue {
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		return v
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		if parsed, err := strconv.Atoi(v); err == nil {
			return parsed
		}
		log.Printf("Failed to convert value '%v' to int for DataTypeNumber", v)
		return v
	default:
		return v
	}
}

// convertToFloat converts a FilterValue to float64 if possible.
func convertToFloat(value FilterValue) FilterValue {
	switch v := value.(type) {
	case float32, float64:
		return v
	case int, int8, int16, int32, int64:
		return float64(v.(int))
	case string:
		if parsed, err := strconv.ParseFloat(v, 64); err == nil {
			return parsed
		}
		log.Printf("Failed to convert value '%v' to float for DataTypeFloat", v)
		return v
	default:
		return v
	}
}

// convertToBool converts a FilterValue to bool if possible.
func convertToBool(value FilterValue) FilterValue {
	switch v := value.(type) {
	case bool:
		return v
	case string:
		if parsed, err := strconv.ParseBool(v); err == nil {
			return parsed
		}
		log.Printf("Failed to convert value '%v' to bool for DataTypeBoolean", v)
		return v
	default:
		return v
	}
}

// convertToTime converts a FilterValue to time.Time based on the data type.
func convertToTime(dataType FilterDataType, value FilterValue) FilterValue {
	switch v := value.(type) {
	case string:
		var layout string
		switch dataType {
		case DataTypeDate:
			layout = "2006-01-02"
		case DataTypeDatetime:
			layout = "2006-01-02 15:04:05"
		case DataTypeTime:
			layout = "15:04:05"
		}
		if parsed, err := time.Parse(layout, v); err == nil {
			return parsed
		}
		log.Printf("Failed to convert value '%v' to time for DataType%s", v, dataType)
		return v
	case time.Time:
		return v
	default:
		return v
	}
}

// isComparable determines if the data type supports comparison operations.
func isComparable(dataType FilterDataType) bool {
	switch dataType {
	case DataTypeNumber, DataTypeFloat, DataTypeDate, DataTypeDatetime, DataTypeTime:
		return true
	default:
		return false
	}
}

func ValidatePaginatedRequest(request PaginatedRequest) error {
	for _, filter := range request.Filters {
		if filter.GetField() == "" {
			return errors.New("filter field cannot be empty")
		}
		if filter.GetMode() == "" {
			return errors.New("filter mode cannot be empty")
		}
		if filter.GetDataType() == "" {
			return errors.New("filter dataType cannot be empty")
		}
	}
	return nil
}
