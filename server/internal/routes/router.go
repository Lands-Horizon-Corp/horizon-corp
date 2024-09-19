package routes

import (
	"horizon-server/cmd/server/docs"
	"horizon-server/config"
	"horizon-server/internal/handlers"
	"horizon-server/internal/models"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/o1egl/paseto"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(cfg *config.Config, userHandler *handlers.UserHandler, fileHandler *handlers.FileHandler) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api"
	router := gin.Default()
	configureCors(router, cfg)
	configureRoutes(
		router,
		cfg,
		userHandler,
		fileHandler,
	)
	configureSwagger(router)
	return router
}

func configureSwagger(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

var v2 = paseto.NewV2()

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
func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header must be Bearer token"})
			c.Abort()
			return
		}

		var payload models.User
		err := v2.Decrypt(tokenParts[1], cfg.App.Token, &payload, nil)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Optionally, check if the user ID from the payload exists in the database here

		c.Set("userID", payload.ID)
		c.Next()
	}
}

func configureRoutes(router *gin.Engine, cfg *config.Config, userHandler *handlers.UserHandler, fileHandler *handlers.FileHandler) {
	v1 := router.Group("/api/v1")
	{
		router.GET("/", handlers.Index)

		// Public endpoints
		v1.POST("/user/register", userHandler.RegisterUser)
		v1.POST("/user/login", userHandler.LoginUser)

		// Protected endpoints
		protected := router.Group("/user")
		protected.Use(AuthMiddleware(cfg))
		{
			protected.POST("/logout", userHandler.LogoutUser)
			protected.GET("/profile", userHandler.GetUserProfile)
			protected.POST("/change-password", userHandler.ChangeUserPassword)
			protected.POST("/update-profile", userHandler.UpdateUserProfile)
			protected.POST("/users", userHandler.ListUsers)
		}

		// File routes
		v1.POST("/files/upload", fileHandler.UploadFile)
		v1.DELETE("/files/:id", fileHandler.DeleteFile)
		v1.GET("/files/:id/download", fileHandler.DownloadFile)
		v1.GET("/files/:id", fileHandler.GetFile)
	}
}
