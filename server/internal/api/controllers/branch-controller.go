package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type BranchController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewBranchController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *BranchController {
	return &BranchController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

// Index retrieves a list of branches with optional filtering for pagination or no pagination.
//
// - Admin: Can retrieve all branches.
// - Employee: Can only retrieve the current branch.
// - Owner: Can retrieve all branches under their ownership.
// - Member: Can only retrieve the current branch.
//
// Endpoint: GET /api/v1/branch
func (c *BranchController) Index(ctx *gin.Context) {}

// Show retrieves details of a specific branch by ID.
//
// - Admin, Employee, Owner, and Member: Can retrieve details of the current branch.
//
// Endpoint: GET /api/v1/branch/:id
func (c *BranchController) Show(ctx *gin.Context) {}

// Store creates a new branch.
//
// - Admin: Allowed but must select a company first.
// - Owner: Allowed.
// - Employee: Not Allowed.
// - Member: Not Allowed.
//
// Endpoint: POST /api/v1/branch
func (c *BranchController) Store(ctx *gin.Context) {}

// Update modifies the details of a specific branch by ID.
//
// - Admin: Allowed.
// - Owner: Allowed.
// - Employee: Not Allowed.
// - Member: Not Allowed.
//
// Endpoint: PUT /api/v1/branch/:id
func (c *BranchController) Update(ctx *gin.Context) {}

// Destroy deletes a specific branch by ID.
//
// - Admin: Allowed but only if there are no members in the branch.
// - Owner: Allowed but only if there are no members in the branch.
// - Employee: Not Allowed.
// - Member: Not Allowed.
//
// Endpoint: DELETE /api/v1/branch/:id
func (c *BranchController) Destroy(ctx *gin.Context) {}

// NearestBranch retrieves branches near a specific location, filtered by radius or postal code, province, city, barangay, or region.
//
// - This returns all nearest and furthest branches based on the provided location parameters.
//
// Endpoint: GET /api/v1/branch/nearest-branch
func (c *BranchController) NearestBranch(ctx *gin.Context) {}

// NearestBranch retrieves branches near a specific location, filtered by radius or postal code, province, city, barangay, or region.
//
// - This returns all nearest and furthest branches based on the provided location parameters.
//
// Endpoint: GET /api/v1/branch/:company-id/nearest-company-branch/:branch-id
func (c *BranchController) NearestCompanyBranch(ctx *gin.Context) {}

// Endpoint: GET /api/v1/branch/current
func (c *CompanyController) CurrentBranch(ctx *gin.Context) {}
