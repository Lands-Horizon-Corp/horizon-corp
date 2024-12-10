package providers

import (
	"errors"
	"fmt"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseService struct {
	Client *gorm.DB
	cfg    *config.AppConfig
	logger *LoggerService
}

func NewDatabaseProvider(
	cfg *config.AppConfig,
	logger *LoggerService,
) (*DatabaseService, error) {

	if cfg.DBHost == "" || cfg.DBName == "" || cfg.DBUsername == "" {
		logger.Error("Database configuration is incomplete",
			zap.String("DBHost", cfg.DBHost),
			zap.String("DBName", cfg.DBName),
			zap.String("DBUsername", cfg.DBUsername),
		)
		return nil, errors.New("incomplete database configuration")
	}

	dsn := buildDSN(cfg)
	logger.Info("Initializing database connection",
		zap.String("host", cfg.DBHost),
		zap.String("database", cfg.DBName),
		zap.String("user", cfg.DBUsername),
	)

	maxRetries := cfg.DBMaxRetries
	if maxRetries == 0 {
		maxRetries = 5
	}

	retryDelay := cfg.DBRetryDelay
	if retryDelay == 0 {
		retryDelay = 2 * time.Second
	}

	var db *gorm.DB
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			logger.Info("Successfully connected to the database",
				zap.String("host", cfg.DBHost),
				zap.String("database", cfg.DBName),
			)

			return &DatabaseService{
				Client: db,
				cfg:    cfg,
				logger: logger,
			}, nil
		}

		logger.Warn("Database connection attempt failed",
			zap.Int("attempt", attempt),
			zap.Int("maxRetries", maxRetries),
			zap.Error(err),
		)

		if attempt < maxRetries {
			logger.Info("Retrying database connection after delay",
				zap.Duration("retryDelay", retryDelay),
			)
			time.Sleep(retryDelay)
		}
	}

	logger.Error("Exceeded maximum number of retries to connect to the database",
		zap.Int("maxRetries", maxRetries),
		zap.Error(err),
	)
	return nil, fmt.Errorf("failed to connect to database after %d retries: %w", maxRetries, err)
}

func buildDSN(cfg *config.AppConfig) string {
	charset := cfg.DBCharset
	if charset == "" {
		charset = "utf8mb4"
	}

	parseTime := cfg.DBParseTime
	if parseTime == "" {
		parseTime = "True"
	}

	loc := cfg.DBLoc
	if loc == "" {
		loc = "Local"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		charset,
		parseTime,
		loc,
	)
}

func (ds *DatabaseService) Ping() error {
	sqlDB, err := ds.Client.DB()
	if err != nil {
		ds.logger.Error("Failed to retrieve SQL DB handle from GORM",
			zap.Error(err))
		return err
	}

	if err := sqlDB.Ping(); err != nil {
		ds.logger.Error("Database ping failed",
			zap.Error(err))
		return err
	}
	ds.logger.Debug("Database ping successful")
	return nil
}

func (ds *DatabaseService) Close() error {
	sqlDB, err := ds.Client.DB()
	if err != nil {
		ds.logger.Error("Failed to retrieve SQL DB handle from GORM",
			zap.Error(err))
		return err
	}
	err = sqlDB.Close()
	if err != nil {
		ds.logger.Error("Failed to close the database connection",
			zap.Error(err))
	} else {
		ds.logger.Info("Database connection closed successfully")
	}
	return err
}
