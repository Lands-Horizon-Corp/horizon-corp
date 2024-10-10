package requests

import "github.com/go-playground/validator/v10"

type SendOTPVerificationRequest struct {
}

func (r *SendOTPVerificationRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
