package controllers

import (
	"horizon/server/config"
	"horizon/server/internal/auth"
	"horizon/server/internal/middleware"
	"horizon/server/internal/models"
	"horizon/server/internal/repositories"
	"horizon/server/internal/requests"
	"horizon/server/internal/requests/user_requests"
	"horizon/server/internal/resources"
	"net/http"

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

	claims, exists := ctx.Get("claims")
	if !exists {
		c.tokenService.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	userClaims, ok := claims.(*auth.UserClaims)
	if !ok {
		c.tokenService.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user claims."})
		return
	}

	user := &models.User{
		AccountType: userClaims.AccountType,
		MediaID:     &req.ID,
	}
	updatedUser, err := c.userRepo.UpdateColumns(userClaims.ID, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload profile picture to user status."})
		return
	}

	// Return updated user resource
	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

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

	claims, exists := ctx.Get("claims")
	if !exists {
		c.tokenService.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}

	userClaims, ok := claims.(*auth.UserClaims)
	if !ok {
		c.tokenService.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user claims."})
		return
	}
	user := &models.User{
		AccountType:      userClaims.AccountType,
		Birthdate:        req.Birthdate,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Description:      req.Description,
		PermanentAddress: req.PermanentAddress,
	}
	updatedUser, err := c.userRepo.UpdateColumns(userClaims.ID, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload profile picture to user status."})
		return
	}

	// Return updated user resource
	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

func (c *UserController) ChangePasswordSetting(ctx *gin.Context) {
	var req user_requests.ChangePasswordSettingRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := req.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	claims, exists := ctx.Get("claims")
	if !exists {
		c.tokenService.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated."})
		return
	}
	userClaims, ok := claims.(*auth.UserClaims)
	if !ok {
		c.tokenService.ClearTokenCookie(ctx)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user claims."})
		return
	}

	user, err := c.userRepo.GetByID(userClaims.AccountType, userClaims.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found."})
		return
	}

	if !config.VerifyPassword(user.Password, req.ConfirmPassword) {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Password verification failed."})
		return
	}
	if err := c.userRepo.UpdatePassword(userClaims.AccountType, userClaims.ID, req.ConfirmPassword); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password."})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully."})
}

func UserRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware, controller *UserController) {
	authGroup := router.Group("/profile")
	{
		authGroup.Use(mw.Middleware())
		authGroup.POST("/profile-picture", controller.ProfilePicture)
		authGroup.POST("/account-setting", controller.AccountSetting)
		authGroup.POST("/change-password", controller.ChangePasswordSetting)
	}
}
