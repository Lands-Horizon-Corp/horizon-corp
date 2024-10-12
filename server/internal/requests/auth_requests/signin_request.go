package auth_requests

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type SignInRequest struct {
	Email       string `json:"email" validate:"omitempty,email,max=255"`
	Username    string `json:"username" validate:"omitempty,max=255"`
	Password    string `json:"password" validate:"required,min=8,max=255"`
	AccountType string `json:"accountType" validate:"required"`
}

func (r *SignInRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.Email == "" && r.Username == "" {
		return errors.New("either email or username is required")
	}

	return nil
}
