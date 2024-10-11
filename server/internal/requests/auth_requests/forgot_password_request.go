package auth_requests

import (
	"github.com/go-playground/validator/v10"
)

type ForgotPasswordRequest struct {
	Email         string `json:"email,omitempty" validate:"required,email,max=255"`
	ContactNumber string `json:"contactNumber,omitempty" validate:"required,contactNumber,max=255"`
}

func (r *ForgotPasswordRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
