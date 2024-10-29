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

// ProfilePicture updates the user's profile picture.
func (c *UserController) ProfilePicture(ctx *gin.Context) {
	var req requests.MediaRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	updatedUser, err := c.userRepo.UpdateColumns(userClaims.ID, userUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile picture."})
		return
	}

	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

// AccountSetting updates the user's account settings.
func (c *UserController) AccountSetting(ctx *gin.Context) {
	var req user_requests.AccountSettingRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	updatedUser, err := c.userRepo.UpdateColumns(userClaims.ID, userUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update account settings."})
		return
	}

	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

// ChangeEmail changes the user's email address.
func (c *UserController) ChangeEmail(ctx *gin.Context) {
	var req user_requests.ChangeEmailRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		return
	}

	user, err := c.userRepo.GetByID(userClaims.AccountType, userClaims.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
		return
	}

	if !config.VerifyPassword(user.Password, req.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Password verification failed."})
		return
	}

	userUpdate := &models.User{
		AccountType:     userClaims.AccountType,
		Email:           req.Email,
		IsEmailVerified: false,
	}

	updatedUser, err := c.userRepo.UpdateColumns(userClaims.ID, userUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update email."})
		return
	}

	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

// ChangeContactNumber changes the user's contact number.
func (c *UserController) ChangeContactNumber(ctx *gin.Context) {
	var req user_requests.ChangeContactNumberRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userClaims, err := c.getUserClaims(ctx)
	if err != nil {
		return
	}

	user, err := c.userRepo.GetByID(userClaims.AccountType, userClaims.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
		return
	}

	if !config.VerifyPassword(user.Password, req.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Password verification failed."})
		return
	}

	userUpdate := &models.User{
		AccountType:       userClaims.AccountType,
		ContactNumber:     req.ContactNumber,
		IsContactVerified: false,
	}

	updatedUser, err := c.userRepo.UpdateColumns(userClaims.ID, userUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update contact number."})
		return
	}

	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

// UserRoutes registers user-related routes.
func UserRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware, controller *UserController) {
	authGroup := router.Group("/profile")
	authGroup.Use(mw.Middleware())
	{
		authGroup.POST("/profile-picture", controller.ProfilePicture)
		authGroup.POST("/account-setting", controller.AccountSetting)
		authGroup.POST("/change-email", controller.ChangeEmail)
		authGroup.POST("/change-contact-number", controller.ChangeContactNumber)
	}
}
