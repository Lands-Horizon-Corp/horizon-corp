package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type AdminController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewAdminController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *AdminController {
	return &AdminController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

// GET: /api/v1/admin
// Retrieve admins with optional filtering for pagination or no pagination. Results can be converted to records.
//
//	Admin: Can retrieve all admins if admin  status is verified
//	Employee: not allowed
//	Member: not allowed
func (c *AdminController) Index(ctx *gin.Context) {

}

// GET: /api/v1/admin/:id
//
//	Admin: if admin  status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *AdminController) Show(ctx *gin.Context) {
	// if id or email or
}

// POST: /api/v1/admin
//
//	Admin: Can create admin and automatically verified and also if admin and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
//	public allowed
func (c *AdminController) Store(ctx *gin.Context) {

}

// PUT: /api/v1/admin/:id
//
//	Admin: Can change status of admin but only if admin and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *AdminController) Update(ctx *gin.Context) {

}

// DELETE:/api/v1/admin/:id
// Verifiy admin
//
//	Admin: only if admin  status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *AdminController) Destroy(ctx *gin.Context) {

}

func (c *AdminController) ForgotPassword(ctx *gin.Context) {

}
