package controllers

import (
	"errors"
	"horizon/server/config"
	"horizon/server/internal/auth"
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

	// Services
	otpService   *services.OTPService
	cfg          *config.AppConfig
	tokenService auth.TokenService
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

	// Services
	otpService *services.OTPService,
	cfg *config.AppConfig,
	tokenService auth.TokenService,

) *AuthController {
	return &AuthController{
		// Database
		adminRepo:    adminRepo,
		employeeRepo: employeeRepo,
		ownerRepo:    ownerRepo,
		memberRepo:   memberRepo,

		// Auth
		adminAuthService:    adminAuthService,
		employeeAuthService: employeeAuthService,
		memberAuthService:   memberAuthService,
		ownerAuthService:    ownerAuthService,

		// Services
		otpService:   otpService,
		cfg:          cfg,
		tokenService: tokenService,
	}
}

func (c *AuthController) CurrentUser(ctx *gin.Context) {
	// Retrieve the cookie
	cookie, err := ctx.Request.Cookie(c.cfg.AppTokenName)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Cookie not found"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Decrypt the token
	decrypted, err := config.Decrypt([]byte(cookie.Value), []byte(c.cfg.AppAdminToken))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// Verify the token
	claims, err := c.tokenService.VerifyToken(string(decrypted))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Token verification failed"})
		return
	}

	// Fetch user based on account type
	var response interface{}
	switch claims.AccountType {
	case "Member":
		member, err := c.memberRepo.GetByID(claims.ID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Member not found"})
			return
		}
		response = resources.ToResourceMember(member)

	case "Employee":
		employee, err := c.employeeRepo.GetByID(claims.ID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
			return
		}
		response = resources.ToResourceEmployee(employee)

	case "Admin":
		admin, err := c.adminRepo.GetByID(claims.ID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
			return
		}
		response = resources.ToResourceAdmin(admin)

	case "Owner":
		owner, err := c.ownerRepo.GetByID(claims.ID)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Owner not found"})
			return
		}
		response = resources.ToResourceOwner(owner)

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
		return
	}

	ctx.JSON(http.StatusOK, response)
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

	// Encrypt the token
	encrypted, err := config.Encrypt([]byte(token), []byte(c.cfg.AppAdminToken))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set the cookie with the encrypted token
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     c.cfg.AppTokenName,
		Value:    string(encrypted),
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

	// Encrypt the token
	encrypted, err := config.Encrypt([]byte(token), []byte(c.cfg.AppAdminToken))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set the cookie with the encrypted token
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     c.cfg.AppTokenName,
		Value:    string(encrypted),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	})

	// Successful sign-in
	ctx.JSON(http.StatusOK, user)
}

func (c *AuthController) SignOut(ctx *gin.Context) {
	cookie, err := ctx.Request.Cookie(c.cfg.AppTokenName)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfully signed out"})
		return
	}

	decryptedToken, err := config.Decrypt([]byte(cookie.Value), []byte(c.cfg.AppAdminToken))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Successfully signed out"})
		return
	}

	_, err = c.tokenService.VerifyToken(string(decryptedToken))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to invalidate token"})
		return
	}

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
func (c *AuthController) ForgotPassword(ctx *gin.Context) {}
func (c *AuthController) ChangePassword(ctx *gin.Context) {}

// Email
func (c *AuthController) ChangeEmail(ctx *gin.Context) {
	// if find email
	// if email exists return error

	// if check if email verified
	// if veriefied continue success

	// if hindi verified
	// send otp to email

}
func (c *AuthController) SendEmailVerification(ctx *gin.Context) {}
func (c *AuthController) VerifyEmail(ctx *gin.Context)           {}

// Contact Number
func (c *AuthController) ChangeContactNumber(ctx *gin.Context)           {}
func (c *AuthController) SendContactNumberVerification(ctx *gin.Context) {}
func (c *AuthController) VerifyContactNumber(ctx *gin.Context)           {}

func AuthRoutes(router *gin.RouterGroup, controller *AuthController) {
	group := router.Group("/auth")
	{
		// Basic Authentication
		group.GET("/current-user", controller.CurrentUser) // for signed in

		// Basic Authentication
		group.POST("/signup", controller.SignUp)   // for signed out
		group.POST("/signin", controller.SignIn)   // for signed out
		group.POST("/signout", controller.SignOut) // for signed in

		// Password
		group.POST("/forgot-password", controller.ForgotPassword) // for signed out
		group.POST("/change-password", controller.ChangePassword) // for signed out

		// Email
		group.POST("/change-email", controller.ChangeEmail)                      // for signed in
		group.POST("/send-email-verification", controller.SendEmailVerification) // for signed in
		group.POST("/verify-email", controller.VerifyEmail)                      // for signed in

		// Contact Number
		group.POST("/change-contact-number", controller.ChangeContactNumber)                      // for signed in
		group.POST("/send-contact-number-verification", controller.SendContactNumberVerification) // for signed in
		group.POST("/verify-otp", controller.VerifyContactNumber)                                 // for signed in
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

func (c *AuthController) RetrieveEmail(accountType string, userID uint) string {
	switch accountType {
	case "Member":
		member, _ := c.memberRepo.GetByID(userID)
		return member.Email
	case "Owner":
		owner, _ := c.ownerRepo.GetByID(userID)
		return owner.Email
	case "Employee":
		employee, _ := c.employeeRepo.GetByID(userID)
		return employee.Email
	case "Admin":
		admin, _ := c.adminRepo.GetByID(userID)
		return admin.Email
	default:
		return ""
	}
}

func (c *AuthController) RetrieveContactNumber(accountType string, userID uint) string {
	switch accountType {
	case "Member":
		member, _ := c.memberRepo.GetByID(userID)
		return member.ContactNumber
	case "Owner":
		owner, _ := c.ownerRepo.GetByID(userID)
		return owner.ContactNumber
	case "Employee":
		employee, _ := c.employeeRepo.GetByID(userID)
		return employee.ContactNumber
	case "Admin":
		admin, _ := c.adminRepo.GetByID(userID)
		return admin.ContactNumber
	default:
		return ""
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
