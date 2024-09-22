package config

import (
	"go.uber.org/zap"
)

// NewLogger constructs a new zap.Logger
func ProvideLogger() (*zap.Logger, error) {
	// Use zap's production configuration
	cfg := zap.NewProductionConfig()

	// Customize the configuration if needed
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	cfg.Encoding = "json" // or "console"

	// Build the logger
	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}
