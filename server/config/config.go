package config

import (
	"fmt"
	"os"
	"strconv"
)

type AppConfig struct {
	// Application
	AppPort          string
	AppTokenName     string
	AppToken         []byte
	AppForwardPort   []byte
	AppAdminToken    []byte
	AppOwnerToken    []byte
	AppEmployeeToken []byte
	AppMemberToken   []byte
	LogLevel         string

	// AWS
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	AWSRegion          string

	// Database
	DBUsername  string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
	DBCharset   string
	DBParseTime string
	DBLoc       string

	// Caching
	CachePort     int
	CacheDB       int
	CacheURL      string
	CachePassword string

	// Storage
	StorageEndpoint    string
	StorageRegion      string
	StorageAccessKey   string
	StorageSecretKey   string
	StorageBucketName  string
	StorageMaxFileSize string

	// Email
	EmailHost     string
	EmailPort     string
	EmailUser     string
	EmailPassword string

	ContactNumber string
}

func LoadConfig() (*AppConfig, error) {
	var errList []string // Collect errors

	cachePort, err := strconv.Atoi(getEnv("CACHE_PORT", "6379"))
	if err != nil {
		errList = append(errList, fmt.Sprintf("Invalid CACHE_PORT: %v", err))
		cachePort = 6379
	}

	cacheURL := getEnv("CACHE_URL", "localhost:"+strconv.Itoa(cachePort))

	requiredVars := []string{"DB_USERNAME", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME", "DB_CHARSET", "STORAGE_ENDPOINT", "STORAGE_REGION", "STORAGE_ACCESS_KEY", "STORAGE_SECRET_KEY", "STORAGE_BUCKET_NAME"}
	for _, v := range requiredVars {
		if value := os.Getenv(v); value == "" {
			errList = append(errList, fmt.Sprintf("%s is required but not set", v))
		}
	}

	if len(errList) > 0 {
		return nil, fmt.Errorf("configuration errors: %v", errList)
	}

	config := AppConfig{
		// Application
		AppPort:          getEnv("APP_PORT", "8080"),
		AppTokenName:     getEnv("APP_TOKEN_NAME", "horizon-corp"),
		AppToken:         []byte(os.Getenv("APP_TOKEN")),
		AppAdminToken:    []byte(os.Getenv("APP_ADMIN_TOKEN")),
		AppOwnerToken:    []byte(os.Getenv("APP_OWNER_TOKEN")),
		AppEmployeeToken: []byte(os.Getenv("APP_EMPLOYEE_TOKEN")),
		AppMemberToken:   []byte(os.Getenv("APP_MEMBER_TOKEN")),
		LogLevel:         getEnv("LOG_LEVEL", "info"),

		// AWS
		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AWSRegion:          os.Getenv("AWS_REGION"),

		// Database
		DBUsername:  os.Getenv("DB_USERNAME"),
		DBPassword:  os.Getenv("DB_PASSWORD"),
		DBHost:      os.Getenv("DB_HOST"),
		DBPort:      os.Getenv("DB_PORT"),
		DBName:      os.Getenv("DB_NAME"),
		DBCharset:   os.Getenv("DB_CHARSET"),
		DBParseTime: os.Getenv("DB_PARSE_TIME"),
		DBLoc:       os.Getenv("DB_LOC"),

		// Cache
		CachePort:     cachePort,
		CacheDB:       0,
		CacheURL:      cacheURL,
		CachePassword: os.Getenv("CACHE_PASSWORD"),

		// Storage
		StorageEndpoint:    os.Getenv("STORAGE_ENDPOINT"),
		StorageRegion:      os.Getenv("STORAGE_REGION"),
		StorageAccessKey:   os.Getenv("STORAGE_ACCESS_KEY"),
		StorageSecretKey:   os.Getenv("STORAGE_SECRET_KEY"),
		StorageBucketName:  os.Getenv("STORAGE_BUCKET_NAME"),
		StorageMaxFileSize: os.Getenv("STORAGE_MAX_FILE_SIZE"),

		// Mail
		EmailHost:     os.Getenv("EMAIL_HOST"),
		EmailPort:     os.Getenv("EMAIL_PORT"),
		EmailUser:     os.Getenv("EMAIL_USER"),
		EmailPassword: os.Getenv("EMAIL_PASSWORD"),

		// SMS Test
		ContactNumber: os.Getenv("CONTACT_NUMBER"),
	}

	return &config, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
