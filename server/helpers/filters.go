package helpers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

func EncodeBase64[T any](data T) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	encodedData := base64.StdEncoding.EncodeToString(jsonData)
	return encodedData, nil
}

func DecodeBase64[T any](encodedData string, target *T) error {
	decodedData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(decodedData, target); err != nil {
		return err
	}
	return nil
}

const maxDecodedLength = 1024 * 1024

func DecodeBase64JSON(encoded string) (map[string]interface{}, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Base64: %w", err)
	}
	if len(decoded) > maxDecodedLength {
		return nil, errors.New("decoded string exceeds maximum allowed size")
	}

	decodedString := string(decoded)
	if containsSQLInjectionRisk(decodedString) {
		return nil, fmt.Errorf("decoded string contains potentially dangerous SQL injection patterns")
	}

	var result map[string]interface{}
	if err := json.Unmarshal(decoded, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return result, nil
}

func containsSQLInjectionRisk(value string) bool {
	injectionPattern := regexp.MustCompile(`(?i)\b(SELECT|INSERT|DELETE|UPDATE|DROP|TRUNCATE|ALTER|CREATE|--|;)\b`)
	return injectionPattern.MatchString(value)
}

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

// GetBase64ArrayString retrieves a slice of strings from a map
func GetBase64ArrayString(data map[string]interface{}, key string) []string {
	if rawArray, ok := data[key].([]interface{}); ok {
		stringArray := []string{}
		for _, item := range rawArray {
			if str, ok := item.(string); ok {
				stringArray = append(stringArray, str)
			}
		}
		return stringArray
	}
	return []string{}
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
		} else if unicode.IsLetter(r) || unicode.IsDigit(r) {
			sb.WriteRune(r)
		}
	}

	result := sb.String()
	if !regexp.MustCompile(`^[a-z0-9_]+$`).MatchString(result) {
		return "", fmt.Errorf("invalid output string: %s", result)
	}
	return result, nil
}
