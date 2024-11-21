package requests

import "github.com/go-playground/validator"

type Validatable interface {
	Validate() error
}

func AccountTypeValidator(fl validator.FieldLevel) bool {
	// Define your account type validation logic here
	validAccountTypes := []string{"user", "admin", "guest"} // Example valid types
	for _, validType := range validAccountTypes {
		if fl.Field().String() == validType {
			return true
		}
	}
	return false
}
