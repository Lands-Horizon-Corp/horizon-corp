package profile_requests

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type AccountSettingRequest struct {
	Birthdate        time.Time `json:"birthDate" validate:"required"`
	FirstName        string    `json:"firstName" validate:"required,max=255"`
	MiddleName       string    `json:"middleName" validate:"required,max=255"`
	LastName         string    `json:"lastName" validate:"required,max=255"`
	Description      string    `json:"description" validate:"max=255"`
	PermanentAddress string    `json:"permanentAddress" validate:"required,max=500"`
}

func (r *AccountSettingRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
