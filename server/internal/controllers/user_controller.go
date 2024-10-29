package controllers

import (
	"errors"
	"net/http"

	"horizon/server/config"
	"horizon/server/internal/auth"
	"horizon/server/internal/middleware"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/requests/user_requests"
	"horizon/server/internal/resources"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepo     *repositories.UserRepository
	tokenService auth.TokenService
}

// NewUserController creates a new instance of UserController.
func NewUserController(userRepo *repositories.UserRepository, tokenService auth.TokenService) *UserController {
	return &UserController{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

// getUserClaims retrieves user claims from the context.
func (c *UserController) getUserClaims(ctx *gin.Context) (*auth.UserClaims, error) {
	claims, exists := ctx.Get("claims")
	if !exists {
		c.tokenService.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return nil, errors.New("user not authenticated")
	}

	userClaims, ok := claims.(*auth.UserClaims)
	if !ok {
		c.tokenService.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user claims."})
		return nil, errors.New("failed to retrieve user claims")
	}

	return userClaims, nil
}

// getUser retrieves the user model based on user claims.
func (c *UserController) getUser(ctx *gin.Context) (*models.User, error) {
	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		return nil, err
	}

	user, err := c.userRepo.GetByID(userClaims.AccountType, userClaims.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
		return nil, err
	}

	return user, nil
}

// verifyPassword checks if the provided password matches the user's password.
func (c *UserController) verifyPassword(ctx *gin.Context, user *models.User, password string) bool {
	if !config.VerifyPassword(user.Password, password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Password verification failed."})
		return false
	}
	return true
}

// updateUser updates the user and sends the response.
func (c *UserController) updateUser(ctx *gin.Context, userUpdate *models.User) {
	userClaims, _ := c.getUserClaims(ctx)

	updatedUser, err := c.userRepo.UpdateColumns(userClaims.ID, userUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user."})
		return
	}

	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

// handleRequest handles the request binding and validation.
func (c *UserController) handleRequest(ctx *gin.Context, req requests.Validatable) bool {
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false
	}
	return true
}

func (c *UserController) ProfilePicture(ctx *gin.Context) {
	var req requests.MediaRequest
	if !c.handleRequest(ctx, &req) {
		return
	}

	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		return
	}

	userUpdate := &models.User{
		AccountType: userClaims.AccountType,
		MediaID:     &req.ID,
	}

	c.updateUser(ctx, userUpdate)
}

// AccountSetting updates the user's account settings.
func (c *UserController) AccountSetting(ctx *gin.Context) {
	var req user_requests.AccountSettingRequest
	if !c.handleRequest(ctx, &req) {
		return
	}

	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		return
	}

	userUpdate := &models.User{
		AccountType:      userClaims.AccountType,
		Birthdate:        req.Birthdate,
		MiddleName:       req.MiddleName,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Description:      req.Description,
		PermanentAddress: req.PermanentAddress,
	}

	c.updateUser(ctx, userUpdate)
}

// ChangeEmail changes the user's email address.
func (c *UserController) ChangeEmail(ctx *gin.Context) {
	var req user_requests.ChangeEmailRequest
	if !c.handleRequest(ctx, &req) {
		return
	}

	user, err := c.getUser(ctx)
	if err != nil {
		return
	}

	if !c.verifyPassword(ctx, user, req.Password) {
		return
	}

	userUpdate := &models.User{
		AccountType:     user.AccountType,
		Email:           req.Email,
		IsEmailVerified: false,
	}

	c.updateUser(ctx, userUpdate)
}

// ChangeContactNumber changes the user's contact number.
func (c *UserController) ChangeContactNumber(ctx *gin.Context) {
	var req user_requests.ChangeContactNumberRequest
	if !c.handleRequest(ctx, &req) {
		return
	}

	user, err := c.getUser(ctx)
	if err != nil {
		return
	}

	if !c.verifyPassword(ctx, user, req.Password) {
		return
	}

	userUpdate := &models.User{
		AccountType:       user.AccountType,
		ContactNumber:     req.ContactNumber,
		IsContactVerified: false,
	}

	c.updateUser(ctx, userUpdate)
}

// ChangeUsername changes the user's username.
func (c *UserController) ChangeUsername(ctx *gin.Context) {
	var req user_requests.ChangeUsernameRequest
	if !c.handleRequest(ctx, &req) {
		return
	}

	user, err := c.getUser(ctx)
	if err != nil {
		return
	}

	if !c.verifyPassword(ctx, user, req.Password) {
		return
	}

	userUpdate := &models.User{
		AccountType: user.AccountType,
		Username:    req.Username,
	}

	c.updateUser(ctx, userUpdate)
}

// UserRoutes sets up the user-related routes.
func UserRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware, controller *UserController) {
	authGroup := router.Group("/profile")
	authGroup.Use(mw.Middleware())
	{
		authGroup.POST("/profile-picture", controller.ProfilePicture)
		authGroup.POST("/account-setting", controller.AccountSetting)
		authGroup.POST("/change-email", controller.ChangeEmail)
		authGroup.POST("/change-contact-number", controller.ChangeContactNumber)
		authGroup.POST("/change-username", controller.ChangeUsername)
	}
}
