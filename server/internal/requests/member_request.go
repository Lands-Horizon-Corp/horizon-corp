package requests

import "github.com/go-playground/validator/v10"

type MemberRequest struct {
	ID               uint         `json:"id,omitempty"`
	FirstName        string       `json:"firstName" validate:"required,max=255"`
	LastName         string       `json:"lastName" validate:"required,max=255"`
	PermanentAddress string       `json:"permanentAddress" validate:"omitempty,max=500"`
	Description      string       `json:"description" validate:"omitempty,max=1000"`
	Birthdate        string       `json:"birthdate" validate:"required"`
	Username         string       `json:"username" validate:"required,max=255"`
	Email            string       `json:"email" validate:"required,email,max=255"`
	Password         string       `json:"password" validate:"required,min=8"`
	Media            MediaRequest `json:"media" validate:"required"`
}

func (r *MemberRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
