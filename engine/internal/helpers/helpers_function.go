package helpers

import (
	"errors"
	"fmt"
	"net/mail"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
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
