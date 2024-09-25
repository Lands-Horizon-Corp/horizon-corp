package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "hello world")
}

func RegisterIndexRoutes(router *gin.RouterGroup) {
	adminGroup := router.Group("/")
	{
		adminGroup.GET("", Index)
	}
}
