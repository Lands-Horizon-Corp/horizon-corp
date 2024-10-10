package requests

import (
	"github.com/go-playground/validator/v10"
)

type VerifyEmailRequest struct {
	Token       string `json:"token" validate:"required,max=500"`
	EmailNumber string `json:"emailNumber" validate:"required,max=15"`
}

func (r *VerifyEmailRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
