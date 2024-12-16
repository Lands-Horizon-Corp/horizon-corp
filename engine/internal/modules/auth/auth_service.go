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
	otpProvider   *providers.OTPService
	authProvider  *AuthProvider
	authAccount   *auth_accounts.AuthAccount
	modelResource *models.ModelResource
}

func NewAuthService(
	engine *providers.EngineService,
	middle *middleware.Middleware,
	tokenProvider *providers.TokenService,
	otpProvider *providers.OTPService,
	authProvider *AuthProvider,
	authAccount *auth_accounts.AuthAccount,
	modelResource *models.ModelResource,
) *AuthService {
	return &AuthService{
		engine:        engine,
		middle:        middle,
		otpProvider:   otpProvider,
		tokenProvider: tokenProvider,
		authProvider:  authProvider,
		authAccount:   authAccount,
		modelResource: modelResource,
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
		as.tokenProvider.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	user, err := as.authAccount.GetByID(userClaims.AccountType, userClaims.ID)
	if err != nil {
		as.tokenProvider.ClearTokenCookie(ctx)
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
	var req VerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("VerifyEmail: JSON binding error: %v", err)})
		return
	}
	if err := as.authProvider.ValidateVerifyEmailRequest(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("VerifyEmail: Validation error: %v", err)})
		return
	}
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	isValid, err := as.otpProvider.ValidateOTP(claims.AccountType, claims.ID, req.Otp, "email")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("VerifyEmail: OTP validation error: %v", err)})
		return
	}
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "VerifyEmail: Invalid or expired OTP"})
		return
	}

	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberVerifyEmail(ctx, claims.ID)
	case "Admin":
		as.authAccount.AdminVerifyEmail(ctx, claims.ID)
	case "Owner":
		as.authAccount.OwnerVerifyEmail(ctx, claims.ID)
	case "Employee":
		as.authAccount.EmployeeVerifyEmail(ctx, claims.ID)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (as AuthService) SendContactNumberVerification(ctx *gin.Context) {
	var req SendContactNumberVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err)})
		return
	}
	if err := as.authProvider.ValidateSendContactNumberVerificationRequest(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: Validation error: %v", err)})
		return
	}
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberSendContactNumberVerification(ctx, claims.ID, req.ContactTemplate)
	case "Admin":
		as.authAccount.AdminSendContactNumberVerification(ctx, claims.ID, req.ContactTemplate)
	case "Owner":
		as.authAccount.OwnerSendContactNumberVerification(ctx, claims.ID, req.ContactTemplate)
	case "Employee":
		as.authAccount.EmployeeSendContactNumberVerification(ctx, claims.ID, req.ContactTemplate)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (as AuthService) VerifyContactNumber(ctx *gin.Context) {

	var req VerifyContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err)})
		return
	}
	if err := as.authProvider.ValidateVerifyContactNumberRequest(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: Validation error: %v", err)})
		return
	}
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	isValid, err := as.otpProvider.ValidateOTP(claims.AccountType, claims.ID, req.Otp, "sms")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("VerifyContactNumber: OTP validation error: %v", err)})
		return
	}
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "VerifyContactNumber: Invalid or expired OTP"})
		return
	}

	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberVerifyContactNumber(ctx, claims.ID)
	case "Admin":
		as.authAccount.AdminVerifyContactNumber(ctx, claims.ID)
	case "Owner":
		as.authAccount.OwnerVerifyContactNumber(ctx, claims.ID)
	case "Employee":
		as.authAccount.EmployeeVerifyContactNumber(ctx, claims.ID)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}

func (as AuthService) ProfilePicture(ctx *gin.Context) {
	var req *models.MediaRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err)})
		return
	}
	if err := as.modelResource.ValidateMediaRequest(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: Validation error: %v", err)})
		return
	}
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberProfilePicture(ctx, claims.ID, *req.ID)
	case "Admin":
		as.authAccount.AdminProfilePicture(ctx, claims.ID, *req.ID)
	case "Owner":
		as.authAccount.OwnerProfilePicture(ctx, claims.ID, *req.ID)
	case "Employee":
		as.authAccount.EmployeeProfilePicture(ctx, claims.ID, *req.ID)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}

}
func (as AuthService) ProfileAccountSetting(ctx *gin.Context) {
	var req *AccountSettingRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err)})
		return
	}
	if err := as.authProvider.ValidateAccountSettingRequest(*req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: Validation error: %v", err)})
		return
	}
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberProfileAccountSetting(ctx, claims.ID, req.BirthDate, req.FirstName, req.MiddleName, req.LastName, req.Description, req.PermanentAddress)
	case "Admin":
		as.authAccount.AdminProfileAccountSetting(ctx, claims.ID, req.BirthDate, req.FirstName, req.MiddleName, req.LastName, req.Description, req.PermanentAddress)
	case "Owner":
		as.authAccount.OwnerProfileAccountSetting(ctx, claims.ID, req.BirthDate, req.FirstName, req.MiddleName, req.LastName, req.Description, req.PermanentAddress)
	case "Employee":
		as.authAccount.EmployeeProfileAccountSetting(ctx, claims.ID, req.BirthDate, req.FirstName, req.MiddleName, req.LastName, req.Description, req.PermanentAddress)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}
func (as AuthService) ProfileChangeEmail(ctx *gin.Context) {
	var req *ChangeEmailRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err)})
		return
	}
	if err := as.authProvider.ValidateChangeEmailRequest(*req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: Validation error: %v", err)})
		return
	}
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberProfileChangeEmail(ctx, claims.ID, req.Password, req.Email)
	case "Admin":
		as.authAccount.AdminProfileChangeEmail(ctx, claims.ID, req.Password, req.Email)
	case "Owner":
		as.authAccount.OwnerProfileChangeEmail(ctx, claims.ID, req.Password, req.Email)
	case "Employee":
		as.authAccount.EmployeeProfileChangeEmail(ctx, claims.ID, req.Password, req.Email)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}
func (as AuthService) ProfileChangeContactNumber(ctx *gin.Context) {
	var req *ChangeContactNumberRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err)})
		return
	}
	if err := as.authProvider.ValidateChangeContactNumberRequest(*req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: Validation error: %v", err)})
		return
	}
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberProfileChangeContactNumber(ctx, claims.ID, req.Password, req.ContactNumber)
	case "Admin":
		as.authAccount.AdminProfileChangeContactNumber(ctx, claims.ID, req.Password, req.ContactNumber)
	case "Owner":
		as.authAccount.OwnerProfileChangeContactNumber(ctx, claims.ID, req.Password, req.ContactNumber)
	case "Employee":
		as.authAccount.EmployeeProfileChangeContactNumber(ctx, claims.ID, req.Password, req.ContactNumber)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}
func (as AuthService) ProfileChangeUsername(ctx *gin.Context) {
	var req *ChangeUsernameRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err)})
		return
	}
	if err := as.authProvider.ValidateChangeUsernameRequest(*req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: Validation error: %v", err)})
		return
	}
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberProfileChangeUsername(ctx, claims.ID, req.Password, req.Username)
	case "Admin":
		as.authAccount.AdminProfileChangeUsername(ctx, claims.ID, req.Password, req.Username)
	case "Owner":
		as.authAccount.OwnerProfileChangeUsername(ctx, claims.ID, req.Password, req.Username)
	case "Employee":
		as.authAccount.EmployeeProfileChangeUsername(ctx, claims.ID, req.Password, req.Username)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
}
func (as AuthService) ProfileChangePassword(ctx *gin.Context) {
	var req *ChangePasswordSettingRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err)})
		return
	}
	if err := as.authProvider.ValidateChangePasswordSetting(*req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("SendContactNumberVerification: Validation error: %v", err)})
		return
	}
	claims, err := as.getUserClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	switch claims.AccountType {
	case "Member":
		as.authAccount.MemberProfileChangePassword(ctx, claims.ID, req.OldPassword, req.NewPassword)
	case "Admin":
		as.authAccount.AdminProfileChangePassword(ctx, claims.ID, req.OldPassword, req.NewPassword)
	case "Owner":
		as.authAccount.OwnerProfileChangePassword(ctx, claims.ID, req.OldPassword, req.NewPassword)
	case "Employee":
		as.authAccount.EmployeeProfileChangePassword(ctx, claims.ID, req.OldPassword, req.NewPassword)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Account type doesn't exist"})
	}
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

	profileAuth := as.engine.Client.Group("/profile")
	profileAuth.Use(as.middle.AuthMiddleware())
	{
		profileAuth.POST("/profile-picture", as.ProfilePicture)
		profileAuth.POST("/account-setting", as.ProfileAccountSetting)
		profileAuth.POST("/change-email", as.ProfileChangeEmail)
		profileAuth.POST("/change-contact-number", as.ProfileChangeContactNumber)
		profileAuth.POST("/change-username", as.ProfileChangeUsername)
		profileAuth.POST("/change-password", as.ChangePassword)
	}
}
