package profile_requests

import (
	"github.com/go-playground/validator/v10"
)

type ChangeEmailRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email"`
}

func (r *ChangeEmailRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

type ChangeContactNumberRequest struct {
	Password      string `json:"password" validate:"required,min=8,max=255"`
	ContactNumber string `json:"contactNumber" validate:"required,min=10,max=15"`
}

// Validate validates the ChangeContactNumberRequest fields
func (r *ChangeContactNumberRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

type ChangeUsernameRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Username string `json:"username" validate:"required,min=5,max=255"`
}

func (r *ChangeUsernameRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}
