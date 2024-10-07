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
	feedbacksController *controllers.FeedbackController,

) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
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
		controllers.FeedbackRoutes(v1, feedbacksController)
	}
	return router
}
