package controllers

import "github.com/gin-gonic/gin"

type AdminController struct{}

func NewAdminController() *AdminController {
	return &AdminController{}
}

// GET: /admin
// Retrieve admins with optional filtering for pagination or no pagination. Results can be converted to records.
//
//	Admin: Can retrieve all admins if admin  status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *AdminController) Index(ctx *gin.Context) {

}

// GET: /admin/:id
//
//	Admin: if admin  status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *AdminController) Show(ctx *gin.Context) {

}

// POST: /admin
//
//	Admin: Can create admin and automatically verified and also if admin and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *AdminController) Store(ctx *gin.Context) {

}

// PUT: /:id
//
//	Admin: Can change status of admin but only if admin and the status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *AdminController) Update(ctx *gin.Context) {

}

// DELETE: /:id
// Verifiy admin
//
//	Admin: only if admin  status is verified
//	Employee: not allowed
//	Owner: not allowed
//	Member: not allowed
func (c *AdminController) Destroy(ctx *gin.Context) {

}
