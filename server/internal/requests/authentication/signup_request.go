package requests

import "github.com/go-playground/validator/v10"

type SignUpRequest struct {
}

func (r *SignUpRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
