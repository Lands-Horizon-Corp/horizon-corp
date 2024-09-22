package router

import (
	"horizon-core/config"
	"horizon-core/internal/handler"
	"horizon-core/internal/websocket"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func ProvideAPI(lc fx.Lifecycle, cfg *config.Config, db *gorm.DB, hub *websocket.Hub) {
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
			"http://localhost:8080",
			"http://localhost:3000",
			"http://localhost:3001",
			"http://localhost:3002",
		},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

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
