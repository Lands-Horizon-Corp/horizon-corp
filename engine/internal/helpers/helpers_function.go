package helpers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/mail"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

// HelpersFunction provides a set of utility functions that can be used
// for file handling, sanitizing inputs, and basic data validations.
type HelpersFunction struct {
	cfg *config.AppConfig
}

// NewHelperFunction initializes and returns a new HelpersFunction instance.
func NewHelperFunction(
	cfg *config.AppConfig,
) *HelpersFunction {
	return &HelpersFunction{
		cfg: cfg,
	}
}

// UniqueFileName takes an original file name and returns a unique file name by
// appending a timestamp and a UUID. If the original name is empty, it defaults to "file".
// Example: "document.pdf" -> "document_1672934671_9f548c2e-ead7-4ceb-9fb9-52be73e6c2ff.pdf"
func (hf *HelpersFunction) UniqueFileName(originalName string) string {
	if originalName == "" {
		originalName = "file"
	}

	fileExt := filepath.Ext(originalName)
	fileBase := strings.TrimSuffix(originalName, fileExt)
	if fileBase == "" {
		fileBase = "file"
	}

	randomID := uuid.New().String()
	timestamp := time.Now().Unix()
	uniqueName := fmt.Sprintf("%s_%d_%s%s", fileBase, timestamp, randomID, fileExt)

	return uniqueName
}

// ValidateEmail checks whether the provided string is a valid email address.
// Returns an error if the email format is invalid.
func (hf *HelpersFunction) ValidateEmail(emailAddr string) error {
	_, err := mail.ParseAddress(emailAddr)
	if err != nil {
		return errors.New("invalid email address")
	}
	return nil
}

// PreviewString takes a string and a maximum length, and returns the string
// truncated to that length followed by "..." if it exceeds the max length.
func (hf *HelpersFunction) PreviewString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}

// SanitizePhoneNumber removes all non-digit characters from the given phone number,
// preserving an optional leading '+' sign if present.
func (hf *HelpersFunction) SanitizePhoneNumber(phone string) string {
	if len(phone) > 0 && phone[0] == '+' {
		return "+" + regexp.MustCompile(`\D`).ReplaceAllString(phone[1:], "")
	}
	return regexp.MustCompile(`\D`).ReplaceAllString(phone, "")
}

// SanitizeMessageVars takes a map of message variables and removes all characters
// not in the whitelist [a-zA-Z0-9\s.,!?@-] from their values.
// Returns a pointer to the sanitized map.
func (hf *HelpersFunction) SanitizeMessageVars(vars *map[string]string) *map[string]string {
	if vars == nil {
		return vars
	}
	whitelist := regexp.MustCompile(`[^a-zA-Z0-9\s.,!?@-]`)
	sanitized := make(map[string]string, len(*vars))
	for k, v := range *vars {
		sanitized[k] = whitelist.ReplaceAllString(v, "")
	}
	return &sanitized
}

// SanitizeBody removes all characters not in the whitelist [a-zA-Z0-9\s.,!?@-]
// from the given body string.
func (hf *HelpersFunction) SanitizeBody(body string) string {
	whitelist := regexp.MustCompile(`[^a-zA-Z0-9\s.,!?@-]`)
	return whitelist.ReplaceAllString(body, "")
}

func (hf *HelpersFunction) FormatValidationError(err error) error {
	var errorMessages []string
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range validationErrors {
			switch vErr.Tag() {
			case "required":
				errorMessages = append(errorMessages, fmt.Sprintf("%s is required", vErr.Field()))
			case "min":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be at least %s characters long", vErr.Field(), vErr.Param()))
			case "max":
				errorMessages = append(errorMessages, fmt.Sprintf("%s cannot exceed %s characters", vErr.Field(), vErr.Param()))
			case "email":
				errorMessages = append(errorMessages, fmt.Sprintf("invalid email format for %s", vErr.Field()))
			case "eqfield":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must match %s", vErr.Field(), vErr.Param()))
			case "oneof":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be one of: %s", vErr.Field(), vErr.Param()))
			case "alphanum":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must contain only alphanumeric characters", vErr.Field()))
			case "url":
				errorMessages = append(errorMessages, fmt.Sprintf("invalid URL format for %s", vErr.Field()))
			case "uuid":
				errorMessages = append(errorMessages, fmt.Sprintf("invalid UUID format for %s", vErr.Field()))
			case "len":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be exactly %s characters long", vErr.Field(), vErr.Param()))
			case "datetime":
				errorMessages = append(errorMessages, fmt.Sprintf("invalid datetime format for %s", vErr.Field()))
			case "numeric":
				errorMessages = append(errorMessages, fmt.Sprintf("%s must be a numeric value", vErr.Field()))
			default:
				errorMessages = append(errorMessages, fmt.Sprintf("validation failed for %s: %s", vErr.Field(), vErr.Tag()))
			}
		}
	}
	if len(errorMessages) > 0 {
		return fmt.Errorf("%s", strings.Join(errorMessages, "; "))
	}
	return err
}

func (hf *HelpersFunction) GetKeyType(key string) string {

	trimmedKey := strings.TrimSpace(key)

	if trimmedKey == "" {
		return "empty"
	}
	if hf.isValidEmail(trimmedKey) {
		return "email"
	}
	if hf.isValidKey(trimmedKey) {
		return "key"
	}

	if hf.isValidUsername(trimmedKey) {
		return "username"
	}
	return "unknown"
}

func (hf *HelpersFunction) isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func (hf *HelpersFunction) isValidKey(key string) bool {
	re := regexp.MustCompile(`^[0-9]+$`)
	return re.MatchString(key)
}

func (hf *HelpersFunction) isValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return re.MatchString(username)
}

const maxDecodedLength = 1024 * 1024

func (hf *HelpersFunction) DecodeBase64JSON(encoded string, v interface{}) error {
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

	decodedString := string(decoded)
	if hf.containsSQLInjectionRisk(decodedString) {
		return errors.New("decoded string contains potentially dangerous SQL injection patterns")
	}

	if err := json.Unmarshal(decoded, v); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

func (hf *HelpersFunction) containsSQLInjectionRisk(value string) bool {
	injectionPattern := regexp.MustCompile(`(?i)\b(SELECT|INSERT|DELETE|UPDATE|DROP|TRUNCATE|ALTER|CREATE|--|;)\b`)
	return injectionPattern.MatchString(value)
}

func (hf *HelpersFunction) GetBase64Int(data map[string]interface{}, key string, defaultValue int) int {
	if value, ok := data[key].(int); ok {
		return value
	}
	return defaultValue
}

func (hf *HelpersFunction) GetBase64String(data map[string]interface{}, key string, defaultValue string) string {
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
func (hf *HelpersFunction) GetBase64ArrayString(data map[string]interface{}, key string) []string {
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

func (hf *HelpersFunction) CamelToSnake(s string) (string, error) {
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
