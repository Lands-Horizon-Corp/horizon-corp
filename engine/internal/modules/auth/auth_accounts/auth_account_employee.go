package auth_accounts

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

func (ac *AuthAccount) EmployeeSignUp(ctx *gin.Context, req *models.Employee, emailTemplate string, contactTemplate string) {
	const accountType = "Employee"
	id, err := ac.Create(req, accountType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	token, err := ac.GenerateUserToken(id, accountType, SignedInExpiration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	emailReq := providers.EmailRequest{
		To:      req.Email,
		Subject: "ECOOP: Email Verification",
		Body:    emailTemplate,
	}
	if err := ac.otpProvider.SendEmailOTP(accountType, req.ID, emailReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	contactReq := providers.SMSRequest{
		To:   req.ContactNumber,
		Body: contactTemplate,
		Vars: &map[string]string{
			"name": fmt.Sprintf("%s %s", req.FirstName, req.LastName),
		},
	}
	if err := ac.otpProvider.SendContactNumberOTP(accountType, req.ID, contactReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     ac.cfg.AppTokenName,
		Value:    *token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})
	ctx.JSON(http.StatusCreated, req)

}
func (ac *AuthAccount) EmployeeSignIn(ctx *gin.Context, key, password string) {
	const accountType = "Employee"
	userID, dbPassword, err := ac.FindByEmailUsernameOrContactForPassword(accountType, key)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("SignIn: User not found: %v", err)})
		return
	}
	if !ac.cryptoHelpers.VerifyPassword(dbPassword, password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "SignIn: Invalid credentials."})
		return
	}
	token, err := ac.GenerateUserToken(userID, accountType, SignedInExpiration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "SignIn: Token generation error"})
		return
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     ac.cfg.AppTokenName,
		Value:    *token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})
	user, err := ac.FindByEmailUsernameOrContact(accountType, key)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("SignIn: User not found: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (ac *AuthAccount) EmployeeForgotPassword(ctx *gin.Context, key, emailTemplate, contactTemplate string) {
	const accountType = "Employee"
	user, err := ac.FindByEmailUsernameOrContactForID(accountType, key)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("SignIn: User not found: %v", err)})
		return
	}

	token, err := ac.GenerateUserToken(user, accountType, time.Minute*10)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	resetLink := fmt.Sprintf("%s/auth/password-reset/%s", ac.cfg.AppClientUrl, *token)
	keyType := ac.helpers.GetKeyType(key)

	name, err := ac.FindByEmailUsernameOrContactForName(accountType, key)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("SignIn: User not found: %v", err)})
		return
	}

	switch keyType {
	case "contact":
		contactReq := providers.SMSRequest{
			To:   key,
			Body: contactTemplate,
			Vars: &map[string]string{
				"name":      name,
				"eventLink": resetLink,
			},
		}

		if err := ac.smsProvider.SendSMS(contactReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ForgotPassword: SMS sending error %v", err)})
		}
	case "email":
		emailReq := providers.EmailRequest{
			To:      key,
			Subject: "ECOOP: Change Password Request",
			Body:    emailTemplate,
			Vars: &map[string]string{
				"name":      name,
				"eventLink": resetLink,
			},
		}

		if err := ac.emailProvider.SendEmail(emailReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ForgotPassword: Email sending error: %v", err)})
			return
		}
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ForgotPassword: Invalid key type: %s", keyType)})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "If the account exists, password reset instructions have been sent."})
}

func (ac *AuthAccount) EmployeeChangePassword(ctx *gin.Context, id uint, password string) {
	const accountType = "Employee"
	if err := ac.UpdatePassword(accountType, id, password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ChangePassword: Password update error: %v", err)})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully."})
}

func (ac *AuthAccount) EmployeeNewPassword(ctx *gin.Context, id uint, newPassword string) {
	const accountType = "Employee"
	password, err := ac.GetByIDForPassword(accountType, id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("NewPassword: User not found: %v", err)})
		return
	}
	if ac.cryptoHelpers.VerifyPassword(password, newPassword) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "NewPassword: NewPassword: Password verification failed"})
		return
	}
	if err := ac.UpdatePassword(accountType, id, newPassword); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("NewPassword: Password update error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully."})
}

func (ac *AuthAccount) EmployeeSkipVerification(ctx *gin.Context)              {}
func (ac *AuthAccount) EmployeeSendEmailVerification(ctx *gin.Context)         {}
func (ac *AuthAccount) EmployeeVerifyEmail(ctx *gin.Context)                   {}
func (ac *AuthAccount) EmployeeSendContactNumberVerification(ctx *gin.Context) {}
func (ac *AuthAccount) EmployeeVerifyContactNumber(ctx *gin.Context)           {}
