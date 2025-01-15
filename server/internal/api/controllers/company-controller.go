package controllers

import "github.com/gin-gonic/gin"

type CompanyController struct{}

func NewCompanyController() *CompanyController {
	return &CompanyController{}
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
func (c *CompanyController) Index(ctx *gin.Context) {}

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
//	Employee: Not allowed
//	Member: Only own company
//
// Endpoint: POST /api/v1/company
func (c *CompanyController) Store(ctx *gin.Context) {}

// Update modifies the details of a specific company by ID.
//
// - Admin: Allowed.
// - Owner: Allowed but only his company
// - Employee: Not allowed.
// - Member: Not allowed.
// Endpoint: PUT /api/v1/company/:id
func (c *CompanyController) Update(ctx *gin.Context) {}

// DELETE: /:id
// Admin: Allowed but if no members on companyes, no companyes, and no employees on companyes.
// Owner: Only own company but if no members on companyes, no companyes, and no employees on companyes.
// Endpoint: DELETE /api/v1/company/:id
func (c *CompanyController) Destroy(ctx *gin.Context) {}

func (c *CompanyController) CurrentCompany(ctx *gin.Context) {}
