package routes

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (ar *APIRoutes) APITestRoute() {
	ar.router.GET("/favicon.ico", func(c *gin.Context) {
		resp, err := http.Get("https://s3.ap-southeast-2.amazonaws.com/horizon.assets/ecoop-logo.png")
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		c.Header("Content-Type", resp.Header.Get("Content-Type"))
		_, err = io.Copy(c.Writer, resp.Body)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return
		}
	})
	ar.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to the secure Gin server!"})
	})
	ar.router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
