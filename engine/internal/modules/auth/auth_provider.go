package auth

import (
	"fmt"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/go-playground/validator"
)

type SignUpRequest struct {
	AccountType      string               `json:"accountType" validate:"required,max=10"`
	FirstName        string               `json:"firstName" validate:"required,max=255"`
	LastName         string               `json:"lastName" validate:"required,max=255"`
	MiddleName       string               `json:"middleName" validate:"max=255"`
	Username         string               `json:"username" validate:"max=255"`
	Email            string               `json:"email" validate:"required,email,max=255"`
	Password         string               `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword  string               `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=Password"`
	BirthDate        time.Time            `json:"birthDate" validate:"required"`
	ContactNumber    string               `json:"contactNumber" validate:"required,max=15"`
	PermanentAddress string               `json:"permanentAddress" validate:"required,max=500"`
	Media            *models.MediaRequest `json:"media" validate:"omitempty"`
	EmailTemplate    string               `json:"emailTemplate" validate:"required"`
	ContactTemplate  string               `json:"contactTemplate" validate:"required"`
}

type AuthProvider struct {
	cfg           *config.AppConfig
	cryptoHelpers *helpers.HelpersCryptography
	modelResource *models.ModelResource
}

func NewAuthProvider(
	cfg *config.AppConfig,
	cryptoHelpers *helpers.HelpersCryptography,
	modelResource *models.ModelResource,
) *AuthProvider {
	return &AuthProvider{
		cfg:           cfg,
		cryptoHelpers: cryptoHelpers,
		modelResource: modelResource,
	}
}

func (ap *AuthProvider) AccountTypeValidator(fl validator.FieldLevel) bool {
	accountType := fl.Field().String()
	validTypes := []string{"Member", "Employee", "Admin", "Owner"}
	for _, validType := range validTypes {
		if accountType == validType {
			return true
		}
	}
	return false
}

func (ap *AuthProvider) Validate(r SignUpRequest) error {
	validate := validator.New()
	err := validate.Struct(r)
	if err := validate.RegisterValidation("accountType", ap.AccountTypeValidator); err != nil {
		return fmt.Errorf("failed to register account type validator: %v", err)
	}
	if err != nil {
		return err
	}
	return nil
}
