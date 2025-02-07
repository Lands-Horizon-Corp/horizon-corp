package providers

import (
	"time"

	"github.com/Lands-Horizon-Corp/horizon-corp/internal/config"
	"github.com/rotisserie/eris"
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
		return nil, eris.New("incomplete database configuration")
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
	return nil, eris.Wrapf(err, "failed to connect to database after %d retries", maxRetries)
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

	return cfg.DBUsername + ":" + cfg.DBPassword + "@tcp(" + cfg.DBHost + ":" + cfg.DBPort + ")/" + cfg.DBName + "?charset=" + charset + "&parseTime=" + parseTime + "&loc=" + loc
}

func (ds *DatabaseService) Ping() error {
	sqlDB, err := ds.Client.DB()
	if err != nil {
		ds.logger.Error("Failed to retrieve SQL DB handle from GORM",
			zap.Error(err))
		return eris.Wrap(err, "failed to retrieve SQL DB handle from GORM")
	}

	if err := sqlDB.Ping(); err != nil {
		ds.logger.Error("Failed to ping the database",
			zap.Error(err))
		return eris.Wrap(err, "failed to ping the database")
	}
	return nil
}

func (ds *DatabaseService) Close() error {
	sqlDB, err := ds.Client.DB()
	if err != nil {
		ds.logger.Error("Failed to retrieve SQL DB handle from GORM",
			zap.Error(err))
		return eris.Wrap(err, "failed to retrieve SQL DB handle from GORM")
	}
	err = sqlDB.Close()
	if err != nil {
		ds.logger.Error("Failed to close the database connection",
			zap.Error(err))
		return eris.Wrap(err, "failed to close the database connection")
	}
	ds.logger.Info("Database connection closed successfully")
	return nil
}
