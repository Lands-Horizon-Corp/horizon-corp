package controllers

import "github.com/gin-gonic/gin"

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (as AuthController) SignUp(ctx *gin.Context)                        {}
func (as AuthController) SignIn(ctx *gin.Context)                        {}
func (as AuthController) ForgotPassword(ctx *gin.Context)                {}
func (as AuthController) ChangePassword(ctx *gin.Context)                {}
func (as AuthController) VerifyResetLink(ctx *gin.Context)               {}
func (as AuthController) SignOut(ctx *gin.Context)                       {}
func (as AuthController) CurrentUser(ctx *gin.Context)                   {}
func (as AuthController) NewPassword(ctx *gin.Context)                   {}
func (as AuthController) SkipVerification(ctx *gin.Context)              {}
func (as AuthController) SendEmailVerification(ctx *gin.Context)         {}
func (as AuthController) VerifyEmail(ctx *gin.Context)                   {}
func (as AuthController) SendContactNumberVerification(ctx *gin.Context) {}
func (as AuthController) VerifyContactNumber(ctx *gin.Context)           {}
