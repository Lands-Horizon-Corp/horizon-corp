package profile_requests

import (
	"fmt"
	"horizon/server/internal/requests"

	"github.com/go-playground/validator/v10"
)

type ChangePasswordSettingRequest struct {
	OldPassword     string `json:"old_password" validate:"required,min=8,max=255"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=255"`
}

func (r *ChangePasswordSettingRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return requests.FormatValidationError(err)
	}
	if r.NewPassword != r.ConfirmPassword {
		return fmt.Errorf("new password and confirm password do not match")
	}
	return nil
}
