package requests

import "github.com/go-playground/validator/v10"

type SignInRequest struct {
}

func (r *SignInRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
