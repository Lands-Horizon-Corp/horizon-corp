package requests

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type SignUpRequest struct {
	FirstName        string    `json:"firstName" validate:"required,max=255"`
	LastName         string    `json:"lastName" validate:"required,max=255"`
	Email            string    `json:"email" validate:"required,email,max=255"`
	Password         string    `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword  string    `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=Password"`
	Birthdate        time.Time `json:"birthdate" validate:"required"`
	ContactNumber    string    `json:"contactNumber" validate:"required,max=15"`
	PermanentAddress string    `json:"permanentAddress" validate:"required,max=500"`
}

func (r *SignUpRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
