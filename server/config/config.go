package config

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
)

type Config struct {
	DB      DBConfig    `yaml:"db"`
	App     AppConfig   `yaml:"app"`
	Log     LogConfig   `yaml:"log"`
	Api     cors.Config `yaml:"app"`
	Storage Storage     `yaml:"storage"`
}

type DBConfig struct {
	Username  string
	Password  string
	Host      string
	Port      string
	Name      string
	Charset   string
	ParseTime string
	Loc       string
}

type AppConfig struct {
	Port string
}

type LogConfig struct {
	Level string `yaml:"level"`
}

type Storage struct {
	Endpoint  string
	Region    string
	AccessKey string
	SecretKey string
}

func LoadConfig() (*Config, error) {

	config := &Config{
		App: AppConfig{
			Port: getEnv("SERVER_APP_PORT", "8080"),
		},
		DB: DBConfig{
			Username:  getEnv("SERVER_DB_USERNAME", ""),
			Password:  getEnv("SERVER_DB_PASSWORD", ""),
			Host:      getEnv("SERVER_DB_HOST", "db"),
			Port:      getEnv("SERVER_DB_PORT", "3306"),
			Name:      getEnv("SERVER_DB_NAME", ""),
			Charset:   getEnv("SERVER_DB_CHARSET", "utf8"),
			ParseTime: getEnv("SERVER_DB_PARSE_TIME", "True"),
			Loc:       getEnv("SERVER_DB_LOC", "Local"),
		},
		Log: LogConfig{
			Level: getEnv("LOG_LEVEL", "debug"),
		},
		Api: cors.Config{
			AllowOrigins: []string{
				"http://0.0.0.0",
				"http://0.0.0.0:8080",
				"http://0.0.0.0:3000",
				"http://0.0.0.0:3001",
				"http://0.0.0.0:80",
				"http://0.0.0.0:3000",
				"http://rea.development",
				"http://rea.pro",
				"http://localhost:8080",
				"http://localhost:3000",
				"http://localhost:3001",
				"http://localhost:3002",
			},
			AllowMethods:     []string{"POST", "GET", "PUT", "PATCH", "DELETE"},
			AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		},
		Storage: Storage{
			Endpoint:  getEnv("SERVER_MINIO_ENDPOINT", ""),
			Region:    getEnv("SERVER_MINIO_REGION", ""),
			AccessKey: getEnv("SERVER_MINIO_ACCESS_KEY", ""),
			SecretKey: getEnv("SERVER_MINIO_SECRET_KEY", ""),
		},
	}
	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}
