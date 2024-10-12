package controllers

import (
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests/auth_requests"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	adminRepo    *repositories.AdminRepository
	employeeRepo *repositories.EmployeeRepository
	ownerRepo    *repositories.OwnerRepository
	memberRepo   *repositories.MemberRepository
}

func NewAuthController(adminRepo *repositories.AdminRepository,
	employeeRepo *repositories.EmployeeRepository,
	ownerRepo *repositories.OwnerRepository,
	memberRepo *repositories.MemberRepository) *AuthController {
	return &AuthController{
		adminRepo:    adminRepo,
		employeeRepo: employeeRepo,
		ownerRepo:    ownerRepo,
		memberRepo:   memberRepo,
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

	// if req.AccountType == "Member" {
	// 	member := models.Member{
	// 		FirstName:         req.FirstName,
	// 		LastName:          req.LastName,
	// 		MiddleName:        req.MiddleName,
	// 		PermanentAddress:  req.PermanentAddress,
	// 		Description:       "",
	// 		Birthdate:         req.Birthdate,
	// 		Username:          req.Username,
	// 		Email:             req.Email,
	// 		Password:          req.Password,
	// 		IsEmailVerified:   false,
	// 		IsContactVerified: false,
	// 		ContactNumber:     req.ContactNumber,
	// 		MediaID:           nil,
	// 	}
	// 	if err := c.memberRepo.Create(&member); err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// } else if req.AccountType == "Owner" {
	// 	owner := models.Owner{
	// 		FirstName:         req.FirstName,
	// 		LastName:          req.LastName,
	// 		MiddleName:        req.MiddleName,
	// 		PermanentAddress:  req.PermanentAddress,
	// 		Description:       "",
	// 		Birthdate:         req.Birthdate,
	// 		Username:          req.Username,
	// 		Email:             req.Email,
	// 		Password:          req.Password,
	// 		IsEmailVerified:   false,
	// 		IsContactVerified: false,
	// 		ContactNumber:     req.ContactNumber,
	// 		MediaID:           nil,
	// 	}
	// 	if err := c.ownerRepo.Create(&owner); err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// } else if req.AccountType == "Employee" {
	// 	employee := models.Employee{
	// 		FirstName:         req.FirstName,
	// 		LastName:          req.LastName,
	// 		MiddleName:        req.MiddleName,
	// 		PermanentAddress:  req.PermanentAddress,
	// 		Description:       "",
	// 		Birthdate:         req.Birthdate,
	// 		Username:          req.Username,
	// 		Email:             req.Email,
	// 		Password:          req.Password,
	// 		IsEmailVerified:   false,
	// 		IsContactVerified: false,
	// 		ContactNumber:     req.ContactNumber,
	// 		MediaID:           nil,
	// 	}
	// 	if err := c.employeeRepo.Create(&employee); err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// } else if req.AccountType == "Admin" {
	// 	admin := models.Admin{
	// 		FirstName:         req.FirstName,
	// 		LastName:          req.LastName,
	// 		MiddleName:        req.MiddleName,
	// 		PermanentAddress:  req.PermanentAddress,
	// 		Description:       "",
	// 		Birthdate:         req.Birthdate,
	// 		Username:          req.Username,
	// 		Email:             req.Email,
	// 		Password:          req.Password,
	// 		IsEmailVerified:   false,
	// 		IsContactVerified: false,
	// 		ContactNumber:     req.ContactNumber,
	// 		MediaID:           nil,
	// 	}
	// 	if err := c.adminRepo.Create(&admin); err != nil {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// }

	// Send ng email notification
	// Send ng otp notification
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
