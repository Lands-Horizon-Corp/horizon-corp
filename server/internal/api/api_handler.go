package api

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/controllers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/middleware"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"go.uber.org/fx"
)

func NewAPIHandlerInvoke(
	lc fx.Lifecycle,
	cfg *config.AppConfig,
	logger *providers.LoggerService,
	cache *providers.CacheService,
	engineService *providers.EngineService,
	middle *middleware.Middleware,

	// Controllers
	adminController *controllers.AdminController,
	authController *controllers.AuthController,
	controller *controllers.Controller,
	profileController *controllers.ProfileController,

) {
	router := engineService.Client
	registerMiddleware(router, middle, cfg)
	router.GET("/", controller.Index)
	router.GET("/favicon.ico", controller.Favico)
	router.GET("/ping", controller.Ping)
	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", authController.SignUp)
			auth.POST("/signin", authController.SignIn)
			auth.POST("/forgot-password", authController.ForgotPassword)
			auth.POST("/change-password", authController.ChangePassword)
			auth.GET("/verify-reset-link/:id", authController.VerifyResetLink)
			auth.POST("/signout", authController.SignOut)
			auth.Use(middle.AuthMiddleware())
			{
				auth.GET("/current-user", authController.CurrentUser)
				auth.POST("/new-password", authController.NewPassword)
				auth.POST("/skip-verification", authController.SkipVerification)
				auth.POST("/send-email-verification", authController.SendEmailVerification)
				auth.POST("/verify-email", authController.VerifyEmail)
				auth.POST("/send-contact-number-verification", authController.SendContactNumberVerification)
				auth.POST("/verify-contact-number", authController.VerifyContactNumber)
			}
		}
		admin := v1.Group("/admin")
		{
			admin.GET("/", adminController.Index)
			admin.GET("/:id", adminController.Show)
			admin.POST("/", adminController.Store)
			admin.PUT("/:id", adminController.Update)
			admin.DELETE("/:id", adminController.Destroy)
		}
		profile := v1.Group("/profile")
		{
			profile.POST("/profile-picture", profileController.ProfilePicture)
			profile.POST("/account-setting", profileController.ProfileAccountSetting)
			profile.POST("/change-email", profileController.ProfileChangeEmail)
			profile.POST("/change-contact-number", profileController.ProfileChangeContactNumber)
			profile.POST("/change-username", profileController.ProfileChangeUsername)
			profile.POST("/change-password", profileController.ProfileChangePassword)
		}
	}
	runServer(lc, cfg, engineService, logger)
}
