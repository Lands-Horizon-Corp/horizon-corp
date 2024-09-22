package router

import (
	"horizon-core/config"
	"horizon-core/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProviedAPI(cfg *config.Config, db *gorm.DB) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// Register handlers
		handler.RegisterAdminRoutes(v1, db)
		handler.RegisterBranchRoutes(v1, db)
		handler.RegisterRoleRoutes(v1, db)
	}
	return router
}
