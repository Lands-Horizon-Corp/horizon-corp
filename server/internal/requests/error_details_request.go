package requests

import "github.com/go-playground/validator/v10"

type ErrorDetailsRequest struct {
	Message  string `json:"message" validate:"required,max=255"`
	Name     string `json:"name" validate:"required,max=255"`
	Stack    string `json:"stack,omitempty" validate:"omitempty,max=1000"`
	Response string `json:"response,omitempty" validate:"omitempty,max=1000"`
	Status   int    `json:"status,omitempty" validate:"omitempty,gte=100,lte=599"`
}

func (r *ErrorDetailsRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
