package controllers

import (
	"fmt"
	"net/http"

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

// CurrentUser retrieves the current authenticated user.
func (c *AuthController) CurrentUser(ctx *gin.Context) {
	user, exists := ctx.Get("current-user")
	if !exists {
		c.tokenService.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required. Please log in."})
		return
	}

	currentUser, ok := user.(models.User)
	if !ok {
		c.tokenService.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user details. Please try again later."})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceUser(currentUser, currentUser.AccountType))
}

// SignUp handles user registration.
func (c *AuthController) SignUp(ctx *gin.Context) {
	var req auth_requests.SignUpRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := c.userAuthService.GenerateUserToken(user, user.AccountType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Send Email OTP
	emailReq := services.EmailRequest{
		To:      req.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}
	if err := c.otpService.SendEmailOTP(req.AccountType, user.ID, emailReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email OTP. Please try again later"})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP to your number. Please try again"})
		return
	}

	// Set Auth Token Cookie
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     c.cfg.AppTokenName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	// Respond with the created user
	response := resources.ToResourceUser(user, user.AccountType)
	ctx.JSON(http.StatusCreated, response)
}

// SignIn handles user authentication.
func (c *AuthController) SignIn(ctx *gin.Context) {
	var req auth_requests.SignInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userRepo.FindByEmailUsernameOrContact(req.AccountType, req.Key)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if !config.VerifyPassword(user.Password, req.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := c.userAuthService.GenerateUserToken(user, user.AccountType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set Auth Token Cookie
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     c.cfg.AppTokenName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	// Respond with the authenticated user
	ctx.JSON(http.StatusOK, resources.ToResourceUser(user, user.AccountType))
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userRepo.FindByEmailUsernameOrContact(req.AccountType, req.Key)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	token, err := c.userAuthService.GenerateUserToken(user, req.AccountType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate reset token"})
		return
	}

	resetLink := fmt.Sprintf("%s/auth/password-reset?token=%s", c.cfg.AppClientUrl, token)
	keyType := helpers.GetKeyType(req.Key)

	if keyType == "contact" {
		contactReq := services.SMSRequest{
			To:   user.ContactNumber,
			Body: req.ContactTemplate,
			Vars: &map[string]string{
				"name":      fmt.Sprintf("%s %s", user.FirstName, user.LastName),
				"eventLink": resetLink,
			},
		}
		if err := c.contactService.SendSMS(contactReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP to your number. Please try again"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Password reset instructions have been sent via SMS"})
	} else if keyType == "email" {
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
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email OTP. Please try again"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Password reset instructions have been sent via email"})
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid key type"})
	}
}

// ChangePassword updates the user's password.
func (c *AuthController) ChangePassword(ctx *gin.Context) {
	var req auth_requests.ChangePasswordRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := c.tokenService.VerifyToken(req.ResetID)
	if err != nil {
		c.tokenService.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired reset token"})
		return
	}

	if err := c.userRepo.UpdatePassword(claims.AccountType, claims.ID, req.ConfirmPassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

// SendEmailVerification sends an email verification OTP to the user.
func (c *AuthController) SendEmailVerification(ctx *gin.Context) {
	var req auth_requests.SendEmailVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation error: " + err.Error()})
		return
	}

	user, exists := ctx.Get("current-user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	currentUser, ok := user.(models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user data"})
		return
	}

	emailReq := services.EmailRequest{
		To:      currentUser.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}

	claims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userClaims, ok := claims.(*auth.UserClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user claims"})
		return
	}

	if err := c.otpService.SendEmailOTP(userClaims.AccountType, currentUser.ID, emailReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email OTP. Please try again later"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Email verification sent successfully. Check your inbox or spam folder"})
}

// VerifyEmail verifies the user's email using the provided OTP.
func (c *AuthController) VerifyEmail(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userClaims, ok := claims.(*auth.UserClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
		return
	}

	var req auth_requests.VerifyEmailRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, err := c.otpService.ValidateOTP(userClaims.AccountType, userClaims.ID, req.Otp, "email")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	updatedUser, err := c.userRepo.UpdateColumns(userClaims.AccountType, userClaims.ID, map[string]interface{}{
		"is_email_verified": true,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user status"})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceUser(updatedUser, userClaims.AccountType))
}

// SendContactNumberVerification sends an SMS verification OTP to the user's contact number.
func (c *AuthController) SendContactNumberVerification(ctx *gin.Context) {
	var req auth_requests.SendContactNumberVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, exists := ctx.Get("current-user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	currentUser, ok := user.(models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user data"})
		return
	}

	claims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userClaims, ok := claims.(*auth.UserClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user claims"})
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP to your number. Please try again"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Contact number verification OTP sent successfully"})
}

// VerifyContactNumber verifies the user's contact number using the provided OTP.
func (c *AuthController) VerifyContactNumber(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userClaims, ok := claims.(*auth.UserClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
		return
	}

	var req auth_requests.VerifyContactNumberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, err := c.otpService.ValidateOTP(userClaims.AccountType, userClaims.ID, req.Otp, "sms")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}

	updatedUser, err := c.userRepo.UpdateColumns(userClaims.AccountType, userClaims.ID, map[string]interface{}{
		"is_contact_verified": true,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user status"})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceUser(updatedUser, userClaims.AccountType))
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

		// Protected Routes
		authGroup.Use(mw.Middleware())
		{
			authGroup.GET("/current-user", controller.CurrentUser)
			authGroup.POST("/signout", controller.SignOut)
			authGroup.POST("/send-email-verification", controller.SendEmailVerification)
			authGroup.POST("/verify-email", controller.VerifyEmail)
			authGroup.POST("/send-contact-number-verification", controller.SendContactNumberVerification)
			authGroup.POST("/verify-contact-number", controller.VerifyContactNumber)
		}
	}
}
