package requests

import "github.com/go-playground/validator/v10"

type ErrorDetailRequest struct {
	ID       uint   `json:"id,omitempty"`
	Message  string `json:"message" validate:"required,max=255"`
	Name     string `json:"name" validate:"required,max=255"`
	Stack    string `json:"stack,omitempty" validate:"omitempty,max=1000"`
	Response string `json:"response,omitempty" validate:"omitempty,max=1000"`
	Status   int    `json:"status,omitempty" validate:"omitempty,gte=100,lte=599"`
}

func (r *ErrorDetailRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return FormatValidationError(err)
	}
	return nil
}
