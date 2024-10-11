package auth_requests

import (
	"github.com/go-playground/validator/v10"
)

type ChangeContactNumberRequest struct {
	ContactNumber string `json:"contactNumber" validate:"required_if=Token '' min=8,max=255"`
}

func (r *ChangeContactNumberRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
