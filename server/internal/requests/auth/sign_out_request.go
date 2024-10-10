package requests

import (
	"github.com/go-playground/validator/v10"
)

type SignOutRequest struct {
	Token string `json:"token" validate:"required,max=500"` // Adjust the max length as needed
}

func (r *SignOutRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
