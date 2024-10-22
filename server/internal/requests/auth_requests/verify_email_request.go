package auth_requests

import (
	"github.com/go-playground/validator/v10"
)

type VerifyEmailRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (r *VerifyEmailRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
