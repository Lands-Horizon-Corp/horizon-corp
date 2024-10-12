package auth_requests

import "github.com/go-playground/validator/v10"

type SendContactNumberVerificationRequest struct{}

func (r *SendContactNumberVerificationRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
