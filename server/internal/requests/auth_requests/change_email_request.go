package auth_requests

import "github.com/go-playground/validator/v10"

type ChangeEmailRequest struct {
	Email         string `json:"email" validate:"required,email,max=255"`
	EmailTemplate string `json:"emailTemplate" validate:"required"`
}

func (r *ChangeEmailRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
