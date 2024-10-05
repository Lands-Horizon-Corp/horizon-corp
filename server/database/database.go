// package database

// import (
// 	"fmt"
// 	"horizon/server/config"

// 	"go.uber.org/zap"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type DB struct {
// 	*gorm.DB
// 	Logger *zap.Logger
// }

// func NewDB(cfg *config.AppConfig, logger *zap.Logger) (*DB, error) {
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
// 		cfg.DBUsername,
// 		cfg.DBPassword,
// 		cfg.DBHost,
// 		cfg.DBPort,
// 		cfg.DBName,
// 		cfg.DBCharset,
// 		cfg.DBParseTime,
// 		cfg.DBLoc,
// 	)

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	logger.Info("Database connection established")
// 	return &DB{db, logger}, nil
// }

// func (db *DB) Close() {
// 	if db.DB != nil {
// 		sqlDB, err := db.DB.DB()
// 		if err != nil {
// 			db.Logger.Fatal("Error getting database from GORM", zap.Error(err))
// 		}
// 		if err := sqlDB.Close(); err != nil {
// 			db.Logger.Fatal("Error closing the database", zap.Error(err))
// 		}
// 	}
// }

package database

import (
	"fmt"
	"horizon/server/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(cfg *config.AppConfig) (*gorm.DB, error) {
	dsn := buildDSN(cfg)

	maxRetries := 5
	retryDelay := 2 * time.Second

	var db *gorm.DB
	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		db, err = tryConnect(dsn)
		if err == nil {
			return db, nil
		}
		if attempt < maxRetries {
			time.Sleep(retryDelay)
		}
	}
	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
}

// buildDSN constructs the data source name (DSN) for the database connection.
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

// tryConnect attempts to open a connection to the database and returns the database object and error.
func tryConnect(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
