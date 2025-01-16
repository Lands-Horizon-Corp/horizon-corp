package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type OwnerController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewOwnerController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *OwnerController {
	return &OwnerController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

// GET :/api/v1/owner/
// For owner only find all owner
func (c *OwnerController) Index(ctx *gin.Context) {

}

// GET :/api/v1/owner/id
// for owner only and self for owner
func (c *OwnerController) Show(ctx *gin.Context) {

}

// POST: /api/v1/owner/
// for owner only create owner
func (c *OwnerController) Store(ctx *gin.Context) {

}

// PUT: /api/v1/owner/:id
func (c *OwnerController) Update(ctx *gin.Context) {

}

// DELETE: /
// For owner only delete owner but if no branch and company and no employee or members
func (c *OwnerController) Destroy(ctx *gin.Context) {

}

// POST: /forgot-password
// owner: only  for email, phone number, and actual link
// Public: ownly  for email, and phone number
func (c *OwnerController) ForgotPassword(ctx *gin.Context) {

}
