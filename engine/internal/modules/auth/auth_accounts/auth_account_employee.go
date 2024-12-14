package auth_accounts

import "github.com/gin-gonic/gin"

func (ac *AuthAccount) EmployeeSignUp(ctx *gin.Context, contextEmailTemplate, contactTemplate string) {
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
