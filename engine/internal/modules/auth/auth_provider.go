package auth

import (
	"fmt"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/helpers"
	"github.com/go-playground/validator"
)

type SignUpRequest struct {
	AccountType     string `json:"accountType" validate:"required,max=10"`
	Password        string `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=Password"`
	EmailTemplate   string `json:"emailTemplate" validate:"required"`
	ContactTemplate string `json:"contactTemplate" validate:"required"`
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
