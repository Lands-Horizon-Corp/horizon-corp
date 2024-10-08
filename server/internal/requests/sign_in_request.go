// requests/sign_in_request.go
package requests

import "github.com/go-playground/validator/v10"

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (r *SignInRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
