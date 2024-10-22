package auth_requests

import (
	"github.com/go-playground/validator/v10"
)

type VerifyContactNumberRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (r *VerifyContactNumberRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
