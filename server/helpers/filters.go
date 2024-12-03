package helpers

import (
	"fmt"
	"regexp"
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

func CamelToSnake(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("input string is empty")
	}

	validInput := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !validInput.MatchString(s) {
		return "", fmt.Errorf("invalid input string: %s", s)
	}

	var sb strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				sb.WriteRune('_')
			}
			sb.WriteRune(unicode.ToLower(r))
		} else if unicode.IsLetter(r) || unicode.IsDigit(r) { // Allow letters and digits only
			sb.WriteRune(r)
		}
	}

	result := sb.String()
	if !regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(result) {
		return "", fmt.Errorf("invalid output string: %s", result)
	}
	return result, nil
}
