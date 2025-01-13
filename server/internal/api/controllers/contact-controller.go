package controllers

import "github.com/gin-gonic/gin"

type ContactController struct{}

func NewContactController() *ContactController {
	return &ContactController{}
}

// GET: /api/v1/contact
// Retrieve contact with optional filtering for pagination or no pagination. Results can be converted to records.
//
//	Admin: Allowed
//	Employee: Not allowed
//	Owner: Not allowed
//	Member: Not allowed
func (c *ContactController) Index(ctx *gin.Context) {}

// GET: /api/v1/contact/:id
//
//	Admin: Allowed
//	Employee: Not allowed
//	Owner: Not allowed
//	Member: Not allowed
//
// Endpoint: GET /api/v1/contact/:id
func (c *ContactController) Show(ctx *gin.Context) {}

// PUBLIC
// Endpoint: POST /api/v1/contact
func (c *ContactController) Store(ctx *gin.Context) {}

//	Admin: Allowed
//	Employee: Not allowed
//	Owner: Not allowed
//	Member: Not allowed
//
// Endpoint: DELETE /api/v1/contact/:id
func (c *ContactController) Destroy(ctx *gin.Context) {}
