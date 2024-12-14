package auth

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	cfg    *config.AppConfig
	engine *providers.EngineService
	middle *middleware.Middleware
}

func NewAuthService(
	cfg *config.AppConfig,
	engine *providers.EngineService,
	middle *middleware.Middleware,
) *AuthService {
	return &AuthService{
		cfg:    cfg,
		engine: engine,
		middle: middle,
	}
}

func (as AuthService) SignUp(ctx *gin.Context) {
	// Admin
	// Employee
	// Owner
	// Member
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
