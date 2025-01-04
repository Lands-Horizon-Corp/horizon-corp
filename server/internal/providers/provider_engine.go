package providers

import (
	"net/http"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/gin-gonic/gin"
)

type EngineService struct {
	Client *gin.Engine
	Server *http.Server
}

func NewEngineProvider(
	cfg *config.AppConfig,
	logger *LoggerService,
) (*EngineService, error) {
	router := gin.Default()
	server := &http.Server{
		Addr:         ":" + cfg.AppPort,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	return &EngineService{
		Client: router,
		Server: server,
	}, nil
}
