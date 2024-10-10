package requests

import "github.com/go-playground/validator/v10"

type CurrentUserRequest struct {
	Email string `json:"email,omitempty" validate:"required_if=Token '' min=8,max=255"`
}

func (r *CurrentUserRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
