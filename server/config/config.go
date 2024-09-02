package config

import (
	"os"
)

type Config struct {
	DB  DBConfig  `yaml:"db"`
	App AppConfig `yaml:"app"`
	Log LogConfig `yaml:"log"`
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
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return defaultValue
}
