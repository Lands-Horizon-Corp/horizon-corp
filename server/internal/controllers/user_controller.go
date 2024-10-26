package controllers

import (
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
		MediaID: &req.ID,
	}
	updatedUser, err := c.userRepo.UpdateColumns(userClaims.ID, *user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload profile picture to user status."})
		return
	}

	// Return updated user resource
	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

func (c *UserController) Description(ctx *gin.Context) {
	var req user_requests.DescriptionRequest

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
		Description: req.Description,
	}
	updatedUser, err := c.userRepo.UpdateColumns(userClaims.ID, *user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload profile picture to user status."})
		return
	}

	// Return updated user resource
	response := resources.ToResourceUser(updatedUser, userClaims.AccountType)
	ctx.JSON(http.StatusOK, response)
}

func UserRoutes(router *gin.RouterGroup, mw *middleware.AuthMiddleware, controller *UserController) {
	authGroup := router.Group("/profile")
	{
		authGroup.Use(mw.Middleware())
		authGroup.POST("/profile-picture", controller.ProfilePicture)
		authGroup.POST("/description", controller.Description)
	}
}
