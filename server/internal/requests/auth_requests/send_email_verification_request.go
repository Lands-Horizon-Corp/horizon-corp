package auth_requests

import (
	"horizon/server/internal/requests"

	"github.com/go-playground/validator/v10"
)

type SendEmailVerificationRequest struct {
	EmailTemplate string `json:"emailTemplate" validate:"required"`
}

func (r *SendEmailVerificationRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return requests.FormatValidationError(err)
	}
	return nil
}
