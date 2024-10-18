package auth_requests

import (
	"errors"

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
	if err := validate.Struct(r); err != nil {
		return err
	}
	if err := validate.RegisterValidation("accountType", AccountTypeValidator); err != nil {
		return errors.New("account type is not valid")
	}
	if r.EmailTemplate == "" && r.ContactTemplate == "" {
		return errors.New("either email or username is required")
	}
	return nil
}
