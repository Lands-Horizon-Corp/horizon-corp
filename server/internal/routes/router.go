package routes

import (
	"horizon-server/cmd/server/docs"
	"horizon-server/config"
	"horizon-server/internal/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(cfg *config.Config, userHandler *handlers.UserHandler) *gin.Engine {

	docs.SwaggerInfo.BasePath = "/api"
	router := gin.Default()
	router.Use(cors.New(cfg.ApiConfig))

	router.GET("/", handlers.Index)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", userHandler.Register)
		v1.GET("/users/:id", userHandler.GetUser)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
