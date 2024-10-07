package requests

import "github.com/go-playground/validator"

type FeedbackRequest struct {
	Email        string `json:"email" validate:"required,max=255"`
	Description  string `json:"description" validate:"max=3000"`
	FeedbackType string `json:"feedbackType" validate:"max=255"`
}

// Validate validates the GenderRequest fields.
func (r *FeedbackRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
