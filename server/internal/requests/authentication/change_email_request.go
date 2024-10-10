package requests

import "github.com/go-playground/validator/v10"

type ChangeEmailRequest struct {
	Email string `json:"email" validate:"required_if=Token '' min=8,max=255"`
}

func (r *ChangeEmailRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
