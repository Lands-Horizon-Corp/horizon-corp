package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewEmployeeController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *EmployeeController {
	return &EmployeeController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

// GET :/
// Admin: Can get anything but must give company id
// Owner: All employee but must be only his company and its branches
// Employee: only employee on current branch
func (c *EmployeeController) Index(ctx *gin.Context) {

}

// GET :/id
// Admin: Can get anything but must give company id
// Owner: All employee but must be only his company and its branches
// Employee: only employee on current branch
func (c *EmployeeController) Show(ctx *gin.Context) {

}

// POST: /
// Admin: Can create employee but must assign company and branch
// Owner: can create but must assign branch
// Employee: not allowed
// Member: not allowed
func (c *EmployeeController) Store(ctx *gin.Context) {

}

// PUT: /
// Admin: Can create employee but must assign company and branch
// Owner: can create but must assign branch
// Employee: not allowed
// Member: not allowed
func (c *EmployeeController) Update(ctx *gin.Context) {

}

// DELETE: /:id
// Admin: allowed
// Owner: allowed
// Employee: not allowed
// Member: not allowed
func (c *EmployeeController) Destroy(ctx *gin.Context) {

}

func (c *EmployeeController) ForgotPassword(ctx *gin.Context) {

}
