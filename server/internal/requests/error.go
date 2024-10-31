package requests

import (
	"fmt"

	"github.com/go-playground/validator"
)

func FormatValidationError(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, vErr := range validationErrors {
			switch vErr.Tag() {
			case "required":
				return fmt.Errorf("%s is required", vErr.Field())
			case "min":
				return fmt.Errorf("%s must be at least %s characters long", vErr.Field(), vErr.Param())
			case "max":
				return fmt.Errorf("%s cannot exceed %s characters", vErr.Field(), vErr.Param())
			case "email":
				return fmt.Errorf("invalid email format for %s", vErr.Field())
			case "eqfield":
				return fmt.Errorf("%s must match %s", vErr.Field(), vErr.Param())
			case "oneof":
				return fmt.Errorf("%s must be one of: %s", vErr.Field(), vErr.Param())
			case "alphanum":
				return fmt.Errorf("%s must contain only alphanumeric characters", vErr.Field())
			case "url":
				return fmt.Errorf("invalid URL format for %s", vErr.Field())
			case "uuid":
				return fmt.Errorf("invalid UUID format for %s", vErr.Field())
			case "len":
				return fmt.Errorf("%s must be exactly %s characters long", vErr.Field(), vErr.Param())
			case "datetime":
				return fmt.Errorf("invalid datetime format for %s", vErr.Field())
			case "numeric":
				return fmt.Errorf("%s must be a numeric value", vErr.Field())
			default:
				return fmt.Errorf("validation failed for %s: %s", vErr.Field(), vErr.Tag())
			}
		}
	}
	return err
}
