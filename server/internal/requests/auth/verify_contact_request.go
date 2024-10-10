package requests

import (
	"github.com/go-playground/validator/v10"
)

type VerifyContactRequest struct {
	Token         string `json:"token" validate:"required,max=500"`
	ContactNumber string `json:"contactNumber" validate:"required,max=15"`
}

func (r *VerifyContactRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
