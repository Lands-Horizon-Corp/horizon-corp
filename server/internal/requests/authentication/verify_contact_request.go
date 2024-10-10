package requests

import (
	"github.com/go-playground/validator/v10"
)

type VerifyContactNumberRequest struct {
	ContactNumber string `json:"contactNumber" validate:"required,email,max=255"`
	Otp           string `json:"otp" validate:"required,max=255"`
}

func (r *VerifyContactNumberRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
