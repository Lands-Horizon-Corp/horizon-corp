package controllers

import "github.com/gin-gonic/gin"

type ProfileController struct{}

func NewProfileController() *ProfileController {
	return &ProfileController{}
}

func (as ProfileController) ProfilePicture(ctx *gin.Context)             {}
func (as ProfileController) ProfileAccountSetting(ctx *gin.Context)      {}
func (as ProfileController) ProfileChangeEmail(ctx *gin.Context)         {}
func (as ProfileController) ProfileChangeContactNumber(ctx *gin.Context) {}
func (as ProfileController) ProfileChangePassword(ctx *gin.Context)      {}
func (as ProfileController) ProfileChangeUsername(ctx *gin.Context)      {}
