package api

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/controllers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/middleware"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"go.uber.org/fx"
)

var ControllerModule = fx.Module(
	"controllers",
	fx.Provide(
		controllers.NewAdminController,
		controllers.NewAuthController,
		controllers.NewBranchController,
		controllers.NewCompanyController,
		controllers.NewContactController,
		controllers.NewController,
		controllers.NewFeedbackController,
		controllers.NewFootstepController,
		controllers.NewMediaController,
		controllers.NewMemberController,
		controllers.NewProfileController,
	),
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
	branchController *controllers.BranchController,
	companyController *controllers.CompanyController,
	contactController *controllers.ContactController,
	controller *controllers.Controller,
	feedbackController *controllers.FeedbackController,
	footstepController *controllers.FootstepController,
	mediaController *controllers.MediaController,
	memberController *controllers.MemberController,
	profileController *controllers.ProfileController,

) {
	router := engineService.Client
	registerMiddleware(router, middle, cfg)
	router.GET("/", controller.Index)
	router.GET("/favicon.ico", controller.Favico)
	router.GET("/ping", controller.Ping)
	v1 := router.Group("/api/v1")
	{
		admin := v1.Group("/admin")
		{
			admin.GET("/", adminController.Index)
			admin.GET("/:id", adminController.Show)
			admin.POST("/", adminController.Store)
			admin.PUT("/:id", adminController.Update)
			admin.DELETE("/:id", adminController.Destroy)
		}
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
		branch := v1.Group("/branch")
		{
			branch.GET("/", branchController.Index)
			branch.GET("/:id", branchController.Show)
			branch.POST("/", branchController.Store)
			branch.PUT("/:id", branchController.Update)
			branch.DELETE("/:id", branchController.Destroy)
			branch.GET("/nearest-branch", branchController.NearestBranch)
			branch.GET("/nearest-company-branch/:id", branchController.NearestCompanyBranch)
		}
		company := v1.Group("/company")
		{
			company.GET("/", companyController.Index)
			company.GET("/:id", companyController.Show)
			company.POST("/", companyController.Store)
			company.PUT("/:id", companyController.Update)
			company.DELETE("/:id", companyController.Destroy)
		}
		contact := v1.Group("/contact")
		{
			contact.GET("/", contactController.Index)
			contact.GET("/:id", contactController.Show)
			contact.POST("/", contactController.Store)
			contact.DELETE("/:id", contactController.Destroy)
		}
		feedback := v1.Group("/feedback")
		{
			feedback.GET("/", feedbackController.Index)
			feedback.GET("/:id", feedbackController.Show)
			feedback.POST("/", feedbackController.Store)
			feedback.DELETE("/:id", feedbackController.Destroy)
		}
		footstep := v1.Group("/footstep")
		{
			footstep.GET("/", footstepController.Index)
			footstep.GET("/:id", footstepController.Show)
			footstep.GET("/team", footstepController.Team)
		}
		media := v1.Group("/media")
		{
			media.GET("/", mediaController.Index)
			media.GET("/:id", mediaController.Show)
			media.POST("/", mediaController.Store)
			media.PUT("/:id", mediaController.Update)
			media.DELETE("/:id", mediaController.Destroy)
			media.GET("/team", mediaController.Team)
		}
		member := v1.Group("/member")
		{
			member.GET("/", memberController.Index)
			member.GET("/:id", memberController.Show)
			member.POST("/", memberController.Store)
			member.PUT("/:id", memberController.Update)
			member.DELETE("/:id", memberController.Destroy)
			member.GET("/member-applications", memberController.Show)
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
