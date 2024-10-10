package requests

import (
	"github.com/go-playground/validator/v10"
)

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

func (r *SignInRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
