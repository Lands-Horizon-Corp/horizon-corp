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

type SignInRequest struct {
	Key         string `json:"key" validate:"required,max=255"`
	Password    string `json:"password" validate:"required,min=8,max=255"`
	AccountType string `json:"accountType" validate:"required"`
}

type ForgotPasswordRequest struct {
	Key             string `json:"key" validate:"required,max=255"`
	AccountType     string `json:"accountType" validate:"required,max=10"`
	EmailTemplate   string `json:"emailTemplate"`
	ContactTemplate string `json:"contactTemplate"`
}

type AuthProvider struct {
	cfg           *config.AppConfig
	cryptoHelpers *helpers.HelpersCryptography
	modelResource *models.ModelResource
}

type ChangePasswordRequest struct {
	ResetID         string `json:"resetId" validate:"required,min=8,max=255"`
	NewPassword     string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

type NewPasswordRequest struct {
	PreviousPassword string `json:"previousPassword" validate:"required,min=8,max=255"`
	NewPassword      string `json:"newPassword" validate:"required,min=8,max=255"`
	ConfirmPassword  string `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=NewPassword"`
}

type SendEmailVerificationRequest struct {
	EmailTemplate string `json:"emailTemplate" validate:"required"`
}

type VerifyEmailRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
}

type SendContactNumberVerificationRequest struct {
	ContactTemplate string `json:"contactTemplate" validate:"required"`
}

type VerifyContactNumberRequest struct {
	Otp string `json:"otp" validate:"required,len=6"`
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

func (ap *AuthProvider) ValidateSignUp(r SignUpRequest) error {
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

func (ap *AuthProvider) ValidateSignIn(r SignInRequest) error {
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

func (ap *AuthProvider) ValidateForgotpassword(r ForgotPasswordRequest) error {
	validate := validator.New()

	if err := validate.RegisterValidation("accountType", ap.AccountTypeValidator); err != nil {
		return fmt.Errorf("failed to register account type validator: %v", err)
	}

	if err := validate.Struct(r); err != nil {
		return err
	}

	if r.EmailTemplate == "" && r.ContactTemplate == "" {
		return fmt.Errorf("either emailTemplate or contactTemplate must be provided")
	}

	return nil

}
func (ap *AuthProvider) ValidateChangePassword(r ChangePasswordRequest) error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

func (ap *AuthProvider) ValidateNewPassword(r NewPasswordRequest) error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

func (ap *AuthProvider) ValidateSendEmailVerification(r SendEmailVerificationRequest) error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

func (ap *AuthProvider) ValidateVerifyEmailRequest(r VerifyEmailRequest) error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

func (ap *AuthProvider) ValidateSendContactNumberVerificationRequest(r SendContactNumberVerificationRequest) error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}

func (ap *AuthProvider) ValidateVerifyContactNumberRequest(r VerifyContactNumberRequest) error {
	validate := validator.New()
	err := validate.Struct(r)
	if err != nil {
		return err
	}
	return nil
}
