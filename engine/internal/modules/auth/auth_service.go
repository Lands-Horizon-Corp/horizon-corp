package auth

import (
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/auth/auth_accounts"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	engine       *providers.EngineService
	middle       *middleware.Middleware
	authProvider *AuthProvider
	authAccount  *auth_accounts.AuthAccount
}

func NewAuthService(
	engine *providers.EngineService,
	middle *middleware.Middleware,
	authProvider *AuthProvider,
	authAccount *auth_accounts.AuthAccount,
) *AuthService {
	return &AuthService{
		engine:      engine,
		middle:      middle,
		authAccount: authAccount,
	}
}

func (as AuthService) SignUp(ctx *gin.Context) {
	var req SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := as.authProvider.Validate(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	switch req.AccountType {
	case "Member":
		as.authAccount.MemberSignUp(ctx, req.EmailTemplate, req.ContactTemplate)
	case "Admin":
		as.authAccount.AdminSignUp(ctx, req.EmailTemplate, req.ContactTemplate)
	case "Owner":
		as.authAccount.OwnerSignUp(ctx, req.EmailTemplate, req.ContactTemplate)
	case "Employee":
		as.authAccount.EmployeeSignUp(ctx, req.EmailTemplate, req.ContactTemplate)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (as AuthService) SignIn(ctx *gin.Context) {

}

func (as AuthService) ForgotPassword(ctx *gin.Context) {

}

func (as AuthService) ChangePassword(ctx *gin.Context) {

}

func (as AuthService) VerifyResetLink(ctx *gin.Context) {

}

func (as AuthService) SignOut(ctx *gin.Context) {

}

func (as AuthService) CurrentUser(ctx *gin.Context) {

}

func (as AuthService) NewPassword(ctx *gin.Context) {

}

func (as AuthService) SkipVerification(ctx *gin.Context) {

}

func (as AuthService) SendEmailVerification(ctx *gin.Context) {

}

func (as AuthService) VerifyEmail(ctx *gin.Context) {

}

func (as AuthService) SendContactNumberVerification(ctx *gin.Context) {

}

func (as AuthService) VerifyContactNumber(ctx *gin.Context) {

}

func (as *AuthService) RegisterRoutes() {
	routes := as.engine.Client.Group("/auth")
	{
		routes.POST("/signup", as.SignUp)
		routes.POST("/signin", as.SignIn)
		routes.POST("/forgot-password", as.ForgotPassword)
		routes.POST("/change-password", as.ChangePassword)
		routes.GET("/verify-reset-link/:id", as.VerifyResetLink)
		routes.POST("/signout", as.SignOut)
		routes.Use(as.middle.AuthMiddleware())
		{
			routes.GET("/current-user", as.CurrentUser)
			routes.POST("/new-password", as.NewPassword)
			routes.POST("/skip-verification", as.SkipVerification)
			routes.POST("/send-email-verification", as.SendEmailVerification)
			routes.POST("/verify-email", as.VerifyEmail)
			routes.POST("/send-contact-number-verification", as.SendContactNumberVerification)
			routes.POST("/verify-contact-number", as.VerifyContactNumber)
		}
	}
}
