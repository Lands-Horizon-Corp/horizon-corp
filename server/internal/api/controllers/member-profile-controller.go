package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type MemberProfileController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewMemberProfileController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *MemberProfileController {
	return &MemberProfileController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

func (c *MemberProfileController) Index(ctx *gin.Context) {

}

func (c *MemberProfileController) Show(ctx *gin.Context) {

}

func (c *MemberProfileController) Store(ctx *gin.Context) {

}

func (c *MemberProfileController) Update(ctx *gin.Context) {

}

func (c *MemberProfileController) Destroy(ctx *gin.Context) {

}
