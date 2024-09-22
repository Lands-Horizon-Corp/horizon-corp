package database

import (
	"fmt"
	"horizon-core/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ProvideDatabase(cfg *config.Config) (*gorm.DB, error) {
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
func buildDSN(cfg *config.Config) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.Charset,
		cfg.DB.ParseTime,
		cfg.DB.Loc,
	)
}

// tryConnect attempts to open a connection to the database and returns the database object and error.
func tryConnect(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
