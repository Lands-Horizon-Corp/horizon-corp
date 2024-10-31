package requests

import "github.com/go-playground/validator/v10"

type RoleRequest struct {
	ID          uint   `json:"id,omitempty"`
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

func (r *RoleRequest) Validate() error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return FormatValidationError(err)
	}
	return nil
}
