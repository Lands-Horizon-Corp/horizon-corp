package auth

import (
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/controllers/media"
)

type SignUpRequest struct {
	AccountType      string              `json:"accountType" validate:"required,max=10"`
	FirstName        string              `json:"firstName" validate:"required,max=255"`
	LastName         string              `json:"lastName" validate:"required,max=255"`
	MiddleName       string              `json:"middleName" validate:"max=255"`
	Username         string              `json:"username" validate:"max=255"`
	Email            string              `json:"email" validate:"required,email,max=255"`
	Password         string              `json:"password" validate:"required,min=8,max=255"`
	ConfirmPassword  string              `json:"confirmPassword" validate:"required,min=8,max=255,eqfield=Password"`
	BirthDate        time.Time           `json:"birthDate" validate:"required"`
	ContactNumber    string              `json:"contactNumber" validate:"required,max=15"`
	PermanentAddress string              `json:"permanentAddress" validate:"required,max=500"`
	Media            *media.MediaRequest `json:"media" validate:"omitempty"`
	EmailTemplate    string              `json:"emailTemplate" validate:"required"`
	ContactTemplate  string              `json:"contactTemplate" validate:"required"`
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

type ChangeEmailRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Email    string `json:"email" validate:"required,email"`
}

type ChangeContactNumberRequest struct {
	Password      string `json:"password" validate:"required,min=8,max=255"`
	ContactNumber string `json:"contactNumber" validate:"required,min=10,max=15"`
}

type ChangeUsernameRequest struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
	Username string `json:"username" validate:"required,min=5,max=255"`
}

type AccountSettingRequest struct {
	BirthDate        time.Time `json:"birthDate" validate:"required"`
	FirstName        string    `json:"firstName" validate:"required,max=255"`
	MiddleName       string    `json:"middleName" validate:"required,max=255"`
	LastName         string    `json:"lastName" validate:"required,max=255"`
	Description      string    `json:"description" validate:"max=2048"`
	PermanentAddress string    `json:"permanentAddress" validate:"required,max=500"`
}

type ChangePasswordSettingRequest struct {
	OldPassword     string `json:"old_password" validate:"required,min=8,max=255"`
	NewPassword     string `json:"new_password" validate:"required,min=8,max=255"`
	ConfirmPassword string `json:"confirm_password" validate:"required,min=8,max=255"`
}
