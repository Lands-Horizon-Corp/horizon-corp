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

func SetupRouter(cfg *config.Config, userHandler *handlers.UserHandler, fileHandler *handlers.FileHandler) *gin.Engine {

	docs.SwaggerInfo.BasePath = "/api"
	router := gin.Default()
	router.Use(cors.New(cfg.Api))

	router.GET("/", handlers.Index)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", userHandler.Register)
		v1.GET("/users/:id", userHandler.GetUser)

		// File handler routes
		v1.POST("/upload", fileHandler.UploadFile)
		v1.DELETE("/delete/:key", fileHandler.DeleteFile)
		v1.GET("/presigned-url/:key", fileHandler.GeneratePresignedURL)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
