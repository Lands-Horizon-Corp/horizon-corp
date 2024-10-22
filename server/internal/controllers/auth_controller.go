package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"horizon/server/config"
	"horizon/server/helpers"
	"horizon/server/internal/auth"
	"horizon/server/internal/middleware"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests/auth_requests"
	"horizon/server/internal/resources"
	"horizon/server/services"

	"github.com/gin-gonic/gin"
)

// AuthController handles authentication-related operations.
type AuthController struct {
	userRepo        *repositories.UserRepository
	userAuthService *auth.UserAuthService
	cfg             *config.AppConfig
	emailService    *services.EmailService
	contactService  *services.SMSService
	otpService      *services.OTPService
	tokenService    auth.TokenService
}

// NewAuthController creates a new instance of AuthController.
func NewAuthController(
	userRepo *repositories.UserRepository,
	userAuthService *auth.UserAuthService,
	cfg *config.AppConfig,
	emailService *services.EmailService,
	contactService *services.SMSService,
	otpService *services.OTPService,
	tokenService auth.TokenService,
) *AuthController {
	return &AuthController{
		userRepo:        userRepo,
		userAuthService: userAuthService,
		cfg:             cfg,
		emailService:    emailService,
		contactService:  contactService,
		otpService:      otpService,
		tokenService:    tokenService,
	}
}

// respondWithError logs the error and sends a generic error message to the client.
func (c *AuthController) respondWithError(ctx *gin.Context, statusCode int, logMessage string, clientMessage string) {
	log.Println(logMessage)
	ctx.JSON(statusCode, gin.H{"error": clientMessage})
}

// getCurrentUser retrieves the current authenticated user from the context.
func (c *AuthController) getCurrentUser(ctx *gin.Context) (models.User, error) {
	user, exists := ctx.Get("current-user")
	if !exists {
		c.tokenService.ClearTokenCookie(ctx)
		return models.User{}, fmt.Errorf("current user not found in context")
	}

	currentUser, ok := user.(models.User)
	if !ok {
		c.tokenService.ClearTokenCookie(ctx)
		return models.User{}, fmt.Errorf("failed to cast user to models.User")
	}

	return currentUser, nil
}

// getUserClaims retrieves the user claims from the context.
func (c *AuthController) getUserClaims(ctx *gin.Context) (*auth.UserClaims, error) {
	claims, exists := ctx.Get("claims")
	if !exists {
		c.tokenService.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("claims not found in context")
	}

	userClaims, ok := claims.(*auth.UserClaims)
	if !ok {
		c.tokenService.ClearTokenCookie(ctx)
		return nil, fmt.Errorf("failed to cast claims to *auth.UserClaims")
	}

	return userClaims, nil
}

// setAuthTokenCookie sets the authentication token as an HTTP-only cookie.
func (c *AuthController) setAuthTokenCookie(ctx *gin.Context, token string) {
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     c.cfg.AppTokenName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})
}

// CurrentUser retrieves the current authenticated user.
func (c *AuthController) CurrentUser(ctx *gin.Context) {
	currentUser, err := c.getCurrentUser(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "Authentication required.")
		return
	}

	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "User not authenticated.")
		return
	}

	currentUser.AccountType = userClaims.AccountType
	response := resources.ToResourceUser(currentUser, currentUser.AccountType)
	ctx.JSON(http.StatusOK, response)
}

// SignUp handles user registration.
func (c *AuthController) SignUp(ctx *gin.Context) {
	var req auth_requests.SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("SignUp: JSON binding error: %v", err), "Invalid request payload.")
		return
	}

	if err := req.Validate(); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("SignUp: Validation error: %v", err), "Invalid input data.")
		return
	}

	userReq := models.User{
		FirstName:         req.FirstName,
		LastName:          req.LastName,
		MiddleName:        req.MiddleName,
		PermanentAddress:  req.PermanentAddress,
		Description:       "",
		Birthdate:         req.Birthdate,
		Username:          req.Username,
		Email:             req.Email,
		Password:          req.Password,
		IsEmailVerified:   false,
		IsContactVerified: false,
		ContactNumber:     req.ContactNumber,
		MediaID:           nil,
		Status:            "Pending",
	}

	user, err := c.userRepo.Create(req.AccountType, userReq)
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("SignUp: User repository create error: %v", err), "Failed to create user account.")
		return
	}

	token, err := c.userAuthService.GenerateUserToken(user, req.AccountType, 0)
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("SignUp: Token generation error: %v", err), "Failed to process registration.")
		return
	}

	// Send Email OTP
	emailReq := services.EmailRequest{
		To:      req.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}
	if err := c.otpService.SendEmailOTP(req.AccountType, user.ID, emailReq); err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("SignUp: Email OTP sending error: %v", err), "Failed to send verification email.")
		return
	}

	// Send SMS OTP
	contactReq := services.SMSRequest{
		To:   req.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": fmt.Sprintf("%s %s", req.FirstName, req.LastName),
		},
	}
	if err := c.contactService.SendSMS(contactReq); err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("SignUp: SMS sending error: %v", err), "Failed to send verification SMS.")
		return
	}

	// Set Auth Token Cookie
	c.setAuthTokenCookie(ctx, token)

	// Respond with the created user
	response := resources.ToResourceUser(user, user.AccountType)
	ctx.JSON(http.StatusCreated, response)
}

// SignIn handles user authentication.
func (c *AuthController) SignIn(ctx *gin.Context) {
	var req auth_requests.SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("SignIn: JSON binding error: %v", err), "Invalid request payload.")
		return
	}

	if err := req.Validate(); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("SignIn: Validation error: %v", err), "Invalid input data.")
		return
	}

	user, err := c.userRepo.FindByEmailUsernameOrContact(req.AccountType, req.Key)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, fmt.Sprintf("SignIn: User not found: %v", err), "Invalid credentials.")
		return
	}

	if !config.VerifyPassword(user.Password, req.Password) {
		c.respondWithError(ctx, http.StatusUnauthorized, "SignIn: Password verification failed", "Invalid credentials.")
		return
	}

	token, err := c.userAuthService.GenerateUserToken(user, req.AccountType, 0)
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("SignIn: Token generation error: %v", err), "Failed to authenticate user.")
		return
	}

	// Set Auth Token Cookie
	c.setAuthTokenCookie(ctx, token)

	// Respond with the authenticated user
	response := resources.ToResourceUser(user, req.AccountType)
	ctx.JSON(http.StatusOK, response)
}

// SignOut handles user sign-out by clearing the auth token.
func (c *AuthController) SignOut(ctx *gin.Context) {
	c.tokenService.ClearTokenCookie(ctx)
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully signed out"})
}

// ForgotPassword initiates the password reset process.
func (c *AuthController) ForgotPassword(ctx *gin.Context) {
	var req auth_requests.ForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("ForgotPassword: JSON binding error: %v", err), "Invalid request payload.")
		return
	}

	if err := req.Validate(); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("ForgotPassword: Validation error: %v", err), "Invalid input data.")
		return
	}

	user, err := c.userRepo.FindByEmailUsernameOrContact(req.AccountType, req.Key)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, fmt.Sprintf("ForgotPassword: User not found: %v", err), "User not found.")
		return
	}

	token, err := c.userAuthService.GenerateUserToken(user, req.AccountType, time.Minute*10)
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("ForgotPassword: Token generation error: %v", err), "Failed to process password reset.")
		return
	}

	resetLink := fmt.Sprintf("%s/auth/password-reset/%s", c.cfg.AppClientUrl, token)
	keyType := helpers.GetKeyType(req.Key)

	switch keyType {
	case "contact":
		contactReq := services.SMSRequest{
			To:   user.ContactNumber,
			Body: req.ContactTemplate,
			Vars: &map[string]string{
				"name":      fmt.Sprintf("%s %s", user.FirstName, user.LastName),
				"eventLink": resetLink,
			},
		}
		if err := c.contactService.SendSMS(contactReq); err != nil {
			c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("ForgotPassword: SMS sending error: %v", err), "Failed to send password reset SMS.")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Password reset instructions have been sent via SMS."})

	case "email":
		emailReq := services.EmailRequest{
			To:      user.Email,
			Subject: "ECOOP: Change Password Request",
			Body:    req.EmailTemplate,
			Vars: &map[string]string{
				"name":      fmt.Sprintf("%s %s", user.FirstName, user.LastName),
				"eventLink": resetLink,
			},
		}
		if err := c.emailService.SendEmail(emailReq); err != nil {
			c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("ForgotPassword: Email sending error: %v", err), "Failed to send password reset email.")
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Password reset instructions have been sent via email."})

	default:
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("ForgotPassword: Invalid key type: %s", keyType), "Invalid key type provided.")
	}
}

// ChangePassword updates the user's password.
func (c *AuthController) ChangePassword(ctx *gin.Context) {
	var req auth_requests.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("ChangePassword: JSON binding error: %v", err), "Invalid request payload.")
		return
	}

	if err := req.Validate(); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("ChangePassword: Validation error: %v", err), "Invalid input data.")
		return
	}

	claims, err := c.tokenService.VerifyToken(req.ResetID)
	if err != nil {
		c.tokenService.ClearTokenCookie(ctx)
		c.respondWithError(ctx, http.StatusUnauthorized, fmt.Sprintf("ChangePassword: Token verification error: %v", err), "Invalid or expired reset token.")
		return
	}

	if err := c.userRepo.UpdatePassword(claims.AccountType, claims.ID, req.ConfirmPassword); err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("ChangePassword: Password update error: %v", err), "Failed to update password.")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully."})
}

func (c *AuthController) VerifyResetLink(ctx *gin.Context) {
	_, err := c.tokenService.VerifyToken(ctx.Param("id"))

	if err != nil {
		c.tokenService.ClearTokenCookie(ctx)
		c.respondWithError(ctx, http.StatusUnauthorized, fmt.Sprintf("VerifyResetLink: Token verification error: %v", err), "Invalid or expired reset token.")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Link verified successfully."})
}

// SendEmailVerification sends an email verification OTP to the user.
func (c *AuthController) SendEmailVerification(ctx *gin.Context) {
	var req auth_requests.SendEmailVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("SendEmailVerification: JSON binding error: %v", err), "Invalid request payload.")
		return
	}

	if err := req.Validate(); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("SendEmailVerification: Validation error: %v", err), "Invalid input data.")
		return
	}

	currentUser, err := c.getCurrentUser(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "User not authenticated.")
		return
	}

	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "User not authenticated.")
		return
	}

	emailReq := services.EmailRequest{
		To:      currentUser.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}

	if err := c.otpService.SendEmailOTP(userClaims.AccountType, currentUser.ID, emailReq); err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("SendEmailVerification: Email OTP sending error: %v", err), "Failed to send email verification.")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email verification sent successfully. Please check your inbox or spam folder."})
}

// VerifyEmail verifies the user's email using the provided OTP.
func (c *AuthController) VerifyEmail(ctx *gin.Context) {
	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "User not authenticated.")
		return
	}

	var req auth_requests.VerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("VerifyEmail: JSON binding error: %v", err), "Invalid request payload.")
		return
	}

	if err := req.Validate(); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("VerifyEmail: Validation error: %v", err), "Invalid input data.")
		return
	}

	isValid, err := c.otpService.ValidateOTP(userClaims.AccountType, userClaims.ID, req.Otp, "email")
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("VerifyEmail: OTP validation error: %v", err), "Internal server error.")
		return
	}
	if !isValid {
		c.respondWithError(ctx, http.StatusUnauthorized, "VerifyEmail: Invalid or expired OTP", "Invalid or expired OTP.")
		return
	}

	updatedUser, err := c.userRepo.UpdateColumns(userClaims.AccountType, userClaims.ID, map[string]interface{}{
		"is_email_verified": true,
	})
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("VerifyEmail: User update error: %v", err), "Failed to update user status.")
		return
	}

	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

// SendContactNumberVerification sends an SMS verification OTP to the user's contact number.
func (c *AuthController) SendContactNumberVerification(ctx *gin.Context) {
	var req auth_requests.SendContactNumberVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("SendContactNumberVerification: JSON binding error: %v", err), "Invalid request payload.")
		return
	}

	if err := req.Validate(); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("SendContactNumberVerification: Validation error: %v", err), "Invalid input data.")
		return
	}

	currentUser, err := c.getCurrentUser(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "User not authenticated.")
		return
	}

	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "User not authenticated.")
		return
	}

	contactReq := services.SMSRequest{
		To:   currentUser.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": fmt.Sprintf("%s %s", currentUser.FirstName, currentUser.LastName),
		},
	}

	if err := c.otpService.SendEContactNumberOTP(userClaims.AccountType, userClaims.ID, contactReq); err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("SendContactNumberVerification: SMS sending error: %v", err), "Failed to send contact number verification OTP.")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Contact number verification OTP sent successfully."})
}

// VerifyContactNumber verifies the user's contact number using the provided OTP.
func (c *AuthController) VerifyContactNumber(ctx *gin.Context) {
	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "User not authenticated.")
		return
	}

	var req auth_requests.VerifyContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("VerifyContactNumber: JSON binding error: %v", err), "Invalid request payload.")
		return
	}

	if err := req.Validate(); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("VerifyContactNumber: Validation error: %v", err), "Invalid input data.")
		return
	}

	isValid, err := c.otpService.ValidateOTP(userClaims.AccountType, userClaims.ID, req.Otp, "sms")
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("VerifyContactNumber: OTP validation error: %v", err), "Internal server error.")
		return
	}
	if !isValid {
		c.respondWithError(ctx, http.StatusUnauthorized, "VerifyContactNumber: Invalid or expired OTP", "Invalid or expired OTP.")
		return
	}

	updatedUser, err := c.userRepo.UpdateColumns(userClaims.AccountType, userClaims.ID, map[string]interface{}{
		"is_contact_verified": true,
	})
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("VerifyContactNumber: User update error: %v", err), "Failed to update user status.")
		return
	}

	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

// ChangeEmail updates the user's email and sends a verification email.
func (c *AuthController) ChangeEmail(ctx *gin.Context) {
	var req auth_requests.ChangeEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("ChangeEmail: JSON binding error: %v", err), "Invalid request payload.")
		return
	}

	if err := req.Validate(); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("ChangeEmail: Validation error: %v", err), "Invalid input data.")
		return
	}

	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "User not authenticated.")
		return
	}

	// Send Email OTP
	emailReq := services.EmailRequest{
		To:      req.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}
	if err := c.otpService.SendEmailOTP(userClaims.AccountType, userClaims.ID, emailReq); err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("ChangeEmail: Email OTP sending error: %v", err), "Failed to send verification email.")
		return
	}

	updatedUser, err := c.userRepo.UpdateColumns(userClaims.AccountType, userClaims.ID, map[string]interface{}{
		"is_email_verified": false,
		"email":             req.Email,
	})
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("ChangeEmail: User update error: %v", err), "Failed to update user email.")
		return
	}

	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

// ChangeContactNumber updates the user's contact number and sends a verification SMS.
func (c *AuthController) ChangeContactNumber(ctx *gin.Context) {
	var req auth_requests.ChangeContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("ChangeContactNumber: JSON binding error: %v", err), "Invalid request payload.")
		return
	}

	if err := req.Validate(); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("ChangeContactNumber: Validation error: %v", err), "Invalid input data.")
		return
	}

	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "User not authenticated.")
		return
	}

	currentUser, err := c.getCurrentUser(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "User not authenticated.")
		return
	}

	contactReq := services.SMSRequest{
		To:   req.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": fmt.Sprintf("%s %s", currentUser.FirstName, currentUser.LastName),
		},
	}
	if err := c.contactService.SendSMS(contactReq); err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("ChangeContactNumber: SMS sending error: %v", err), "Failed to send verification SMS.")
		return
	}

	updatedUser, err := c.userRepo.UpdateColumns(userClaims.AccountType, userClaims.ID, map[string]interface{}{
		"is_contact_verified": false,
		"contact_number":      req.ContactNumber,
	})
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("ChangeContactNumber: User update error: %v", err), "Failed to update contact number.")
		return
	}

	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

func (c *AuthController) SkipVerification(ctx *gin.Context) {
	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "User not authenticated.")
		return
	}
	updatedUser, err := c.userRepo.UpdateColumns(userClaims.AccountType, userClaims.ID, map[string]interface{}{
		"is_skip_verification": false,
	})
	if err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("ChangeContactNumber: User update error: %v", err), "Failed to update contact number.")
		return
	}
	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

func (c *AuthController) NewPassword(ctx *gin.Context) {

	var req auth_requests.NewPasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("Change Password: JSON binding error: %v", err), "Invalid request payload.")
		return
	}

	if err := req.Validate(); err != nil {
		c.respondWithError(ctx, http.StatusBadRequest, fmt.Sprintf("Change Password: Validation error: %v", err), "Invalid input data.")
		return
	}

	claims, err := c.getCurrentUser(ctx)
	if err != nil {
		c.respondWithError(ctx, http.StatusUnauthorized, err.Error(), "Authentication required.")
		return
	}

	if !config.VerifyPassword(claims.Password, req.PreviousPassword) {
		c.respondWithError(ctx, http.StatusUnauthorized, "SignIn: Password verification failed", "Invalid credentials.")
		return
	}

	if err := c.userRepo.UpdatePassword(claims.AccountType, claims.ID, req.ConfirmPassword); err != nil {
		c.respondWithError(ctx, http.StatusInternalServerError, fmt.Sprintf("ChangePassword: Password update error: %v", err), "Failed to update password.")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully."})
}

// AuthRoutes sets up the authentication routes.
func AuthRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware, controller *AuthController) {
	authGroup := router.Group("/auth")
	{
		// Public Routes
		authGroup.POST("/signup", controller.SignUp)
		authGroup.POST("/signin", controller.SignIn)
		authGroup.POST("/forgot-password", controller.ForgotPassword)
		authGroup.POST("/change-password", controller.ChangePassword)
		authGroup.GET("/verify-reset-link/:id", controller.VerifyResetLink)
		// Protected Routes
		authGroup.Use(mw.Middleware())
		{
			authGroup.GET("/current-user", controller.CurrentUser)
			authGroup.POST("/signout", controller.SignOut)

			authGroup.POST("/new-password", controller.NewPassword)

			authGroup.POST("/skip-verification", controller.SkipVerification)

			authGroup.POST("/send-email-verification", controller.SendEmailVerification)
			authGroup.POST("/verify-email", controller.VerifyEmail)
			authGroup.POST("/change-email", controller.ChangeEmail)

			authGroup.POST("/send-contact-number-verification", controller.SendContactNumberVerification)
			authGroup.POST("/verify-contact-number", controller.VerifyContactNumber)
			authGroup.POST("/change-contact-number", controller.ChangeContactNumber)
		}
	}
}
