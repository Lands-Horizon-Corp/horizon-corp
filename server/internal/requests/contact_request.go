package requests

import (
	"github.com/go-playground/validator/v10"
)

type ContactsRequest struct {
	FirstName   string `json:"first_name" validate:"required,max=255"`
	LastName    string `json:"last_name" validate:"required,max=255"`
	Email       string `json:"email" validate:"required,email,max=255"`
	Description string `json:"description,omitempty" validate:"omitempty,max=3000"`
}

func (r *ContactsRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
