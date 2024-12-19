package auth_requests

import (
	"fmt"
	"horizon/server/internal/requests"

	"github.com/go-playground/validator/v10"
)

type SignInRequest struct {
	Key         string `json:"key" validate:"required,max=255"`
	Password    string `json:"password" validate:"required,min=8,max=255"`
	AccountType string `json:"accountType" validate:"required"`
}

func (r *SignInRequest) Validate() error {
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
