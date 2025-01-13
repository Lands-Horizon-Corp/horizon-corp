package controllers

import "github.com/gin-gonic/gin"

type QRScannerController struct{}

func NewQRScannerController() *QRScannerController {
	return &QRScannerController{}
}

// GET: /api/v1/qr/profile
// This QR will activate base on the JSON Data on the codes
// This will find a user - will give user id only
func (qc *QRScannerController) Profile(ctx *gin.Context) {

}

// GET: /api/v1/qr/find-profile
func (qc *QRScannerController) FindProfile(ctx *gin.Context) {

}
