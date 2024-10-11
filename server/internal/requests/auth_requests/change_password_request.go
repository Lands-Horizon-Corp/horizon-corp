package auth_requests

import (
	"github.com/go-playground/validator/v10"
)

type ChangePasswordRequest struct {
	CurrentPassword string `json:"currentPassword,omitempty" validate:"required_if=Token '' min=8,max=255"`
	NewPassword     string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
	Token           string `json:"token,omitempty" validate:"omitempty,max=255"`
}

func (r *ChangePasswordRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
