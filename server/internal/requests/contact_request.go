package requests

import (
	"github.com/go-playground/validator/v10"
)

type ContactRequest struct {
	ID            uint   `json:"id,omitempty"`
	FirstName     string `json:"firstName" validate:"required,max=255"`
	LastName      string `json:"lastName" validate:"required,max=255"`
	Email         string `json:"email" validate:"required,email,max=255"`
	Description   string `json:"description,omitempty" validate:"omitempty,max=3000"`
	ContactNumber string `json:"contactNumber,omitempty" validate:"omitempty,max=3000"`
}

func (r *ContactRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
