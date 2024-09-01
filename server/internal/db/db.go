package db

import (
	"fmt"
	"horizon-server/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.Charset,
		cfg.DB.ParseTime,
		cfg.DB.Loc,
	)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
