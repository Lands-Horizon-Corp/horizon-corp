package auth_requests

import (
	"github.com/go-playground/validator/v10"
)

type SendEmailVerificationRequest struct{}

func (r *SendEmailVerificationRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
