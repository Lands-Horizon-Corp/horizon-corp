package config

import (
	"github.com/spf13/viper"
)

// AppConfig holds the configuration settings for the application.
type AppConfig struct {
	AppPort            string `mapstructure:"APP_PORT"`
	AppForwardPort     string `mapstructure:"APP_FORWARD_PORT"`
	AppToken           string `mapstructure:"APP_TOKEN"`
	LogLevel           string `mapstructure:"LOG_LEVEL"`
	DBUsername         string `mapstructure:"DB_USERNAME"`
	DBPassword         string `mapstructure:"DB_PASSWORD"`
	DBHost             string `mapstructure:"DB_HOST"`
	DBPort             string `mapstructure:"DB_PORT"`
	DBName             string `mapstructure:"DB_NAME"`
	DBCharset          string `mapstructure:"DB_CHARSET"`
	DBParseTime        bool   `mapstructure:"DB_PARSE_TIME"`
	DBLoc              string `mapstructure:"DB_LOC"`
	StorageEndpoint    string `mapstructure:"STORAGE_ENDPOINT"`
	StorageRegion      string `mapstructure:"STORAGE_REGION"`
	StorageAccessKey   string `mapstructure:"STORAGE_ACCESS_KEY"`
	StorageSecretKey   string `mapstructure:"STORAGE_SECRET_KEY"`
	StorageBucketName  string `mapstructure:"STORAGE_BUCKET_NAME"`
	StorageMaxFileSize string `mapstructure:"STORAGE_MAX_FILE_SIZE"`
}

// LoadConfig loads the configuration from environment variables.
func LoadConfig() (*AppConfig, error) {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config AppConfig
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
