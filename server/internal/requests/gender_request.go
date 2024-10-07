package requests

import "github.com/go-playground/validator"

type GenderRequest struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description" validate:"omitempty,max=500"`
}

// Validate validates the GenderRequest fields.
func (r *GenderRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
