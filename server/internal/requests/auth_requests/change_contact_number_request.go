package auth_requests

import (
	"github.com/go-playground/validator/v10"
)

type ChangeContactNumberRequest struct {
	ContactNumber   string `json:"contactNumber" validate:"required,min=8,max=255"`
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

func (r *ChangeContactNumberRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
