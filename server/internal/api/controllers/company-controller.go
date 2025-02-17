package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewCompanyController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *CompanyController {
	return &CompanyController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

// GET: /
// Retrieve companyes with optional filtering for pagination or no pagination. Results can be converted to records.
//
//	Admin: Allowed
//	Employee: Only the company of current company
//	Owner: Only own company
//	Member: Only own company
//
// Endpoint: GET /api/v1/company
func (c *CompanyController) Index(ctx *gin.Context) {
	// Owner, Media
}

// GET: /api/v1/company/:id
//
//	Admin: Allowed
//	Employee: Only the company of current company
//	Owner: Only own company
//	Member: Only own company
//
// Endpoint: GET /api/v1/company/:id
func (c *CompanyController) Show(ctx *gin.Context) {}

//	Admin: Can create company but need to assign owner
//	Owner: Can create company and automaticall assign owner
//	Employee: Not Allowed
//	Member: Only own company
//
// Endpoint: POST /api/v1/company
func (c *CompanyController) Store(ctx *gin.Context) {}

// Update modifies the details of a specific company by ID.
//
// - Admin: Allowed.
// - Owner: Allowed but only his company
// - Employee: Not Allowed.
// - Member: Not Allowed.
// Endpoint: PUT /api/v1/company/:id
func (c *CompanyController) Update(ctx *gin.Context) {}

// DELETE: /:id
// Admin: Allowed but if no members on companyes, no companyes, and no employees on companyes.
// Owner: Only own company but if no members on companyes, no companyes, and no employees on companyes.
// Endpoint: DELETE /api/v1/company/:id
func (c *CompanyController) Destroy(ctx *gin.Context) {}

// Endpoint: GET /api/v1/company/current
func (c *CompanyController) CurrentCompany(ctx *gin.Context) {}
