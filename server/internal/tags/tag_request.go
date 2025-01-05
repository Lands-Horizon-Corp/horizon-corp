package tags

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/rotisserie/eris"
)

type RequestTag struct {
	validator *validator.Validate
}

func NewRequestTag() *RequestTag {
	v := validator.New()
	v.RegisterValidation("request", validateRequestTag)
	return &RequestTag{validator: v}
}

func validateRequestTag(fl validator.FieldLevel) bool {
	field, found := fl.Parent().Type().FieldByName(fl.FieldName())
	if !found {
		return false
	}
	tagValue := field.Tag.Get("request")
	return tagValue != ""
}

func (rt *RequestTag) Validate(input interface{}) error {
	if err := rt.validator.Struct(input); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return formatValidationErrors(input, validationErrors)
		}
		return eris.Wrap(err, "validation failed")
	}
	return nil
}

func formatValidationErrors(input interface{}, validationErrors validator.ValidationErrors) error {
	inputType := reflect.TypeOf(input).Elem()
	var errorMessages []string

	for _, err := range validationErrors {
		field, _ := inputType.FieldByName(err.StructField())
		fieldName := field.Name
		tag := field.Tag.Get("request")

		// Translate validation rules into user-friendly messages
		ruleDescription := translateRule(err.Tag(), err.Param())

		// Add a readable error message for each field
		errorMessages = append(errorMessages, fmt.Sprintf(
			"Field '%s' (%s) failed validation: %s.",
			fieldName, tag, ruleDescription,
		))
	}

	// Join all messages into a single readable error
	combinedMessage := strings.Join(errorMessages, "\n")
	return eris.New(fmt.Sprintf("Validation errors:\n%s", combinedMessage))
}

// translateRule converts technical validation rules into user-friendly descriptions.
func translateRule(rule, param string) string {
	switch rule {
	case "required":
		return "is required and cannot be empty"
	case "gte":
		return fmt.Sprintf("must be greater than or equal to %s", param)
	case "gt":
		return fmt.Sprintf("must be greater than %s", param)
	case "lte":
		return fmt.Sprintf("must be less than or equal to %s", param)
	case "lt":
		return fmt.Sprintf("must be less than %s", param)
	case "min":
		return fmt.Sprintf("must have at least %s characters", param)
	case "max":
		return fmt.Sprintf("must have no more than %s characters", param)
	case "email":
		return "must be a valid email address"
	case "url":
		return "must be a valid URL"
	case "len":
		return fmt.Sprintf("must be exactly %s characters long", param)
	case "eq":
		return fmt.Sprintf("must be equal to %s", param)
	case "ne":
		return fmt.Sprintf("must not be equal to %s", param)
	case "numeric":
		return "must be a valid number"
	case "alpha":
		return "must only contain alphabetic characters"
	case "alphanum":
		return "must only contain alphanumeric characters"
	case "uuid":
		return "must be a valid UUID"
	case "boolean":
		return "must be a valid boolean (true or false)"
	case "datetime":
		return fmt.Sprintf("must be a valid datetime in the format '%s'", param)
	default:
		return fmt.Sprintf("violates the '%s' rule", rule)
	}
}
