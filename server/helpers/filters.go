package helpers

import (
	"strings"
	"unicode"
)

func GetBase64Int(data map[string]interface{}, key string, defaultValue int) int {
	if value, ok := data[key].(int); ok {
		return value
	}
	return defaultValue
}

func GetBase64String(data map[string]interface{}, key string, defaultValue string) string {
	if value, ok := data[key].(string); ok {
		return value
	}
	return defaultValue
}

func GetBase64Filters(data map[string]interface{}, key string) []interface{} {
	if filters, ok := data[key].([]interface{}); ok {
		return filters
	}
	return []interface{}{}
}

func CamelToSnake(s string) string {
	if len(s) == 0 {
		return s
	}
	var sb strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				sb.WriteRune('_')
			}
			sb.WriteRune(unicode.ToLower(r))
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
