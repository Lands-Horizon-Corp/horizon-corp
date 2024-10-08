package requests

import "github.com/go-playground/validator/v10"

type VerifyContactRequest struct {
	OTP string `json:"otp" validate:"required"`
}

func (r *VerifyContactRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
