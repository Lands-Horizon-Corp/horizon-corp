package auth

import (
	"fmt"
	"net/http"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/database/models"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/modules/auth/auth_accounts"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/providers"
	"github.com/Lands-Horizon-Corp/horizon-corp/server/middleware"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	engine        *providers.EngineService
	middle        *middleware.Middleware
	tokenProvider *providers.TokenService
	authProvider  *AuthProvider
	authAccount   *auth_accounts.AuthAccount
}

func NewAuthService(
	engine *providers.EngineService,
	middle *middleware.Middleware,
	tokenProvider *providers.TokenService,
	authProvider *AuthProvider,
	authAccount *auth_accounts.AuthAccount,
) *AuthService {
	return &AuthService{
		engine:        engine,
		middle:        middle,
		tokenProvider: tokenProvider,
		authAccount:   authAccount,
	}
}

func (as *AuthService) getUserClaims(ctx *gin.Context) (*providers.UserClaims, error) {
	claims, exists := ctx.Get("claims")
	if !exists {
		as.tokenProvider.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("claims not found in context")
	}

	userClaims, ok := claims.(*providers.UserClaims)
	if !ok {
		as.tokenProvider.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("failed to cast claims to *auth.UserClaims")
	}

	return userClaims, nil
}

func (as AuthService) SignUp(ctx *gin.Context) {
	var req SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := as.authProvider.ValidateSignUp(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	switch req.AccountType {
	case "Member":
		user := &models.Member{
			FirstName:         req.FirstName,
			LastName:          req.LastName,
			MiddleName:        req.MiddleName,
			PermanentAddress:  req.PermanentAddress,
			Description:       "",
			BirthDate:         req.BirthDate,
			Username:          req.Username,
			Email:             req.Email,
			Password:          req.Password,
			IsEmailVerified:   false,
			IsContactVerified: false,
			ContactNumber:     req.ContactNumber,
			MediaID:           nil,
			Status:            "Pending",
		}
		as.authAccount.MemberSignUp(ctx, user, req.EmailTemplate, req.ContactTemplate)
	case "Admin":
		user := &models.Admin{
			FirstName:         req.FirstName,
			LastName:          req.LastName,
			MiddleName:        req.MiddleName,
			PermanentAddress:  req.PermanentAddress,
			Description:       "",
			BirthDate:         req.BirthDate,
			Username:          req.Username,
			Email:             req.Email,
			Password:          req.Password,
			IsEmailVerified:   false,
			IsContactVerified: false,
			ContactNumber:     req.ContactNumber,
			MediaID:           nil,
			Status:            "Pending",
		}
		as.authAccount.AdminSignUp(ctx, user, req.EmailTemplate, req.ContactTemplate)
	case "Owner":
		user := &models.Owner{
			FirstName:         req.FirstName,
			LastName:          req.LastName,
			MiddleName:        req.MiddleName,
			PermanentAddress:  req.PermanentAddress,
			Description:       "",
			BirthDate:         req.BirthDate,
			Username:          req.Username,
			Email:             req.Email,
			Password:          req.Password,
			IsEmailVerified:   false,
			IsContactVerified: false,
			ContactNumber:     req.ContactNumber,
			MediaID:           nil,
			Status:            "Pending",
		}
		as.authAccount.OwnerSignUp(ctx, user, req.EmailTemplate, req.ContactTemplate)
	case "Employee":
		user := &models.Employee{
			FirstName:         req.FirstName,
			LastName:          req.LastName,
			MiddleName:        req.MiddleName,
			PermanentAddress:  req.PermanentAddress,
			Description:       "",
			BirthDate:         req.BirthDate,
			Username:          req.Username,
			Email:             req.Email,
			Password:          req.Password,
			IsEmailVerified:   false,
			IsContactVerified: false,
			ContactNumber:     req.ContactNumber,
			MediaID:           nil,
			Status:            "Pending",
		}
		as.authAccount.EmployeeSignUp(ctx, user, req.EmailTemplate, req.ContactTemplate)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (as AuthService) SignIn(ctx *gin.Context) {
	var req SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := as.authProvider.ValidateSignIn(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	switch req.AccountType {
	case "Member":
		as.authAccount.MemberSignIn(ctx, req.Key, req.Password)
	case "Admin":
		as.authAccount.AdminSignIn(ctx, req.Key, req.Password)
	case "Owner":
		as.authAccount.OwnerSignIn(ctx, req.Key, req.Password)
	case "Employee":
		as.authAccount.EmployeeSignIn(ctx, req.Key, req.Password)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}

}

func (as AuthService) ForgotPassword(ctx *gin.Context) {
	var req ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := as.authProvider.ValidateForgotpassword(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	switch req.AccountType {
	case "Member":
		as.authAccount.MemberForgotPassword(ctx, req.Key, req.EmailTemplate, req.ContactTemplate)
	case "Admin":
		as.authAccount.AdminForgotPassword(ctx, req.Key, req.EmailTemplate, req.ContactTemplate)
	case "Owner":
		as.authAccount.OwnerForgotPassword(ctx, req.Key, req.EmailTemplate, req.ContactTemplate)
	case "Employee":
		as.authAccount.EmployeeForgotPassword(ctx, req.Key, req.EmailTemplate, req.ContactTemplate)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (as AuthService) ChangePassword(ctx *gin.Context) {
	var req ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ChangePassword: JSON binding error: %v", err)})
		return
	}

	if err := as.authProvider.ValidateChangePassword(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ChangePassword: Validation error: %v", err)})
		return
	}
	claims, err := as.tokenProvider.VerifyToken(req.ResetID)
	if err != nil {
		as.tokenProvider.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ChangePassword: Token verification error: %v", err)})
		return
	}
	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberChangePassword(ctx, claims.ID, req.NewPassword)
	case "Admin":
		as.authAccount.AdminChangePassword(ctx, claims.ID, req.NewPassword)
	case "Owner":
		as.authAccount.OwnerChangePassword(ctx, claims.ID, req.NewPassword)
	case "Employee":
		as.authAccount.EmployeeChangePassword(ctx, claims.ID, req.NewPassword)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
	as.tokenProvider.DeleteToken(req.ResetID)
}

func (as AuthService) VerifyResetLink(ctx *gin.Context) {
	_, err := as.tokenProvider.VerifyToken(ctx.Param("id"))
	if err != nil {
		as.tokenProvider.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("VerifyResetLink: Token verification error: %v", err)})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Link verified successfully."})
}

func (as AuthService) SignOut(ctx *gin.Context) {
	as.tokenProvider.ClearTokenCookie(ctx)
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully signed out"})
}

func (as AuthService) CurrentUser(ctx *gin.Context) {
	userClaims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	user, err := as.authAccount.GetByID(userClaims.AccountType, userClaims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found."})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (as AuthService) NewPassword(ctx *gin.Context) {
	var req NewPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ChangePassword: JSON binding error: %v", err)})
		return
	}
	if err := as.authProvider.ValidateNewPassword(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ChangePassword: Validation error: %v", err)})
		return
	}
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberChangePassword(ctx, claims.ID, req.NewPassword)
	case "Admin":
		as.authAccount.AdminChangePassword(ctx, claims.ID, req.NewPassword)
	case "Owner":
		as.authAccount.OwnerChangePassword(ctx, claims.ID, req.NewPassword)
	case "Employee":
		as.authAccount.EmployeeChangePassword(ctx, claims.ID, req.NewPassword)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}

}

func (as AuthService) SkipVerification(ctx *gin.Context) {
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberSkipVerification(ctx, claims.ID)
	case "Admin":
		as.authAccount.AdminSkipVerification(ctx, claims.ID)
	case "Owner":
		as.authAccount.OwnerSkipVerification(ctx, claims.ID)
	case "Employee":
		as.authAccount.EmployeeSkipVerification(ctx, claims.ID)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (as AuthService) SendEmailVerification(ctx *gin.Context) {
	var req SendEmailVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ChangePassword: JSON binding error: %v", err)})
		return
	}
	if err := as.authProvider.ValidateSendEmailVerification(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("ChangePassword: Validation error: %v", err)})
		return
	}

	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberSendEmailVerification(ctx, claims.ID, req.EmailTemplate)
	case "Admin":
		as.authAccount.AdminSendEmailVerification(ctx, claims.ID, req.EmailTemplate)
	case "Owner":
		as.authAccount.OwnerSendEmailVerification(ctx, claims.ID, req.EmailTemplate)
	case "Employee":
		as.authAccount.EmployeeSendEmailVerification(ctx, claims.ID, req.EmailTemplate)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}

}

func (as AuthService) VerifyEmail(ctx *gin.Context) {

}

func (as AuthService) SendContactNumberVerification(ctx *gin.Context) {

}

func (as AuthService) VerifyContactNumber(ctx *gin.Context) {

}

func (as *AuthService) RegisterRoutes() {
	routes := as.engine.Client.Group("/api/v1/auth")
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
