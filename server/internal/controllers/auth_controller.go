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
	"time"

	"github.com/gin-gonic/gin"
)

type AuthController struct {

	// Database
	adminRepo    *repositories.AdminRepository
	employeeRepo *repositories.EmployeeRepository
	ownerRepo    *repositories.OwnerRepository
	memberRepo   *repositories.MemberRepository

	// Auth
	adminAuthService    *auth.AdminAuthService
	employeeAuthService *auth.EmployeeAuthService
	memberAuthService   *auth.MemberAuthService
	ownerAuthService    *auth.OwnerAuthService

	// App
	cfg *config.AppConfig

	// Services
	emailService   *services.EmailService
	contactService *services.SMSService
	otpService     *services.OTPService
	tokenService   auth.TokenService
}

func NewAuthController(
	// Database
	adminRepo *repositories.AdminRepository,
	employeeRepo *repositories.EmployeeRepository,
	ownerRepo *repositories.OwnerRepository,
	memberRepo *repositories.MemberRepository,

	// Auth
	adminAuthService *auth.AdminAuthService,
	employeeAuthService *auth.EmployeeAuthService,
	memberAuthService *auth.MemberAuthService,
	ownerAuthService *auth.OwnerAuthService,

	// App
	cfg *config.AppConfig,

	// Services
	emailService *services.EmailService,
	contactService *services.SMSService,
	otpService *services.OTPService,
	tokenService auth.TokenService,

) *AuthController {
	return &AuthController{
		adminRepo:           adminRepo,
		employeeRepo:        employeeRepo,
		ownerRepo:           ownerRepo,
		memberRepo:          memberRepo,
		adminAuthService:    adminAuthService,
		employeeAuthService: employeeAuthService,
		memberAuthService:   memberAuthService,
		ownerAuthService:    ownerAuthService,
		cfg:                 cfg,
		emailService:        emailService,
		contactService:      contactService,
		otpService:          otpService,
		tokenService:        tokenService,
	}
}
func (c *AuthController) CurrentUser(ctx *gin.Context) {
	user, exists := ctx.Get("current-user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication required. Please log in."})
		return
	}
	currentUser, ok := user.(models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to retrieve user details. Please try again later."})
		return
	}

	switch currentUser.AccountType {
	case "Owner":
		owner, _ := user.(models.Owner)
		resource := resources.ToResourceOwner(owner)
		ctx.JSON(http.StatusOK, resource)
		return
	case "Employee":
		employee, _ := user.(models.Employee)
		resource := resources.ToResourceEmployee(employee)
		ctx.JSON(http.StatusOK, resource)
		return
	case "Admin":
		admin, _ := user.(models.Admin)
		resource := resources.ToResourceAdmin(admin)
		ctx.JSON(http.StatusOK, resource)
		return
	case "Member":
		member, _ := user.(models.Member)

		resource := resources.ToResourceMember(member)
		ctx.JSON(http.StatusOK, resource)
		return
	}

	ctx.JSON(http.StatusForbidden, gin.H{"error": "Access denied. You do not have permission to view this resource."})
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

	// Hash the password
	hashed, err := config.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Password = hashed

	var userID uint
	var response interface{}
	var token string

	// Create user based on account type
	switch req.AccountType {
	case "Member":
		member := c.createMember(req)
		if err := c.memberRepo.Create(&member); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = member.ID
		token, err = c.memberAuthService.GenerateMemberToken(member)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		response = resources.ToResourceMember(member)

	case "Owner":
		owner := c.createOwner(req)
		if err := c.ownerRepo.Create(&owner); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = owner.ID
		token, err = c.ownerAuthService.GenerateOwnerToken(owner)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		response = resources.ToResourceOwner(owner)

	case "Employee":
		employee := c.createEmployee(req)
		if err := c.employeeRepo.Create(&employee); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = employee.ID
		token, err = c.employeeAuthService.GenerateEmployeeToken(employee)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		response = resources.ToResourceEmployee(employee)

	case "Admin":
		admin := c.createAdmin(req)
		if err := c.adminRepo.Create(&admin); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = admin.ID
		token, err = c.adminAuthService.GenerateAdminToken(admin)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}
		response = resources.ToResourceAdmin(admin)

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
		return
	}

	// Send OTP via email
	if err := c.sendEmailOTP(req, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send OTP via SMS
	if err := c.sendSMSOTP(req, userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	var password string
	var user interface{}
	var token string

	switch req.AccountType {
	case "Member":
		member, err := c.memberRepo.FindByEmailUsernameOrContact(req.Key)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		password = member.Password
		user = resources.ToResourceMember(member)
		token, err = c.memberAuthService.GenerateMemberToken(member)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

	case "Owner":
		owner, err := c.ownerRepo.FindByEmailUsernameOrContact(req.Key)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		password = owner.Password
		user = resources.ToResourceOwner(owner)
		token, err = c.ownerAuthService.GenerateOwnerToken(owner)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

	case "Employee":
		employee, err := c.employeeRepo.FindByEmailUsernameOrContact(req.Key)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		password = employee.Password
		user = resources.ToResourceEmployee(employee)
		token, err = c.employeeAuthService.GenerateEmployeeToken(employee)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

	case "Admin":
		admin, err := c.adminRepo.FindByEmailUsernameOrContact(req.Key)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		password = admin.Password
		user = resources.ToResourceAdmin(admin)
		token, err = c.adminAuthService.GenerateAdminToken(admin)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
		return
	}

	// Verify password
	isValid := config.VerifyPassword(password, req.Password)
	if !isValid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
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
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     c.cfg.AppTokenName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteNoneMode,
	})
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
	var userClaims *auth.UserClaims
	var user models.User
	switch req.AccountType {
	case "Member":
		member, err := c.memberRepo.FindByEmailUsernameOrContact(req.Key)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		user = models.User{
			ContactNumber: member.ContactNumber,
			FirstName:     member.FirstName,
			LastName:      member.LastName,
			Email:         member.Email,
		}
		userClaims = &auth.UserClaims{
			ID:          member.ID,
			AccountType: "Member",
		}
	case "Owner":
		owner, err := c.ownerRepo.FindByEmailUsernameOrContact(req.Key)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		user = models.User{
			ContactNumber: owner.ContactNumber,
			FirstName:     owner.FirstName,
			LastName:      owner.LastName,
			Email:         owner.Email,
		}
		userClaims = &auth.UserClaims{
			ID:          owner.ID,
			AccountType: "Owner",
		}
	case "Employee":
		employee, err := c.employeeRepo.FindByEmailUsernameOrContact(req.Key)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		user = models.User{
			ContactNumber: employee.ContactNumber,
			FirstName:     employee.FirstName,
			LastName:      employee.LastName,
			Email:         employee.Email,
		}
		userClaims = &auth.UserClaims{
			ID:          employee.ID,
			AccountType: "Employee",
		}
	case "Admin":
		admin, err := c.adminRepo.FindByEmailUsernameOrContact(req.Key)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			return
		}
		user = models.User{
			ContactNumber: admin.ContactNumber,
			FirstName:     admin.FirstName,
			LastName:      admin.LastName,
			Email:         admin.Email,
		}
		userClaims = &auth.UserClaims{
			ID:          admin.ID,
			AccountType: "Admin",
		}
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
		return
	}

	token, err := c.tokenService.GenerateToken(userClaims)
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
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired reset token"})
		return
	}
	hashedNewPassword, err := config.HashPassword(req.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash new password"})
		return
	}
	if err := c.updatePassword(claims.AccountType, claims.ID, hashedNewPassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

func (c *AuthController) SendEmailVerification(ctx *gin.Context) {
	var req auth_requests.SendEmailVerificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	user, exists := ctx.Get("current-user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}
	currentUser, ok := user.(models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user email"})
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
	getClaims, err := claims.(*auth.UserClaims)
	if !err {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user details"})
		return
	}

	if err := c.otpService.SendEmailOTP(getClaims.AccountType, currentUser.ID, emailReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed send email OTP. Please try again"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Sent you the mail. Check your inbox or spam folder"})
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
	var response interface{}

	switch userDetails.AccountType {
	case "Member":
		updatedMember, err := c.memberRepo.UpdateColumns(userDetails.ID, map[string]interface{}{
			"is_email_verified": true,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member"})
			return
		}
		response = resources.ToResourceMember(updatedMember)
	case "Owner":
		updatedOwner, err := c.ownerRepo.UpdateColumns(userDetails.ID, map[string]interface{}{
			"is_email_verified": true,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner"})
			return
		}
		response = resources.ToResourceOwner(updatedOwner)

	case "Employee":
		updatedEmployee, err := c.employeeRepo.UpdateColumns(userDetails.ID, map[string]interface{}{
			"is_email_verified": true,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
			return
		}
		response = resources.ToResourceEmployee(updatedEmployee)

	case "Admin":
		updatedAdmin, err := c.adminRepo.UpdateColumns(userDetails.ID, map[string]interface{}{
			"is_email_verified": true,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin"})
			return
		}
		response = resources.ToResourceAdmin(updatedAdmin)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
		return
	}
	ctx.JSON(http.StatusOK, response)
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
	var response interface{}

	switch userDetails.AccountType {
	case "Member":
		updatedMember, err := c.memberRepo.UpdateColumns(userDetails.ID, map[string]interface{}{
			"is_contact_verified": true,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update member"})
			return
		}
		response = resources.ToResourceMember(updatedMember)
	case "Owner":
		updatedOwner, err := c.ownerRepo.UpdateColumns(userDetails.ID, map[string]interface{}{
			"is_contact_verified": true,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update owner"})
			return
		}
		response = resources.ToResourceOwner(updatedOwner)

	case "Employee":
		updatedEmployee, err := c.employeeRepo.UpdateColumns(userDetails.ID, map[string]interface{}{
			"is_contact_verified": true,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update employee"})
			return
		}
		response = resources.ToResourceEmployee(updatedEmployee)

	case "Admin":
		updatedAdmin, err := c.adminRepo.UpdateColumns(userDetails.ID, map[string]interface{}{
			"is_contact_verified": true,
		})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin"})
			return
		}
		response = resources.ToResourceAdmin(updatedAdmin)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
		return
	}
	ctx.JSON(http.StatusOK, response)
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

func (c *AuthController) updatePassword(accountType string, userID uint, newPassword string) error {
	switch accountType {
	case "Member":
		_, err := c.memberRepo.UpdateColumns(userID, map[string]interface{}{
			"password": newPassword,
		})
		return err
	case "Owner":
		_, err := c.ownerRepo.UpdateColumns(userID, map[string]interface{}{
			"password": newPassword,
		})
		return err
	case "Employee":
		_, err := c.employeeRepo.UpdateColumns(userID, map[string]interface{}{
			"password": newPassword,
		})
		return err
	case "Admin":
		_, err := c.adminRepo.UpdateColumns(userID, map[string]interface{}{
			"password": newPassword,
		})
		return err
	default:
		return fmt.Errorf("invalid account type")
	}
}

// Helper functions to create user models
func (c *AuthController) createMember(req auth_requests.SignUpRequest) models.Member {
	return models.Member{
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
}

func (c *AuthController) createOwner(req auth_requests.SignUpRequest) models.Owner {
	return models.Owner{
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
}

func (c *AuthController) createEmployee(req auth_requests.SignUpRequest) models.Employee {
	return models.Employee{
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
}

func (c *AuthController) createAdmin(req auth_requests.SignUpRequest) models.Admin {
	return models.Admin{
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
}

// Sending email OTP
func (c *AuthController) sendEmailOTP(req auth_requests.SignUpRequest, userID uint) error {
	emailReq := services.EmailRequest{
		To:      req.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}
	return c.otpService.SendEmailOTP(req.AccountType, userID, emailReq)
}

// Sending SMS OTP
func (c *AuthController) sendSMSOTP(req auth_requests.SignUpRequest, userID uint) error {
	contactReq := services.SMSRequest{
		To:   req.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": req.FirstName + " " + req.LastName,
		},
	}
	return c.otpService.SendEContactNumberOTP(req.AccountType, userID, contactReq)
}
