package requests

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

// FormatValidationError formats the error messages from validation.
func FormatValidationError(err error) error {
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
