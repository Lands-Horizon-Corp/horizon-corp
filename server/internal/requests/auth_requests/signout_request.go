package auth_requests

import (
	"github.com/go-playground/validator/v10"
)

type SignOutRequest struct{}

func (r *SignOutRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
