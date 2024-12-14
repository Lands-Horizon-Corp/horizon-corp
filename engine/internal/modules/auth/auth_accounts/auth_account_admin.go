package auth_accounts

import (
	"fmt"
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/gin-gonic/gin"
)

func (ac *AuthAccount) AdminSignUp(ctx *gin.Context, emailTemplate, contactTemplate string) {
	const accountType = "Admin"
	var req models.AdminRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := ac.modelResource.ValidateAdminRequest(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user := &models.Admin{
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		MiddleName:        req.MiddleName,
		PermanentAddress:  req.PermanentAddress,
		Description:       "",
		BirthDate:         req.BirthDate,
		Username:          req.Username,
		Email:             req.Email,
		Password:          req.Password,
		IsEmailVerified:   false,
		IsContactVerified: false,
		ContactNumber:     req.ContactNumber,
		MediaID:           nil,
		Status:            "Pending",
	}

	id, err := ac.Create(user, accountType)
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
	if err := ac.otpProvider.SendEmailOTP(accountType, user.ID, emailReq); err != nil {
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
	if err := ac.otpProvider.SendContactNumberOTP(accountType, user.ID, contactReq); err != nil {
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
	ctx.JSON(http.StatusCreated, user)

}

func (ac *AuthAccount) AdminSignIn(ctx *gin.Context)                        {}
func (ac *AuthAccount) AdminForgotPassword(ctx *gin.Context)                {}
func (ac *AuthAccount) AdminChangePassword()                                {}
func (ac *AuthAccount) AdminVerifyResetLink(ctx *gin.Context)               {}
func (ac *AuthAccount) AdminSignOut(ctx *gin.Context)                       {}
func (ac *AuthAccount) AdminCurrentUser(ctx *gin.Context)                   {}
func (ac *AuthAccount) AdminNewPassword(ctx *gin.Context)                   {}
func (ac *AuthAccount) AdminSkipVerification(ctx *gin.Context)              {}
func (ac *AuthAccount) AdminSendEmailVerification(ctx *gin.Context)         {}
func (ac *AuthAccount) AdminVerifyEmail(ctx *gin.Context)                   {}
func (ac *AuthAccount) AdminSendContactNumberVerification(ctx *gin.Context) {}
func (ac *AuthAccount) AdminVerifyContactNumber(ctx *gin.Context)           {}
