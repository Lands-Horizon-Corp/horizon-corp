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
	configureCors(router, cfg)
	configureRoutes(router, userHandler, fileHandler)
	configureSwagger(router)
	return router
}

func configureSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func configureCors(router *gin.Engine, cfg *config.Config) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     cfg.Api.AllowOrigins,
		AllowMethods:     cfg.Api.AllowMethods,
		AllowHeaders:     cfg.Api.AllowHeaders,
		ExposeHeaders:    cfg.Api.ExposeHeaders,
		AllowCredentials: cfg.Api.AllowCredentials,
		MaxAge:           cfg.Api.MaxAge,
	}))
}

func configureRoutes(router *gin.Engine, userHandler *handlers.UserHandler, fileHandler *handlers.FileHandler) {
	v1 := router.Group("/api/v1")
	{
		// Home route
		router.GET("/", handlers.Index)

		// User routes
		v1.POST("/users", userHandler.Register)
		v1.GET("/users/:id", userHandler.GetUser)

		// File routes

		v1.POST("/files/upload", fileHandler.UploadFile)
		v1.DELETE("/files/:id", fileHandler.DeleteFile)
		v1.GET("/files/:id/download", fileHandler.DownloadFile)
		v1.GET("/files/:id", fileHandler.GetFile)

	}
}
