package requests

import (
	"github.com/go-playground/validator/v10"
)

type VerifyOTPRequest struct {
	ContactNumber string `json:"contactNumber" validate:"required,email,max=255"`
	Token         string `json:"token" validate:"required,max=255"`
}

func (r *VerifyOTPRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
