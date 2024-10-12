package auth_requests

import (
	"horizon/server/internal/requests"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
)

type SignUpRequest struct {
	AccountType      string                 `json:"accountType" validate:"required,max=10"`
	FirstName        string                 `json:"firstName" validate:"required,max=255"`
	LastName         string                 `json:"lastName" validate:"required,max=255"`
	MiddleName       string                 `json:"middleName" validate:"max=255"`
	Username         string                 `json:"username" validate:"max=255"`
	Email            string                 `json:"email" validate:"required,email,max=255"`
	Password         string                 `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword  string                 `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=Password"`
	Birthdate        time.Time              `json:"birthdate" validate:"required"`
	ContactNumber    string                 `json:"contactNumber" validate:"required,max=15"`
	PermanentAddress string                 `json:"permanentAddress" validate:"required,max=500"`
	Media            *requests.MediaRequest `json:"media" validate:"omitempty"`
}

func accountTypeValidator(fl validator.FieldLevel) bool {
	accountType := fl.Field().String()
	validTypes := []string{"Member", "Employee", "Admin", "Owner"}
	for _, validType := range validTypes {
		if accountType == validType {
			return true
		}
	}
	return false
}

func (r *SignUpRequest) Validate() error {
	validate := validator.New()
	if err := validate.RegisterValidation("accounttype", accountTypeValidator); err != nil {
		log.Fatalf("Could not register validation: %v", err)
	}
	return validate.Struct(r)
}
