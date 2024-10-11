package auth_requests

import "github.com/go-playground/validator/v10"

type CurrentUserRequest struct{}

func (r *CurrentUserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
