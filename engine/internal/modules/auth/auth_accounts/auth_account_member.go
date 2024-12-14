package auth_accounts

import "github.com/gin-gonic/gin"

func (ac *AuthAccount) MemberSignUp(ctx *gin.Context)                        {}
func (ac *AuthAccount) MemberSignIn(ctx *gin.Context)                        {}
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
