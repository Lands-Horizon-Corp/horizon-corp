package auth_requests

import (
	"horizon/server/internal/requests"

	"github.com/go-playground/validator/v10"
)

type ChangePasswordRequest struct {
	ResetID         string `json:"resetId" validate:"required,min=8,max=255"`
	NewPassword     string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

func (r *ChangePasswordRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return requests.FormatValidationError(err)
	}
	return nil
}
