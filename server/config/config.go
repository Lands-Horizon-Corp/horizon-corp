package config

import (
	"os"
)

type AppConfig struct {
	AppPort            string
	AppForwardPort     string
	AppToken           string
	LogLevel           string
	DBUsername         string
	DBPassword         string
	DBHost             string
	DBPort             string
	DBName             string
	DBCharset          string
	DBParseTime        string
	DBLoc              string
	StorageEndpoint    string
	StorageRegion      string
	StorageAccessKey   string
	StorageSecretKey   string
	StorageBucketName  string
	StorageMaxFileSize string
}

func LoadConfig() (*AppConfig, error) {
	config := AppConfig{
		AppPort:            getEnv("APP_PORT", "8080"),         // default to 8080 if not set
		AppForwardPort:     getEnv("APP_FORWARD_PORT", "8081"), // default to 8081 if not set
		AppToken:           os.Getenv("APP_TOKEN"),
		LogLevel:           os.Getenv("LOG_LEVEL"),
		DBUsername:         os.Getenv("DB_USERNAME"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBName:             os.Getenv("DB_NAME"),
		DBCharset:          os.Getenv("DB_CHARSET"),
		DBParseTime:        os.Getenv("DB_PARSE_TIME"),
		DBLoc:              os.Getenv("DB_LOC"),
		StorageEndpoint:    os.Getenv("STORAGE_ENDPOINT"),
		StorageRegion:      os.Getenv("STORAGE_REGION"),
		StorageAccessKey:   os.Getenv("STORAGE_ACCESS_KEY"),
		StorageSecretKey:   os.Getenv("STORAGE_SECRET_KEY"),
		StorageBucketName:  os.Getenv("STORAGE_BUCKET_NAME"),
		StorageMaxFileSize: os.Getenv("STORAGE_MAX_FILE_SIZE"),
	}

	return &config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
