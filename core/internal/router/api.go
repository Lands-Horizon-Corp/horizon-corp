package router

import (
	"horizon-core/config"
	"horizon-core/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ProvideAPI(cfg *config.Config, db *gorm.DB) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// Register handlers
		handler.RegisterAdminRoutes(v1, db)
		handler.RegisterBranchRoutes(v1, db)
		handler.RegisterCompanyRoutes(v1, db)
		handler.RegisterEmployeeRoutes(v1, db)
		handler.RegisterMediaRoutes(v1, db)
		handler.RegisterMemberRoutes(v1, db)
		handler.RegisterOwnerRoutes(v1, db)
		handler.RegisterPermissionRoutes(v1, db)
		handler.RegisterRoleRoutes(v1, db)
	}
	return router
}
