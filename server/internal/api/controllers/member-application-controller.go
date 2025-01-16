package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type MemberApplicationController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewMemberApplicationController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *MemberApplicationController {
	return &MemberApplicationController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

func (c *MemberApplicationController) Index(ctx *gin.Context) {

}

func (c *MemberApplicationController) Show(ctx *gin.Context) {

}

func (c *MemberApplicationController) Store(ctx *gin.Context) {

}

func (c *MemberApplicationController) Update(ctx *gin.Context) {

}

func (c *MemberApplicationController) Destroy(ctx *gin.Context) {

}
