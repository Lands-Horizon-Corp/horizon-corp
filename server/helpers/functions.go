package helpers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func GetKeyType(key string) string {

	trimmedKey := strings.TrimSpace(key)

	if trimmedKey == "" {
		return "empty"
	}
	if isValidEmail(trimmedKey) {
		return "email"
	}
	if isValidKey(trimmedKey) {
		return "key"
	}

	if isValidUsername(trimmedKey) {
		return "username"
	}
	return "unknown"
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func isValidKey(key string) bool {
	re := regexp.MustCompile(`^[0-9]+$`)
	return re.MatchString(key)
}

func isValidUsername(username string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
	return re.MatchString(username)
}

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

func DecodeBase64JSON(encoded string) (map[string]interface{}, error) {
	// Decode Base64
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("failed to decode Base64: %w", err)
	}

	// Parse JSON
	var result map[string]interface{}
	if err := json.Unmarshal(decoded, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return result, nil
}
