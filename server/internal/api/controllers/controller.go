package controllers

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

// Endpoint: GET /
// Description: Returns a welcome message for the ECOOP: Horizon server.
func (c *Controller) Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to the ECOOP: Horizon server"})
}

// Endpoint: GET /favico
// Description: Serves the favicon image by proxying the request to an S3 bucket.
func (c *Controller) Favico(ctx *gin.Context) {
	resp, err := http.Get("https://s3.ap-southeast-2.amazonaws.com/horizon.assets/ecoop-logo.png")
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	ctx.Header("Content-Type", resp.Header.Get("Content-Type"))
	_, err = io.Copy(ctx.Writer, resp.Body)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		return
	}
}

// Endpoint: GET /ping
// Description: A health check endpoint that returns a "pong" response to indicate the server is running.
func (c *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
