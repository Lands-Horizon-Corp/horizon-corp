package requests

import (
	"github.com/go-playground/validator/v10"
)

type VerifyOTPRequest struct {
}

func (r *VerifyOTPRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
