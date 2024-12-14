package auth_accounts

import "github.com/gin-gonic/gin"

func (ac *AuthAccount) AdminSignUp(ctx *gin.Context)                        {}
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
