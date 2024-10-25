package requests

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type TimeInRequest struct {
	TimeIn  time.Time    `json:"timeIn" validate:"required"`
	MediaIn MediaRequest `json:"mediaIn" validate:"required"`
}

type TimeOutRequest struct {
	TimeOut  time.Time    `json:"timeOut" validate:"required"`
	MediaOut MediaRequest `json:"mediaOut" validate:"required"`
}

func (r *TimeInRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}

func (r *TimeOutRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
