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
			Username:  getEnv("DB_USERNAME", ""),
			Password:  getEnv("DB_PASSWORD", ""),
			Host:      getEnv("DB_HOST", "db"),
			Port:      getEnv("DB_PORT", "3306"),
			Name:      getEnv("DB_NAME", ""),
			Charset:   getEnv("DB_CHARSET", "utf8"),
			ParseTime: getEnv("DB_PARSE_TIME", "True"),
			Loc:       getEnv("DB_LOC", "Local"),
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
