package controllers

import (
	"horizon/server/config"
	"horizon/server/internal/auth"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests/auth_requests"
	"horizon/server/internal/resources"
	"horizon/server/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	adminRepo    *repositories.AdminRepository
	employeeRepo *repositories.EmployeeRepository
	ownerRepo    *repositories.OwnerRepository
	memberRepo   *repositories.MemberRepository
	otpService   *services.OTPService
	cfg          *config.AppConfig
	tokenService auth.TokenService
}

func NewAuthController(
	adminRepo *repositories.AdminRepository,
	employeeRepo *repositories.EmployeeRepository,
	ownerRepo *repositories.OwnerRepository,
	memberRepo *repositories.MemberRepository,
	otpService *services.OTPService,
	cfg *config.AppConfig,
	tokenService auth.TokenService,
) *AuthController {
	return &AuthController{
		adminRepo:    adminRepo,
		employeeRepo: employeeRepo,
		ownerRepo:    ownerRepo,
		memberRepo:   memberRepo,
		otpService:   otpService,
		cfg:          cfg,
		tokenService: tokenService,
	}
}

func (c *AuthController) CurrentUser(ctx *gin.Context) {}

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

	var userID uint
	hashed, err := config.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Password = hashed
	var response interface{}

	switch req.AccountType {
	case "Member":
		member := models.Member{
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
		if err := c.memberRepo.Create(&member); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = member.ID
		response = resources.ToResourceMember(member) // Use ToResourceAdmin for the member

	case "Owner":
		owner := models.Owner{
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
		if err := c.ownerRepo.Create(&owner); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = owner.ID
		response = resources.ToResourceOwner(owner) // Use ToResourceAdmin for the owner

	case "Employee":
		employee := models.Employee{
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
		if err := c.employeeRepo.Create(&employee); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = employee.ID
		response = resources.ToResourceEmployee(employee) // Use ToResourceAdmin for the employee

	case "Admin":
		admin := models.Admin{
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
		if err := c.adminRepo.Create(&admin); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = admin.ID
		response = resources.ToResourceAdmin(admin) // Use ToResourceAdmin for the admin

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
		return
	}

	// Token generation
	token, err := c.tokenService.GenerateToken(&auth.UserClaims{ID: userID, Mode: req.AccountType})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Sending OTP via email
	emailReq := services.EmailRequest{
		To:      req.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}
	if err := c.otpService.SendEmailOTP(req.AccountType, userID, emailReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP via email"})
		return
	}

	// Sending OTP via SMS
	contactReq := services.SMSRequest{
		To:   req.ContactNumber,
		Body: req.ContactTemplate,
		Vars: &map[string]string{
			"name": req.FirstName + " " + req.LastName,
		},
	}
	if err := c.otpService.SendEContactNumberOTP(req.AccountType, userID, contactReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP via SMS"})
		return
	}

	// Set token in cookie
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

func (c *AuthController) SignIn(ctx *gin.Context)  {}
func (c *AuthController) SignOut(ctx *gin.Context) {}

// Password
func (c *AuthController) ForgotPassword(ctx *gin.Context) {}
func (c *AuthController) ChangePassword(ctx *gin.Context) {}

// Email
func (c *AuthController) ChangeEmail(ctx *gin.Context)           {}
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
