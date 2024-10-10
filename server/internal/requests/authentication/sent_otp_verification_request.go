package requests

import "github.com/go-playground/validator/v10"

type SendOTPVerificationRequest struct {
	ContactNumber string `json:"contactNumber" validate:"required,max=15"`
}

func (r *SendOTPVerificationRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
