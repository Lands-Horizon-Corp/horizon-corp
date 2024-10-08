package requests

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type AdminRequest struct {
	ID               uint         `json:"id,omitempty"`
	FirstName        string       `json:"firstName" validate:"required,max=255"`
	LastName         string       `json:"lastName" validate:"required,max=255"`
	PermanentAddress string       `json:"permanentAddress" validate:"omitempty,max=500"`
	Description      string       `json:"description" validate:"omitempty,max=1000"`
	Birthdate        time.Time    `json:"birthdate" validate:"required"`
	Username         string       `json:"username" validate:"required,max=255"`
	Email            string       `json:"email" validate:"required,email,max=255"`
	Password         string       `json:"password" validate:"required,min=8"`
	ConfirmPassword  string       `json:"confirmPassword" validate:"required,eqfield=Password"`
	Media            MediaRequest `json:"media" validate:"required"`
}

func (r *AdminRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
