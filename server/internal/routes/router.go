package routes

import (
	"horizon-server/cmd/server/docs"
	"horizon-server/internal/handlers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(userHandler *handlers.UserHandler) *gin.Engine {

	docs.SwaggerInfo.BasePath = "/api"
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://0.0.0.0",
			"http://0.0.0.0:8080",
			"http://0.0.0.0:3000",
			"http://0.0.0.0:3001",
			"http://0.0.0.0:80",
			"http://0.0.0.0:3000",
			"http://rea.development",
			"http://rea.pro",
			"http://ec2-54-146-249-91.compute-1.amazonaws.com",
		},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/", handlers.Index)
	v1 := router.Group("/api/v1")
	{
		v1.POST("/users", userHandler.Register)
		v1.GET("/users/:id", userHandler.GetUser)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
