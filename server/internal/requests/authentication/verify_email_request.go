package requests

import (
	"github.com/go-playground/validator/v10"
)

type VerifyEmailRequest struct {
}

func (r *VerifyEmailRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
