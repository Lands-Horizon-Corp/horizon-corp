package routes

import (
	"horizon-server/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", userHandler.Register)
		v1.GET("/users/:id", userHandler.GetUser)
	}

	return router
}
