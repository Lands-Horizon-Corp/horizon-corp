package requests

import "github.com/go-playground/validator/v10"

type RolesRequest struct {
	Name               string `json:"name" validate:"required,max=255"`
	Description        string `json:"description,omitempty" validate:"omitempty,max=3000"`
	ApiKey             string `json:"api_key" validate:"required,max=255"`
	ReadRole           bool   `json:"read_role,omitempty"`
	WriteRole          bool   `json:"write_role,omitempty"`
	UpdateRole         bool   `json:"update_role,omitempty"`
	DeleteRole         bool   `json:"delete_role,omitempty"`
	ReadErrorDetails   bool   `json:"read_error_details,omitempty"`
	WriteErrorDetails  bool   `json:"write_error_details,omitempty"`
	UpdateErrorDetails bool   `json:"update_error_details,omitempty"`
	DeleteErrorDetails bool   `json:"delete_error_details,omitempty"`
	ReadGender         bool   `json:"read_gender,omitempty"`
	WriteGender        bool   `json:"write_gender,omitempty"`
	UpdateGender       bool   `json:"update_gender,omitempty"`
	DeleteGender       bool   `json:"delete_gender,omitempty"`
}

func (r *RolesRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
