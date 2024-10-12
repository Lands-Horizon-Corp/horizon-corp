package auth_requests

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type ForgotPasswordRequest struct {
	Email         string `json:"email,omitempty" validate:"required,email,max=255"`
	ContactNumber string `json:"contactNumber,omitempty" validate:"required,contactNumber,max=255"`
}

func (r *ForgotPasswordRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	if r.Email == "" && r.ContactNumber == "" {
		return errors.New("either email or username is required")
	}
	return nil
}
