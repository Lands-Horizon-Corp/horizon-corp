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

	authController *controllers.AuthController,
	roleController *controllers.RolesController,
	genderController *controllers.GenderController,
	errorDetailController *controllers.ErrorDetailController,
	contactController *controllers.ContactsController,
	feedbackController *controllers.FeedbackController,
	mediaController *controllers.MediaController,

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
			"http://localhost:80",
			"http://localhost:8080",
			"http://localhost:3000",
			"http://localhost:3001",
			"http://localhost:3002",
		},
		AllowMethods:  []string{"POST", "GET", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},

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
		controllers.AuthRoutes(v1, authController)
		controllers.GenderRoutes(v1, genderController)
		controllers.ErrorDetailRoutes(v1, errorDetailController)
		controllers.ContactsRoutes(v1, contactController)
		controllers.FeedbackRoutes(v1, feedbackController)
		controllers.MediaRoutes(v1, mediaController)
	}
	return router
}
