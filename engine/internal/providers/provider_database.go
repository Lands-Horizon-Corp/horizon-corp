package providers

import (
	"fmt"
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseProvider(
	cfg *config.AppConfig,
	logger *zap.Logger,
) (*gorm.DB, error) {
	dsn := buildDSN(cfg)
	maxRetries := 5
	retryDelay := 2 * time.Second

	var db *gorm.DB
	var err error

	for attempt := 1; attempt <= maxRetries; attempt++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			logger.Info("Successfully connected to the database.")
			return db, nil
		}
		logger.Warn(fmt.Sprintf("Database connection attempt %d failed: %v", attempt, err))
		if attempt < maxRetries {
			time.Sleep(retryDelay)
		}
	}
	logger.Error("Exceeded maximum number of retries to connect to the database.", zap.Error(err))
	return nil, err
}

func buildDSN(cfg *config.AppConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBCharset,
		cfg.DBParseTime,
		cfg.DBLoc,
	)
}
