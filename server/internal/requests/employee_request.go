package requests

import "github.com/go-playground/validator/v10"

type EmployeeRequest struct {
	FirstName         string       `json:"firstName" validate:"required,max=255"`
	LastName          string       `json:"lastName" validate:"required,max=255"`
	PermanentAddress  string       `json:"permanentAddress" validate:"omitempty,max=500"`
	Description       string       `json:"description" validate:"omitempty,max=1000"`
	Birthdate         string       `json:"birthdate" validate:"required"`
	Username          string       `json:"username" validate:"required,max=255"`
	Email             string       `json:"email" validate:"required,email,max=255"`
	Password          string       `json:"password" validate:"required,min=8"`
	Media             MediaRequest `json:"media" validate:"required"`
	IsEmailVerified   bool         `json:"isEmailVerified" validate:"omitempty"`
	IsContactVerified bool         `json:"isContactVerified" validate:"omitempty"`
}

func (r *EmployeeRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
