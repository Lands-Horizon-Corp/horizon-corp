package routes

import (
	"horizon/server/config"
	"horizon/server/internal/controllers"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func ProvideAPI(
	lc fx.Lifecycle,
	cfg *config.AppConfig,

	rolesController *controllers.RolesController,
	gendersController *controllers.GenderController,
	errorDetailsController *controllers.ErrorDetailsController,
	contactsController *controllers.ContactsController,

) *gin.Engine {
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

	router.GET("/", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	v1 := router.Group("/api/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		controllers.GenderRoutes(v1, gendersController)
		controllers.ErrorDetailsRoutes(v1, errorDetailsController)
		controllers.ContactsRoutes(v1, contactsController)
	}
	return router
}