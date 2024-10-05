package database

import (
	"database/sql"
	"fmt"

	"horizon/server/config"

	"go.uber.org/zap"
)

type DB struct {
	*sql.DB
	Logger *zap.Logger
}

func NewDB(cfg *config.AppConfig, logger *zap.Logger) (*DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBCharset,
		cfg.DBParseTime,
		cfg.DBLoc,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	logger.Info("Database connection established")
	return &DB{db, logger}, nil
}

func (db *DB) Close() {
	if err := db.DB.Close(); err != nil {
		db.Logger.Fatal("Error closing the database", zap.Error(err))
	}
}
