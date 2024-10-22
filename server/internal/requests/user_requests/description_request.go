package user_requests

import (
	"github.com/go-playground/validator/v10"
)

type DescriptionRequest struct {
	Description string `json:"description" validate:"required,min=20,max=255"`
}

func (r *DescriptionRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
