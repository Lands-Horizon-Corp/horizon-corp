package controllers

import "github.com/gin-gonic/gin"

type BranchController struct{}

func NewBranchController() *BranchController {
	return &BranchController{}
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
// - Employee: Not allowed.
// - Member: Not allowed.
//
// Endpoint: POST /api/v1/branch
func (c *BranchController) Store(ctx *gin.Context) {}

// Update modifies the details of a specific branch by ID.
//
// - Admin: Allowed.
// - Owner: Allowed.
// - Employee: Not allowed.
// - Member: Not allowed.
//
// Endpoint: PUT /api/v1/branch/:id
func (c *BranchController) Update(ctx *gin.Context) {}

// Destroy deletes a specific branch by ID.
//
// - Admin: Allowed but only if there are no members in the branch.
// - Owner: Allowed but only if there are no members in the branch.
// - Employee: Not allowed.
// - Member: Not allowed.
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
// Endpoint: GET /api/v1/branch/nearest-company-branch/:id
func (c *BranchController) NearestCompanyBranch(ctx *gin.Context) {}
