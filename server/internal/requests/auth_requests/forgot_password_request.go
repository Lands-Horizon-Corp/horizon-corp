package auth_requests

import (
	"fmt"
	"horizon/server/internal/requests"

	"github.com/go-playground/validator/v10"
)

type ForgotPasswordRequest struct {
	Key             string `json:"key" validate:"required,max=255"`
	AccountType     string `json:"accountType" validate:"required,max=10"`
	EmailTemplate   string `json:"emailTemplate"`
	ContactTemplate string `json:"contactTemplate"`
}

func (r *ForgotPasswordRequest) Validate() error {
	validate := validator.New()

	if err := validate.RegisterValidation("accountType", AccountTypeValidator); err != nil {
		return fmt.Errorf("failed to register account type validator: %v", err)
	}

	if err := validate.Struct(r); err != nil {
		return requests.FormatValidationError(err)
	}

	if r.EmailTemplate == "" && r.ContactTemplate == "" {
		return fmt.Errorf("either emailTemplate or contactTemplate must be provided")
	}

	return nil
}
