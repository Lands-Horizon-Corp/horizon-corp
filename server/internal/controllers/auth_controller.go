package controllers

import (
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests/auth_requests"
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
}

func NewAuthController(
	adminRepo *repositories.AdminRepository,
	employeeRepo *repositories.EmployeeRepository,
	ownerRepo *repositories.OwnerRepository,
	memberRepo *repositories.MemberRepository,
	otpService *services.OTPService,
) *AuthController {
	return &AuthController{
		adminRepo:    adminRepo,
		employeeRepo: employeeRepo,
		ownerRepo:    ownerRepo,
		memberRepo:   memberRepo,
		otpService:   otpService,
	}
}

// User
func (c *AuthController) CurrentUser(ctx *gin.Context) {}

// Basic Authentication
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

	// Create the user based on AccountType and capture the userID
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
		}
		if err := c.memberRepo.Create(&member); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = member.ID

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
		}
		if err := c.ownerRepo.Create(&owner); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = owner.ID

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
		}
		if err := c.employeeRepo.Create(&employee); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = employee.ID

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
		}
		if err := c.adminRepo.Create(&admin); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		userID = admin.ID

	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type"})
		return
	}

	emailReq := services.EmailRequest{
		To:      req.Email,
		Subject: "ECOOP: Email Verification",
		Body:    req.EmailTemplate,
	}
	if err := c.otpService.SendEmailOTP(req.AccountType, userID, emailReq); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send OTP via email"})
		return
	}

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

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully! OTP sent to email and contact number."})
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
