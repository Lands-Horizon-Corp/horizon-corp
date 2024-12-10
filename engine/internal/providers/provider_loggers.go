package providers

import (
	"fmt"
	"strings"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerService struct {
	Client *zap.Logger
	cfg    *config.AppConfig
}

func NewLoggerProvider(cfg *config.AppConfig) (*LoggerService, error) {
	logLevel := mapLogLevel(cfg.LogLevel)

	var (
		logger *zap.Logger
		err    error
	)

	switch strings.ToLower(cfg.AppEnv) {
	case "production":
		prodConfig := zap.NewProductionConfig()
		prodConfig.Encoding = "json"
		prodConfig.Level = zap.NewAtomicLevelAt(logLevel)
		prodConfig.Sampling = &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		}
		logger, err = prodConfig.Build()
		if err != nil {
			return nil, fmt.Errorf("failed to create production logger: %w", err)
		}

	default:
		devConfig := zap.NewDevelopmentConfig()
		devConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		devConfig.Level = zap.NewAtomicLevelAt(logLevel)
		logger, err = devConfig.Build()
		if err != nil {
			return nil, fmt.Errorf("failed to create development logger: %w", err)
		}
	}

	// Add contextual fields if available
	if cfg.AppName != "" {
		logger = logger.With(zap.String("appName", cfg.AppName))
	}
	if cfg.AppEnv != "" {
		logger = logger.With(zap.String("environment", cfg.AppEnv))
	}
	if cfg.AppVersion != "" {
		logger = logger.With(zap.String("version", cfg.AppVersion))
	}

	logger.Info("Logger initialized successfully",
		zap.String("environment", cfg.AppEnv),
		zap.String("appName", cfg.AppName),
		zap.String("appVersion", cfg.AppVersion),
		zap.String("logLevel", logLevel.String()),
	)

	return &LoggerService{
		Client: logger,
		cfg:    cfg,
	}, nil
}

func mapLogLevel(levelStr string) zapcore.Level {
	switch strings.ToLower(levelStr) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn", "warning":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	default:
		return zap.InfoLevel
	}
}

func (ls *LoggerService) Info(msg string, fields ...zap.Field) {
	ls.Client.Info(msg, fields...)
}

func (ls *LoggerService) Error(msg string, fields ...zap.Field) {
	ls.Client.Error(msg, fields...)
}

func (ls *LoggerService) Warn(msg string, fields ...zap.Field) {
	ls.Client.Warn(msg, fields...)
}

func (ls *LoggerService) Debug(msg string, fields ...zap.Field) {
	ls.Client.Debug(msg, fields...)
}

func (ls *LoggerService) Sync() error {
	return ls.Client.Sync()
}
