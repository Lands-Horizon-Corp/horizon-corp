package requests

import "github.com/go-playground/validator/v10"

type SendContactNumberVerificationRequest struct {
	ContactNumber string `json:"contactNumber" validate:"required,max=15"`
}

func (r *SendContactNumberVerificationRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
