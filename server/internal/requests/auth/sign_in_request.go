package requests

import (
	"github.com/go-playground/validator/v10"
)

type VerifyEmailRequest struct {
	Token string `json:"token" validate:"required,max=500"`
	Email string `json:"email" validate:"required,email,max=255"`
}

func (r *VerifyEmailRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
