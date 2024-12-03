package helpers

import (
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
