package controllers

import (
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/api/handlers"
	"github.com/Lands-Horizon-Corp/horizon-corp/internal/models"
	"github.com/gin-gonic/gin"
)

type QRScannerController struct {
	repository  *models.ModelRepository
	footstep    *handlers.FootstepHandler
	currentUser *handlers.CurrentUser
}

func NewQRScannerController(
	repository *models.ModelRepository,
	footstep *handlers.FootstepHandler,
	currentUser *handlers.CurrentUser,
) *QRScannerController {
	return &QRScannerController{
		repository:  repository,
		footstep:    footstep,
		currentUser: currentUser,
	}
}

// GET: /api/v1/qr/profile
// This QR will activate base on the JSON Data on the codes
// This will find a user - will give user id only
func (qc *QRScannerController) Profile(ctx *gin.Context) {

}

// GET: /api/v1/qr/find-profile
func (qc *QRScannerController) FindProfile(ctx *gin.Context) {

}
