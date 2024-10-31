package auth_requests

import (
	"fmt"
	"horizon/server/internal/requests"
	"time"

	"github.com/go-playground/validator/v10"
)

type SignUpRequest struct {
	AccountType      string                 `json:"accountType" validate:"required,max=10"`
	FirstName        string                 `json:"firstName" validate:"required,max=255"`
	LastName         string                 `json:"lastName" validate:"required,max=255"`
	MiddleName       string                 `json:"middleName" validate:"max=255"`
	Username         string                 `json:"username" validate:"max=255"`
	Email            string                 `json:"email" validate:"required,email,max=255"`
	Password         string                 `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword  string                 `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=Password"`
	BirthDate        time.Time              `json:"birthDate" validate:"required"`
	ContactNumber    string                 `json:"contactNumber" validate:"required,max=15"`
	PermanentAddress string                 `json:"permanentAddress" validate:"required,max=500"`
	Media            *requests.MediaRequest `json:"media" validate:"omitempty"`
	EmailTemplate    string                 `json:"emailTemplate" validate:"required"`
	ContactTemplate  string                 `json:"contactTemplate" validate:"required"`
}

func AccountTypeValidator(fl validator.FieldLevel) bool {
	accountType := fl.Field().String()
	validTypes := []string{"Member", "Employee", "Admin", "Owner"}
	for _, validType := range validTypes {
		if accountType == validType {
			return true
		}
	}
	return false
}

func (r *SignUpRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err := validate.RegisterValidation("accountType", AccountTypeValidator); err != nil {
		return fmt.Errorf("failed to register account type validator: %v", err)
	}

	if err != nil {
		return requests.FormatValidationError(err)
	}
	return nil
}
