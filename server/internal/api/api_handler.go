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
	branchController *controllers.BranchController,
	companyController *controllers.CompanyController,
	contactController *controllers.ContactController,
	controller *controllers.Controller,
	employeeController *controllers.EmployeeController,
	feedbackController *controllers.FeedbackController,
	footstepController *controllers.FootstepController,
	genderController *controllers.GenderController,
	mediaController *controllers.MediaController,
	memberController *controllers.MemberController,
	memberProfileController *controllers.MemberProfileController,
	ownerController *controllers.OwnerController,
	profileController *controllers.ProfileController,
	qrController *controllers.QRScannerController,
	timesheetController *controllers.TimesheetController,

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
			admin.GET("/", middle.AccountTypeMiddleware("Admin"), adminController.Index)
			admin.GET("/:id", middle.AccountTypeMiddleware("Admin"), adminController.Show)
			admin.POST("/", middle.AccountTypeMiddleware("Admin"), adminController.Store)
			admin.PUT("/:id", middle.AccountTypeMiddleware("Admin"), adminController.Update)
			admin.DELETE("/:id", middle.AccountTypeMiddleware("Admin"), adminController.Destroy)

			admin.POST("/change-password", middle.AccountTypeMiddleware("Admin"), adminController.ChangePassword)
			admin.POST("/forgot-password", middle.AccountTypeMiddleware("Admin"), adminController.ForgotPassword)
			admin.POST("/forgot-password-reset-link", middle.AccountTypeMiddleware("Admin"), adminController.ForgotPassword)
			admin.POST("/new-password", middle.AccountTypeMiddleware("Admin"), adminController.NewPassword)

			admin.POST("/skip-verification", middle.AccountTypeMiddleware("Admin"), adminController.SkipVerification)
			admin.POST("/send-email-verification", middle.AccountTypeMiddleware("Admin"), adminController.SendEmailVerification)
			admin.POST("/verify-email", middle.AccountTypeMiddleware("Admin"), adminController.VerifyEmail)

			admin.POST("/send-contact-number-verification", middle.AccountTypeMiddleware("Admin"), adminController.SendContactNumberVerification)
			admin.POST("/verify-contact-number", middle.AccountTypeMiddleware("Admin"), adminController.VerifyContactNumber)

			admin.POST("/profile-picture", middle.AccountTypeMiddleware("Admin"), adminController.ProfilePicture)
			admin.PUT("/profile-account-setting", middle.AccountTypeMiddleware("Admin"), adminController.ProfileAccountSetting)
			admin.PUT("/profile-change-email", middle.AccountTypeMiddleware("Admin"), adminController.ProfileChangeEmail)
			admin.PUT("/profile-change-contact-number", middle.AccountTypeMiddleware("Admin"), adminController.ProfileChangeContactNumber)
			admin.PUT("/profile-change-username", middle.AccountTypeMiddleware("Admin"), adminController.ProfileChangeUsername)
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

		employee := v1.Group("/employee")
		{
			employee.GET("/", middle.AccountTypeMiddleware("Admin", "Owner", "Employee"), employeeController.Index)
			employee.GET("/:id", middle.AccountTypeMiddleware("Admin", "Owner", "Employee"), employeeController.Show)
			employee.POST("/", middle.AccountTypeMiddleware("Admin", "Owner", "Employee"), employeeController.Store)
			employee.PUT("/:id", middle.AccountTypeMiddleware("Admin", "Owner", "Employee"), employeeController.Update)
			employee.DELETE("/:id", middle.AccountTypeMiddleware("Admin", "Owner"), employeeController.Destroy)

			employee.POST("/forgot-password-reset-link", middle.AccountTypeMiddleware("Admin", "Owner", "Employee"), employeeController.ForgotPassword)
			employee.POST("/change-password", middle.AccountTypeMiddleware("Employee"), employeeController.ChangePassword)
			employee.POST("/forgot-password", middle.AccountTypeMiddleware("Employee"), employeeController.ForgotPassword)
			employee.POST("/new-password", middle.AccountTypeMiddleware("Employee"), employeeController.NewPassword)

			employee.POST("/skip-verification", middle.AccountTypeMiddleware("Employee"), employeeController.SkipVerification)
			employee.POST("/send-email-verification", middle.AccountTypeMiddleware("Employee"), employeeController.SendEmailVerification)
			employee.POST("/verify-email", middle.AccountTypeMiddleware("Employee"), employeeController.VerifyEmail)

			employee.POST("/send-contact-number-verification", middle.AccountTypeMiddleware("Employee"), employeeController.SendContactNumberVerification)
			employee.POST("/verify-contact-number", middle.AccountTypeMiddleware("Employee"), employeeController.VerifyContactNumber)

			employee.POST("/profile-picture", middle.AccountTypeMiddleware("Employee"), employeeController.ProfilePicture)
			employee.PUT("/profile-account-setting", middle.AccountTypeMiddleware("Employee"), employeeController.ProfileAccountSetting)
			employee.PUT("/profile-change-email", middle.AccountTypeMiddleware("Employee"), employeeController.ProfileChangeEmail)
			employee.PUT("/profile-change-contact-number", middle.AccountTypeMiddleware("Employee"), employeeController.ProfileChangeContactNumber)
			employee.PUT("/profile-change-username", middle.AccountTypeMiddleware("Employee"), employeeController.ProfileChangeUsername)
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
		gender := v1.Group("/gender")
		{
			gender.GET("/", genderController.Index)
			gender.GET("/:id", genderController.Show)
			gender.POST("/", genderController.Store)
			gender.PUT("/:id", genderController.Update)
			gender.DELETE("/:id", genderController.Destroy)
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
			member.GET("/", middle.AccountTypeMiddleware("Admin", "Owner", "Employee", "Member"), memberController.Index)
			member.GET("/:id", middle.AccountTypeMiddleware("Admin", "Owner", "Employee", "Member"), memberController.Show)
			member.POST("/", middle.AccountTypeMiddleware("Admin", "Owner", "Employee"), memberController.Store)
			member.PUT("/:id", middle.AccountTypeMiddleware("Admin", "Owner", "Employee", "Member"), memberController.Update)
			member.DELETE("/:id", middle.AccountTypeMiddleware("Admin", "Owner", "Employee"), memberController.Destroy)

			member.POST("/forgot-password-reset-link", middle.AccountTypeMiddleware("Admin", "Owner", "Employee", "Member"), memberController.ForgotPassword)
			member.POST("/change-password", middle.AccountTypeMiddleware("Member"), memberController.ChangePassword)
			member.POST("/forgot-password", middle.AccountTypeMiddleware("Member"), memberController.ForgotPassword)
			member.POST("/new-password", middle.AccountTypeMiddleware("Member"), memberController.NewPassword)

			member.POST("/skip-verification", middle.AccountTypeMiddleware("Member"), memberController.SkipVerification)
			member.POST("/send-email-verification", middle.AccountTypeMiddleware("Member"), memberController.SendEmailVerification)
			member.POST("/verify-email", middle.AccountTypeMiddleware("Member"), memberController.VerifyEmail)

			member.POST("/send-contact-number-verification", middle.AccountTypeMiddleware("Member"), memberController.SendContactNumberVerification)
			member.POST("/verify-contact-number", middle.AccountTypeMiddleware("Member"), memberController.VerifyContactNumber)

			member.POST("/profile-picture", middle.AccountTypeMiddleware("Member"), memberController.ProfilePicture)
			member.PUT("/profile-account-setting", middle.AccountTypeMiddleware("Member"), memberController.ProfileAccountSetting)
			member.PUT("/profile-change-email", middle.AccountTypeMiddleware("Member"), memberController.ProfileChangeEmail)
			member.PUT("/profile-change-contact-number", middle.AccountTypeMiddleware("Member"), memberController.ProfileChangeContactNumber)
			member.PUT("/profile-change-username", middle.AccountTypeMiddleware("Member"), memberController.ProfileChangeUsername)
		}

		memberProfile := v1.Group("/member-profile")
		{
			memberProfile.GET("/", memberProfileController.Index)
			memberProfile.GET("/:id", memberProfileController.Show)
			memberProfile.POST("/", memberProfileController.Store)
			memberProfile.PUT("/:id", memberProfileController.Update)
			memberProfile.DELETE("/:id", memberProfileController.Destroy)
		}

		owner := v1.Group("/owner")
		{
			owner.GET("/", middle.AccountTypeMiddleware("Admin"), ownerController.Index)
			owner.GET("/:id", middle.AccountTypeMiddleware("Admin", "Owner", "Employee"), ownerController.Show)
			owner.POST("/", middle.AccountTypeMiddleware("Admin"), ownerController.Store)
			owner.PUT("/:id", middle.AccountTypeMiddleware("Admin", "Owner"), ownerController.Update)
			owner.DELETE("/:id", middle.AccountTypeMiddleware("Amin"), ownerController.Destroy)

			owner.POST("/change-password", middle.AccountTypeMiddleware("Owner"), ownerController.ChangePassword)
			owner.POST("/forgot-password", middle.AccountTypeMiddleware("Owner"), ownerController.ForgotPassword)
			owner.POST("/forgot-password-reset-link", middle.AccountTypeMiddleware("Admin", "Owner"), ownerController.ForgotPassword)
			owner.POST("/new-password", middle.AccountTypeMiddleware("Owner"), ownerController.NewPassword)

			owner.POST("/skip-verification", middle.AccountTypeMiddleware("Owner"), ownerController.SkipVerification)
			owner.POST("/send-email-verification", middle.AccountTypeMiddleware("Owner"), ownerController.SendEmailVerification)
			owner.POST("/verify-email", middle.AccountTypeMiddleware("Owner"), ownerController.VerifyEmail)

			owner.POST("/send-contact-number-verification", middle.AccountTypeMiddleware("Owner"), ownerController.SendContactNumberVerification)
			owner.POST("/verify-contact-number", middle.AccountTypeMiddleware("Owner"), ownerController.VerifyContactNumber)

			owner.POST("/profile-picture", middle.AccountTypeMiddleware("Owner"), ownerController.ProfilePicture)
			owner.PUT("/profile-account-setting", middle.AccountTypeMiddleware("Owner"), ownerController.ProfileAccountSetting)
			owner.PUT("/profile-change-email", middle.AccountTypeMiddleware("Owner"), ownerController.ProfileChangeEmail)
			owner.PUT("/profile-change-contact-number", middle.AccountTypeMiddleware("Owner"), ownerController.ProfileChangeContactNumber)
			owner.PUT("/profile-change-username", middle.AccountTypeMiddleware("Owner"), ownerController.ProfileChangeUsername)
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

		timesheet := v1.Group("/timesheet")
		{
			timesheet.GET("/", timesheetController.Index)
			timesheet.GET("/current", timesheetController.Current)
			timesheet.POST("/time-in", timesheetController.TimeIn)
			timesheet.POST("/time-out", timesheetController.TimeOut)
		}

		qr := v1.Group("/qr")
		{
			qr.GET("/profile", qrController.Profile)
			qr.GET("/find-profile", qrController.FindProfile)
		}
	}

	runServer(lc, cfg, engineService, logger)
}
