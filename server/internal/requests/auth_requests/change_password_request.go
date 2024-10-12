package auth_requests

import (
	"github.com/go-playground/validator/v10"
)

type ChangePasswordRequest struct {
	NewPassword     string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

func (r *ChangePasswordRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
