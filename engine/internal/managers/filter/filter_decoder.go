package filter

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const maxDecodedLength = 1024 * 8192

func DecodeBase64JSON(encoded string, v interface{}) error {
	if encoded == "" {
		return errors.New("encoded string is empty")
	}
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return fmt.Errorf("failed to decode Base64: %w", err)
	}
	if len(decoded) > maxDecodedLength {
		return errors.New("decoded string exceeds maximum allowed size")
	}
	if containsSQLInjectionRisk(string(decoded)) {
		return errors.New("decoded string contains potentially dangerous SQL injection patterns")
	}

	var temp map[string]interface{}
	err = json.Unmarshal(decoded, &temp)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	err = json.Unmarshal(decoded, v)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into target: %w", err)
	}

	return nil
}

func containsSQLInjectionRisk(value string) bool {
	injectionPattern := regexp.MustCompile(`(?i)\b(SELECT|INSERT|DELETE|UPDATE|DROP|TRUNCATE|ALTER|CREATE|--|;)\b`)
	return injectionPattern.MatchString(value)
}

func convertToTypedSlice(value FilterValue, targetType FilterDataType) []FilterValue {
	switch targetType {
	case DataTypeBoolean:
		if slice, ok := value.([]FilterValue); ok {
			result := make([]FilterValue, len(slice))
			for i, v := range slice {
				if b, ok := v.(bool); ok {
					result[i] = b
				} else if s, ok := v.(string); ok {
					if parsed, err := strconv.ParseBool(s); err == nil {
						result[i] = parsed
					}
				}
			}
			return result
		}

	case DataTypeText:
		if slice, ok := value.([]FilterValue); ok {
			result := make([]FilterValue, len(slice))
			for i, v := range slice {
				if s, ok := v.(string); ok {
					result[i] = s
				} else {
					result[i] = fmt.Sprintf("%v", v)
				}
			}
			return result
		}

	case DataTypeFloat:
		if slice, ok := value.([]FilterValue); ok {
			result := make([]FilterValue, len(slice))
			for i, v := range slice {
				if f, ok := v.(float64); ok {
					result[i] = f
				} else if s, ok := v.(string); ok {
					if parsed, err := strconv.ParseFloat(s, 64); err == nil {
						result[i] = parsed
					}
				}
			}
			return result
		}

	case DataTypeNumber:
		if slice, ok := value.([]FilterValue); ok {
			result := make([]FilterValue, len(slice))
			for i, v := range slice {
				if f, ok := v.(float64); ok {
					result[i] = int(f)
				} else if s, ok := v.(string); ok {
					if parsed, err := strconv.Atoi(s); err == nil {
						result[i] = parsed
					}
				}
			}
			return result
		}
	}

	return value.([]FilterValue)
}
