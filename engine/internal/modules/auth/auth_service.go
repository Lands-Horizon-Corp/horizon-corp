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
		as.authAccount.MemberSignUp(ctx)
	case "Admin":
		as.authAccount.AdminSignUp(ctx)
	case "Owner":
		as.authAccount.OwnerSignUp(ctx)
	case "Employee":
		as.authAccount.EmployeeSignUp(ctx)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
	as.authAccount.AdminSignUp(ctx)

	// switch req.AccountType {
	// case "Member":
	// 	user := &models.Member{
	// 		FirstName:         req.FirstName,
	// 		LastName:          req.LastName,
	// 		MiddleName:        req.MiddleName,
	// 		PermanentAddress:  req.PermanentAddress,
	// 		Description:       "",
	// 		BirthDate:         req.BirthDate,
	// 		Username:          req.Username,
	// 		Email:             req.Email,
	// 		Password:          req.Password,
	// 		IsEmailVerified:   false,
	// 		IsContactVerified: false,
	// 		ContactNumber:     req.ContactNumber,
	// 		MediaID:           nil,
	// 		Status:            "Pending",
	// 	}
	// 	id, err := as.authProvider.Create(user, req.AccountType)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 		return
	// 	}
	// 	token, err := as.authToken.GenerateUserToken(id, req.AccountType)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 		return
	// 	}

	// 	// Send Email OTP
	// 	emailReq := providers.EmailRequest{
	// 		To:      req.Email,
	// 		Subject: "ECOOP: Email Verification",
	// 		Body:    req.EmailTemplate,
	// 	}
	// 	if err := as.otpProvider.SendEmailOTP(req.AccountType, user.ID, emailReq); err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 		return
	// 	}

	// 	// Send SMS OTP
	// 	contactReq := providers.SMSRequest{
	// 		To:   req.ContactNumber,
	// 		Body: req.ContactTemplate,
	// 		Vars: &map[string]string{
	// 			"name": fmt.Sprintf("%s %s", req.FirstName, req.LastName),
	// 		},
	// 	}
	// 	if err := as.otpProvider.SendContactNumberOTP(req.AccountType, user.ID, contactReq); err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 		return
	// 	}

	// 	http.SetCookie(ctx.Writer, &http.Cookie{
	// 		Name:     as.cfg.AppTokenName,
	// 		Value:    *token,
	// 		Path:     "/",
	// 		HttpOnly: true,
	// 		Secure:   true,
	// 		SameSite: http.SameSiteNoneMode,
	// 	})
	// 	ctx.JSON(http.StatusCreated, user)
	// 	return

	// case "Employee":
	// 	user := &models.Employee{
	// 		FirstName:         req.FirstName,
	// 		LastName:          req.LastName,
	// 		MiddleName:        req.MiddleName,
	// 		PermanentAddress:  req.PermanentAddress,
	// 		Description:       "",
	// 		BirthDate:         req.BirthDate,
	// 		Username:          req.Username,
	// 		Email:             req.Email,
	// 		Password:          req.Password,
	// 		IsEmailVerified:   false,
	// 		IsContactVerified: false,
	// 		ContactNumber:     req.ContactNumber,
	// 		MediaID:           nil,
	// 		Status:            "Pending",
	// 	}
	// 	created, err := as.authProvider.Create(user, req.AccountType)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 		return
	// 	}
	// case "Admin":
	// 	user := &models.Admin{
	// 		FirstName:         req.FirstName,
	// 		LastName:          req.LastName,
	// 		MiddleName:        req.MiddleName,
	// 		PermanentAddress:  req.PermanentAddress,
	// 		Description:       "",
	// 		BirthDate:         req.BirthDate,
	// 		Username:          req.Username,
	// 		Email:             req.Email,
	// 		Password:          req.Password,
	// 		IsEmailVerified:   false,
	// 		IsContactVerified: false,
	// 		ContactNumber:     req.ContactNumber,
	// 		MediaID:           nil,
	// 		Status:            "Pending",
	// 	}
	// 	created, err := as.authProvider.Create(user, req.AccountType)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 		return
	// 	}
	// case "Owner":
	// 	user := &models.Owner{
	// 		FirstName:         req.FirstName,
	// 		LastName:          req.LastName,
	// 		MiddleName:        req.MiddleName,
	// 		PermanentAddress:  req.PermanentAddress,
	// 		Description:       "",
	// 		BirthDate:         req.BirthDate,
	// 		Username:          req.Username,
	// 		Email:             req.Email,
	// 		Password:          req.Password,
	// 		IsEmailVerified:   false,
	// 		IsContactVerified: false,
	// 		ContactNumber:     req.ContactNumber,
	// 		MediaID:           nil,
	// 		Status:            "Pending",
	// 	}
	// 	created, err := as.authProvider.Create(user, req.AccountType)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	// 		return
	// 	}
	// }
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
