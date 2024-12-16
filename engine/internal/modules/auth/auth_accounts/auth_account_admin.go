package auth_accounts

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

func (ac *AuthAccount) AdminSignUp(ctx *gin.Context, req *models.Admin, emailTemplate string, contactTemplate string) {
	const accountType = "Admin"
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

func (ac *AuthAccount) AdminSignIn(ctx *gin.Context, key, password string) {
	const accountType = "Admin"
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

func (ac *AuthAccount) AdminForgotPassword(ctx *gin.Context, key, emailTemplate, contactTemplate string) {
	const accountType = "Admin"
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

func (ac *AuthAccount) AdminChangePassword(ctx *gin.Context, id uint, password string) {
	const accountType = "Admin"
	if err := ac.UpdatePassword(accountType, id, password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("ChangePassword: Password update error: %v", err)})
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully."})
}

func (ac *AuthAccount) AdminNewPassword(ctx *gin.Context, id uint, newPassword string) {
	const accountType = "Admin"
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

func (ac *AuthAccount) AdminSkipVerification(ctx *gin.Context, id uint) {
	const accountType = "Admin"
	resource, err := ac.UpdateVerification(accountType, id, "skip", true)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("SkipVerification: User update error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, resource)
}
func (ac *AuthAccount) AdminSendEmailVerification(ctx *gin.Context, id uint, emailTemplate string) {
	const accountType = "Admin"

	email, err := ac.GetByIDForEmail(accountType, id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("NewPassword: User not found: %v", err)})
		return
	}

	emailReq := providers.EmailRequest{
		To:      email,
		Subject: "ECOOP: Email Verification",
		Body:    emailTemplate,
	}

	if err := ac.otpProvider.SendEmailOTP(accountType, id, emailReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email verification sent successfully. Please check your inbox or spam folder."})
}

func (ac *AuthAccount) AdminVerifyEmail(ctx *gin.Context, id uint) {
	const accountType = "Admin"
	updatedUser, err := ac.UpdateVerification(accountType, id, "email", true)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("AdminVerifyEmail: User update error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func (ac *AuthAccount) AdminSendContactNumberVerification(ctx *gin.Context, id uint, contactTemplate string) {
	const accountType = "Admin"
	contact, err := ac.GetByIDForContact(accountType, id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("NewPassword: User not found: %v", err)})
		return
	}
	name, err := ac.GetByIDForName(accountType, id)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("NewPassword: User not found: %v", err)})
		return
	}
	contactReq := providers.SMSRequest{
		To:   contact,
		Body: contact,
		Vars: &map[string]string{
			"name": name,
		},
	}
	if err := ac.otpProvider.SendContactNumberOTP(accountType, id, contactReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: SMS sending error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Contact number verification OTP sent successfully."})

}
func (ac *AuthAccount) AdminVerifyContactNumber(ctx *gin.Context, id uint) {
	const accountType = "Admin"
	updatedUser, err := ac.UpdateVerification(accountType, id, "contact", true)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("User update error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func (ac *AuthAccount) AdminProfilePicture(ctx *gin.Context, id uint, mediaId uint) {
	const accountType = "Admin"
	updatedUser, err := ac.UpdateProfilePicture(accountType, id, &mediaId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("User update error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}
func (ac *AuthAccount) AdminProfileAccountSetting(ctx *gin.Context, id uint, birthDate time.Time, firstName, middleName, lastName, description, permanentAddress string) {
	const accountType = "Admin"
	updatedUser, err := ac.UpdateProfileAccountSettings(accountType, id, birthDate, firstName, middleName, lastName, description, permanentAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("User update error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}

func (ac *AuthAccount) AdminProfileChangeEmail(ctx *gin.Context, id uint, password, email string) {
	const accountType = "Admin"
	if !ac.VerifyPassword(accountType, id, password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedUser, err := ac.UpdateProfileChangeEmail(accountType, id, email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("User update error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}
func (ac *AuthAccount) AdminProfileChangeContactNumber(ctx *gin.Context, id uint, password, contactNumber string) {
	const accountType = "Admin"
	if !ac.VerifyPassword(accountType, id, password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedUser, err := ac.UpdateProfileChangeContactNumber(accountType, id, contactNumber)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("User update error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}
func (ac *AuthAccount) AdminProfileChangeUsername(ctx *gin.Context, id uint, password, userName string) {
	const accountType = "Admin"
	if !ac.VerifyPassword(accountType, id, password) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	updatedUser, err := ac.UpdateProfileChangeUsername(accountType, id, userName)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("User update error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, updatedUser)
}
func (ac *AuthAccount) AdminProfileChangePassword(ctx *gin.Context, id uint, oldPassword, newPassword string) {
	const accountType = "Admin"
	if !ac.VerifyPassword(accountType, id, oldPassword) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Wrong password"})
		return
	}
	err := ac.UpdatePassword(accountType, id, newPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("User update error: %v", err)})
		return
	}
}
