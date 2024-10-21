package controllers

import (
	"fmt"
	"horizon/server/config"
	"horizon/server/helpers"
	"horizon/server/internal/auth"
	"horizon/server/internal/middleware"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests/auth_requests"
	"horizon/server/internal/resources"
	"horizon/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {

	// Database
	userRepo *repositories.UserRepository

	// Auth
	userAuthService *auth.UserAuthService

	// App
	cfg *config.AppConfig

	// Services
	emailService   *services.EmailService
	contactService *services.SMSService
	otpService     *services.OTPService
	tokenService   auth.TokenService
}

func NewAuthController(
	userRepo *repositories.UserRepository,

	// Auth
	userAuthService *auth.UserAuthService,

	// App
	cfg *config.AppConfig,

	// Services
	emailService *services.EmailService,
	contactService *services.SMSService,
	otpService *services.OTPService,
	tokenService auth.TokenService,

) *AuthController {
	return &AuthController{
		userRepo:        userRepo,
		userAuthService: userAuthService,

		cfg:            cfg,
		emailService:   emailService,
		contactService: contactService,
		otpService:     otpService,
		tokenService:   tokenService,
	}
}
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
	response := resources.ToResourceUser(user, user.AccountType)
	token, err := c.userAuthService.GenerateUserToken(user, user.AccountType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	emailReq := services.EmailRequest{
		To:      req.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}
	if err := c.otpService.SendEmailOTP(req.AccountType, user.ID, emailReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email OTP. Please try again later"})
		return
	}
	contactReq := services.SMSRequest{
		To:   req.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": req.FirstName + " " + req.LastName,
		},
	}
	if err := c.contactService.SendSMS(contactReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed send OTP from your number. Please try again"})
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     c.cfg.AppTokenName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	// Final response
	ctx.JSON(http.StatusCreated, response)
}

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

	isValid := config.VerifyPassword(user.Password, req.Password)
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := c.userAuthService.GenerateUserToken(user, user.AccountType)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     c.cfg.AppTokenName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	// Successful sign-in
	ctx.JSON(http.StatusOK, user)
}

func (c *AuthController) SignOut(ctx *gin.Context) {
	c.tokenService.ClearTokenCookie(ctx)
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully signed out"})
}

// Password
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
	keyType := helpers.GetKeyType(req.Key)

	resetLink := fmt.Sprintf("%s/auth/password-reset?token=%s", c.cfg.AppClientUrl, token)
	if keyType == "contact" {
		contactReq := services.SMSRequest{
			To:   user.ContactNumber,
			Body: req.ContactTemplate,
			Vars: &map[string]string{
				"name":      user.FirstName + " " + user.LastName,
				"eventLink": resetLink,
			},
		}
		if err := c.contactService.SendSMS(contactReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed send OTP from your number. Please try again"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Sent you the mail. Check your inbox or spam folder"})
	} else if keyType == "email" {
		emailReq := services.EmailRequest{
			To:      user.Email,
			Subject: "ECOOP: Change Password Request",
			Body:    req.EmailTemplate,
			Vars: &map[string]string{
				"name":      user.FirstName + " " + user.LastName,
				"eventLink": resetLink,
			},
		}
		if err := c.emailService.SendEmail(emailReq); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed send email OTP. Please try again"})
			return
		}
	}

}

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

func (c *AuthController) SendEmailVerification(ctx *gin.Context) {
	// Bind and validate the request
	var req auth_requests.SendEmailVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request: " + err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Validation error: " + err.Error()})
		return
	}

	// Get the current authenticated user from the context
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

	// Prepare the email request
	emailReq := services.EmailRequest{
		To:      currentUser.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}

	// Get the user claims from the context
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

	// Send the email OTP
	if err := c.otpService.SendEmailOTP(userClaims.AccountType, currentUser.ID, emailReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email OTP. Please try again later"})
		return
	}

	// Respond with success
	ctx.JSON(http.StatusOK, gin.H{"message": "Email verification sent successfully. Check your inbox or spam folder"})
}

func (c *AuthController) VerifyEmail(ctx *gin.Context) {
	user, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userDetails, ok := user.(*auth.UserClaims)
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
	isValid, err := c.otpService.ValidateOTP(userDetails.AccountType, userDetails.ID, req.Otp, "email")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}
	response, err := c.userRepo.UpdateColumns(userDetails.AccountType, userDetails.ID, map[string]interface{}{
		"is_email_verified": true,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member"})
		return
	}

	ctx.JSON(http.StatusOK, resources.ToResourceUser(response, userDetails.AccountType))
}

// Contact Number
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user user"})
		return
	}

	claims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	getClaims, err := claims.(*auth.UserClaims)
	if !err {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
		return
	}

	contactReq := services.SMSRequest{
		To:   currentUser.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": currentUser.FirstName + " " + currentUser.LastName,
		},
	}
	if err := c.otpService.SendEContactNumberOTP(getClaims.AccountType, getClaims.ID, contactReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed send OTP from your number. Please try again"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Sent you the mail. Check your inbox or spam folder"})
}

func (c *AuthController) VerifyContactNumber(ctx *gin.Context) {
	user, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	userDetails, ok := user.(*auth.UserClaims)
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
	isValid, err := c.otpService.ValidateOTP(userDetails.AccountType, userDetails.ID, req.Otp, "sms")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired OTP"})
		return
	}
	updated, err := c.userRepo.UpdateColumns(userDetails.AccountType, userDetails.ID, map[string]interface{}{
		"is_contact_verified": true,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member"})
		return
	}
	ctx.JSON(http.StatusOK,
		resources.ToResourceUser(updated, userDetails.AccountType))
}

func AuthRoutes(router *gin.RouterGroup, middleware *middleware.AuthMiddleware, controller *AuthController) {
	group := router.Group("/auth")
	{
		// Basic Authentication
		group.POST("/signup", controller.SignUp)
		group.POST("/signin", controller.SignIn)
		// Password
		group.POST("/forgot-password", controller.ForgotPassword)
		group.POST("/change-password", controller.ChangePassword)

		group.Use(middleware.Middleware())
		{
			// Route for getting the current user information
			group.GET("/current-user", controller.CurrentUser)
			group.POST("/signout", controller.SignOut)

			// Email
			group.POST("/send-email-verification", controller.SendEmailVerification)
			group.POST("/verify-email", controller.VerifyEmail)

			// Contact Number
			group.POST("/send-contact-number-verification", controller.SendContactNumberVerification)
			group.POST("/verify-contact-number", controller.VerifyContactNumber)
		}

	}
}
