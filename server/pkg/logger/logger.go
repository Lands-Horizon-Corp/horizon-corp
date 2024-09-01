package logger

import (
	"horizon-server/config"
	"log"
	"os"
)

func NewLogger(cfg *config.Config) *log.Logger {
	return log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
