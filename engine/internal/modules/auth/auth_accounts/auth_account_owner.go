package auth_accounts

import "github.com/gin-gonic/gin"

func (ac *AuthAccount) OwnerSignUp(ctx *gin.Context)                        {}
func (ac *AuthAccount) OwnerSignIn(ctx *gin.Context)                        {}
func (ac *AuthAccount) OwnerForgotPassword(ctx *gin.Context)                {}
func (ac *AuthAccount) OwnerChangePassword(ctx *gin.Context)                {}
func (ac *AuthAccount) OwnerVerifyResetLink(ctx *gin.Context)               {}
func (ac *AuthAccount) OwnerSignOut(ctx *gin.Context)                       {}
func (ac *AuthAccount) OwnerCurrentUser(ctx *gin.Context)                   {}
func (ac *AuthAccount) OwnerNewPassword(ctx *gin.Context)                   {}
func (ac *AuthAccount) OwnerSkipVerification(ctx *gin.Context)              {}
func (ac *AuthAccount) OwnerSendEmailVerification(ctx *gin.Context)         {}
func (ac *AuthAccount) OwnerVerifyEmail(ctx *gin.Context)                   {}
func (ac *AuthAccount) OwnerSendContactNumberVerification(ctx *gin.Context) {}
func (ac *AuthAccount) OwnerVerifyContactNumber(ctx *gin.Context)           {}
