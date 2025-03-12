package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type ContactController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewContactController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *ContactController {
	return &ContactController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

// GET: /api/v1/contact
// Retrieve contact with optional filtering for pagination or no pagination. Results can be converted to records.
//
//	Admin: Allowed
//	Employee: Not Allowed
//	Owner: Not Allowed
//	Member: Not Allowed
func (c *ContactController) Index(ctx *gin.Context) {}

// GET: /api/v1/contact/:id
//
//	Admin: Allowed
//	Employee: Not Allowed
//	Owner: Not Allowed
//	Member: Not Allowed
//
// Endpoint: GET /api/v1/contact/:id
func (c *ContactController) Show(ctx *gin.Context) {}

// PUBLIC
// Endpoint: POST /api/v1/contact
func (c *ContactController) Store(ctx *gin.Context) {}

//	Admin: Allowed
//	Employee: Not Allowed
//	Owner: Not Allowed
//	Member: Not Allowed
//
// Endpoint: DELETE /api/v1/contact/:id
func (c *ContactController) Destroy(ctx *gin.Context) {}
