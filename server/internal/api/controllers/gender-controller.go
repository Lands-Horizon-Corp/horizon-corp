package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type GenderController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewGenderController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *GenderController {
	return &GenderController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

// GET: /
// Retrieve genderes with optional filtering for pagination or no pagination. Results can be converted to records.
//
//	Admin: Allowed
//	Employee: Only the gender of current gender
//	Owner: Only own gender
//	Member: Only own gender
//
// Endpoint: GET /api/v1/gender
func (c *GenderController) Index(ctx *gin.Context) {}

// GET: /api/v1/gender/:id
//
//	Admin: Allowed
//	Employee: Only the gender of current gender
//	Owner: Only own gender
//	Member: Only own gender
//
// Endpoint: GET /api/v1/gender/:id
func (c *GenderController) Show(ctx *gin.Context) {}

//	Admin: Can create gender but need to assign owner
//	Owner: Can create gender and automaticall assign owner
//	Employee: Not allowed
//	Member: Only own gender
//
// Endpoint: POST /api/v1/gender
func (c *GenderController) Store(ctx *gin.Context) {}

// Update modifies the details of a specific gender by ID.
//
// - Admin: Allowed.
// - Owner: Allowed but only his gender
// - Employee: Not allowed.
// - Member: Not allowed.
// Endpoint: PUT /api/v1/gender/:id
func (c *GenderController) Update(ctx *gin.Context) {}

// DELETE: /:id
// Admin: Allowed but if no members on genderes, no genderes, and no employees on genderes.
// Owner: Only own gender but if no members on genderes, no genderes, and no employees on genderes.
// Endpoint: DELETE /api/v1/gender/:id
func (c *GenderController) Destroy(ctx *gin.Context) {}
