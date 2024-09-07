package handlers

import (
	"net/http"

	"horizon-server/internal/models"
	"horizon-server/internal/services"
	"horizon-server/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService services.UserService
}

// NewUserHandler creates a new instance of UserHandler.
func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// RegisterUser handles user registration.

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var req struct {
		Username        string `json:"username"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
		Email           string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.userService.Register(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
func (h *UserHandler) LoginUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, err := h.userService.Login(username, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create a token or session ID here (implementation depends on your auth system)
	sessionToken, err := helpers.CreateSessionToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session token"})
		return
	}

	c.SetCookie("session_token", sessionToken, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h *UserHandler) LogoutUser(c *gin.Context) {
	c.SetCookie("session_token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// ChangeUserPassword handles changing the password for an authenticated user.
func (h *UserHandler) ChangeUserPassword(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in request"})
		return
	}

	var passwordDetails struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&passwordDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	err := h.userService.ChangePassword(userID.(uint), passwordDetails.OldPassword, passwordDetails.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to change password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

// GetUserProfile retrieves the profile of the authenticated user.
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in request"})
		return
	}

	user, err := h.userService.GetUserByID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return the user excluding the password
	c.JSON(http.StatusOK, gin.H{
		"id":                user.ID,
		"username":          user.Username,
		"email":             user.Email,
		"first_name":        user.FirstName,
		"last_name":         user.LastName,
		"permanent_address": user.PermanentAddress,
		"description":       user.Description,
		"birthdate":         user.Birthdate,
	})
}

// UpdateUserProfile updates the profile for the authenticated user.
func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in request"})
		return
	}

	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data provided"})
		return
	}

	updatedUser.ID = userID.(uint) // Ensure the ID is the authenticated user's ID to prevent unauthorized updates.
	err := h.userService.UpdateUser(&updatedUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}
