package database

import (
	"fmt"
	"horizon/server/config"
	"log" // Import log package for logging
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseService(cfg *config.AppConfig) (*gorm.DB, error) {
	dsn := buildDSN(cfg)
	maxRetries := 5
	retryDelay := 2 * time.Second

	var db *gorm.DB
	var err error
	for attempt := 1; attempt <= maxRetries; attempt++ {
		db, err = tryConnect(dsn)
		if err == nil {
			log.Println("Successfully connected to the database.")
			return db, nil
		}
		if attempt < maxRetries {
			time.Sleep(retryDelay)
		}
	}
	return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
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

func tryConnect(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
