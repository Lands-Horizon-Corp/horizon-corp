package requests

import "github.com/go-playground/validator/v10"

type RolesRequest struct {
	Name        string `json:"name" validate:"required,max=255"`
	Description string `json:"description,omitempty" validate:"omitempty,max=3000"`
	ApiKey      string `json:"apiKey" validate:"required,max=255"`
	Color       string `json:"color" validate:"required"`

	ReadRole           bool `json:"readRole,omitempty"`
	WriteRole          bool `json:"writeRole,omitempty"`
	UpdateRole         bool `json:"updateRole,omitempty"`
	DeleteRole         bool `json:"deleteRole,omitempty"`
	ReadErrorDetails   bool `json:"readErrorDetails,omitempty"`
	WriteErrorDetails  bool `json:"writeErrorDetails,omitempty"`
	UpdateErrorDetails bool `json:"updateErrorDetails,omitempty"`
	DeleteErrorDetails bool `json:"deleteErrorDetails,omitempty"`
	ReadGender         bool `json:"readGender,omitempty"`
	WriteGender        bool `json:"writeGender,omitempty"`
	UpdateGender       bool `json:"updateGender,omitempty"`
	DeleteGender       bool `json:"deleteGender,omitempty"`
}

func (r *RolesRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
