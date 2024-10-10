package requests

import (
	"github.com/go-playground/validator/v10"
)

type ChangePasswordRequest struct {
}

func (r *ChangePasswordRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
