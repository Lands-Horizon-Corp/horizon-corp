package routes

import (
	"horizon-server/internal/handlers"

	docs "horizon-server/cmd/server/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"

	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", userHandler.Register)
		v1.GET("/users/:id", userHandler.GetUser)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
