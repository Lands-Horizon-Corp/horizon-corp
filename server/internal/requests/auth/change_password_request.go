package requests

import (
	"github.com/go-playground/validator/v10"
)

type ChangePasswordRequest struct {
	Token           string `json:"token" validate:"required,max=500"`                       // Token for password reset
	NewPassword     string `json:"newPassword" validate:"required,min=8,max=255"`           // Minimum length as per your policy
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=NewPassword"` // Must match NewPassword
}

func (r *ChangePasswordRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
