package auth_accounts

import (
	"fmt"
	"net/http"

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
func (ac *AuthAccount) EmployeeSignIn(ctx *gin.Context)                        {}
func (ac *AuthAccount) EmployeeForgotPassword(ctx *gin.Context)                {}
func (ac *AuthAccount) EmployeeChangePassword(ctx *gin.Context)                {}
func (ac *AuthAccount) EmployeeVerifyResetLink(ctx *gin.Context)               {}
func (ac *AuthAccount) EmployeeSignOut(ctx *gin.Context)                       {}
func (ac *AuthAccount) EmployeeCurrentUser(ctx *gin.Context)                   {}
func (ac *AuthAccount) EmployeeNewPassword(ctx *gin.Context)                   {}
func (ac *AuthAccount) EmployeeSkipVerification(ctx *gin.Context)              {}
func (ac *AuthAccount) EmployeeSendEmailVerification(ctx *gin.Context)         {}
func (ac *AuthAccount) EmployeeVerifyEmail(ctx *gin.Context)                   {}
func (ac *AuthAccount) EmployeeSendContactNumberVerification(ctx *gin.Context) {}
func (ac *AuthAccount) EmployeeVerifyContactNumber(ctx *gin.Context)           {}
