package requests

import (
	"github.com/go-playground/validator/v10"
)

type SignOutRequest struct {
	Token string `json:"token" validate:"required,max=255"`
}

func (r *SignOutRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
