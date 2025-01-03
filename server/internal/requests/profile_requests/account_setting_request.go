package profile_requests

import (
	"horizon/server/internal/requests"
	"time"

	"github.com/go-playground/validator/v10"
)

type AccountSettingRequest struct {
	BirthDate        time.Time `json:"birthDate" validate:"required"`
	FirstName        string    `json:"firstName" validate:"required,max=255"`
	MiddleName       string    `json:"middleName" validate:"required,max=255"`
	LastName         string    `json:"lastName" validate:"required,max=255"`
	Description      string    `json:"description" validate:"max=2048"`
	PermanentAddress string    `json:"permanentAddress" validate:"required,max=500"`
}

func (r *AccountSettingRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return requests.FormatValidationError(err)
	}
	return nil
}
