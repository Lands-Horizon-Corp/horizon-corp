package auth_accounts

import (
	"fmt"
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

func (ac *AuthAccount) MemberSignUp(ctx *gin.Context, req *models.Member, emailTemplate string, contactTemplate string) {
	const accountType = "Member"
	id, err := ac.Create(req, accountType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	token, err := ac.GenerateUserToken(id, accountType)
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

func (ac *AuthAccount) MemberSignIn(ctx *gin.Context, key, password string) {
	const accountType = "Member"
	userID, dbPassword, err := ac.FindByEmailUsernameOrContactForPassword(accountType, key)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("SignIn: User not found: %v", err)})
		return
	}
	if !ac.cryptoHelpers.VerifyPassword(dbPassword, password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "SignIn: Invalid credentials."})
		return
	}
	token, err := ac.GenerateUserToken(userID, accountType)
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
func (ac *AuthAccount) MemberForgotPassword(ctx *gin.Context)                {}
func (ac *AuthAccount) MemberChangePassword(ctx *gin.Context)                {}
func (ac *AuthAccount) MemberVerifyResetLink(ctx *gin.Context)               {}
func (ac *AuthAccount) MemberSignOut(ctx *gin.Context)                       {}
func (ac *AuthAccount) MemberCurrentUser(ctx *gin.Context)                   {}
func (ac *AuthAccount) MemberNewPassword(ctx *gin.Context)                   {}
func (ac *AuthAccount) MemberSkipVerification(ctx *gin.Context)              {}
func (ac *AuthAccount) MemberSendEmailVerification(ctx *gin.Context)         {}
func (ac *AuthAccount) MemberVerifyEmail(ctx *gin.Context)                   {}
func (ac *AuthAccount) MemberSendContactNumberVerification(ctx *gin.Context) {}
func (ac *AuthAccount) MemberVerifyContactNumber(ctx *gin.Context)           {}
