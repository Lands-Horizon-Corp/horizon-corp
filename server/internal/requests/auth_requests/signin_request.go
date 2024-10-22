package auth_requests

import (
	"github.com/go-playground/validator/v10"
)

type SignInRequest struct {
	Key         string `json:"key" validate:"required,max=255"`
	Password    string `json:"password" validate:"required,min=8,max=255"`
	AccountType string `json:"accountType" validate:"required"`
}

func (r *SignInRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}
