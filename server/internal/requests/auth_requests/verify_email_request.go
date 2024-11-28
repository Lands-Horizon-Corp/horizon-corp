package auth_requests

import (
	"horizon/server/internal/requests"

	"github.com/go-playground/validator/v10"
)

type VerifyEmailRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

func (r *VerifyEmailRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return requests.FormatValidationError(err)
	}
	return nil
}
