package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	fmt.Println("Heloo")
	c.JSON(http.StatusOK, gin.H{"message": "hello world"})
}
