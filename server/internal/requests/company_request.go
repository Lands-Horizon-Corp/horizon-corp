package requests

import (
	"github.com/go-playground/validator/v10"
)

type CompanyRequest struct {
	ID              uint    `json:"id,omitempty"`
	Name            string  `json:"name" validate:"required,max=255"`
	Description     string  `json:"description,omitempty" validate:"omitempty,max=3000"`
	Address         string  `json:"address,omitempty" validate:"omitempty,max=500"`
	Longitude       float64 `json:"longitude,omitempty" validate:"omitempty,latitudeLongitude"`
	Latitude        float64 `json:"latitude,omitempty" validate:"omitempty,latitudeLongitude"`
	ContactNumber   string  `json:"contactNumber" validate:"required,max=255"`
	OwnerID         *uint   `json:"ownerId,omitempty"`
	MediaID         *uint   `json:"mediaId,omitempty"`
	IsAdminVerified bool    `json:"isAdminVerified,omitempty"`
}

func (r *CompanyRequest) Validate() error {
	validate := validator.New()

	// Register custom validation for latitude and longitude if needed
	validate.RegisterValidation("latitudeLongitude", func(fl validator.FieldLevel) bool {
		// Simple validation: longitude should be between -180 and 180, latitude between -90 and 90
		val := fl.Field().Float()
		if fl.FieldName() == "longitude" {
			return val >= -180 && val <= 180
		}
		if fl.FieldName() == "latitude" {
			return val >= -90 && val <= 90
		}
		return true
	})

	err := validate.Struct(r)
	if err != nil {
		return FormatValidationError(err)
	}
	return nil
}
