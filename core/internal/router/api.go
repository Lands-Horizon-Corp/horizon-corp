package router

import (
	"horizon-core/config"

	"github.com/gin-gonic/gin"
)

func API(router *gin.Engine, cfg *config.Config) {
	// v1 := router.Group("/api/v1")
	// {
	// 	router.GET("/", handlers.Index)

	// 	// Public endpoints
	// 	v1.POST("/user/register", userHandler.RegisterUser)
	// 	v1.POST("/user/login", userHandler.LoginUser)

	// 	// Protected endpoints
	// 	protected := router.Group("/user")
	// 	protected.Use(AuthMiddleware(cfg))
	// 	{
	// 		protected.POST("/logout", userHandler.LogoutUser)
	// 		protected.GET("/profile", userHandler.GetUserProfile)
	// 		protected.POST("/change-password", userHandler.ChangeUserPassword)
	// 		protected.POST("/update-profile", userHandler.UpdateUserProfile)
	// 		protected.POST("/users", userHandler.ListUsers)
	// 	}

	// 	// File routes
	// 	v1.POST("/files/upload", fileHandler.UploadFile)
	// 	v1.DELETE("/files/:id", fileHandler.DeleteFile)
	// 	v1.GET("/files/:id/download", fileHandler.DownloadFile)
	// 	v1.GET("/files/:id", fileHandler.GetFile)
	// }
}
