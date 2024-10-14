package auth_requests

import "github.com/go-playground/validator/v10"

type SendContactNumberVerificationRequest struct {
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

func (r *SendContactNumberVerificationRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
