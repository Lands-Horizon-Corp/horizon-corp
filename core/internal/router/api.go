package router

import (
	"horizon-core/config"
	"horizon-core/internal/handler"
	"horizon-core/internal/websocket"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func ProvideAPI(lc fx.Lifecycle, cfg *config.Config, db *gorm.DB, hub *websocket.Hub) {
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
	router.GET("/ws", websocket.ServeWs(hub))
}
