package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type AppConfig struct {
	// Application
	AppName        string
	AppVersion     string
	AppEnv         string
	AppClientUrl   string
	AppPort        string
	AppSeeder      string
	AppTokenName   string
	AppToken       []byte
	AppForwardPort []byte
	AppLogo        string
	LogLevel       string

	// AWS
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	AWSRegion          string

	// Database
	DBUsername   string
	DBPassword   string
	DBHost       string
	DBPort       string
	DBName       string
	DBCharset    string
	DBParseTime  string
	DBLoc        string
	DBMaxRetries int
	DBRetryDelay time.Duration

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

	AdminEmail       string
	AdminPassword    string
	OwnerEmail       string
	OwnerPassword    string
	MemberEmail      string
	MemberPassword   string
	EmployeeEmail    string
	EmployeePassword string
}

func NewAppConfig() *AppConfig {
	var errList []string

	// Parse CACHE_PORT or default to 6379 if invalid
	cachePort, err := strconv.Atoi(getEnv("CACHE_PORT", "6379"))
	if err != nil {
		errList = append(errList, fmt.Sprintf("Invalid CACHE_PORT: %v (defaulting to 6379)", err))
		cachePort = 6379
	}

	// Construct Cache URL with fallback to redis:6379
	cacheURL := getEnv("CACHE_URL", "redis:"+strconv.Itoa(cachePort))

	// Required environment variables for database and storage
	requiredVars := []string{
		"DB_USERNAME", "DB_PASSWORD", "DB_HOST", "DB_PORT", "DB_NAME", "DB_CHARSET",
		"STORAGE_ENDPOINT", "STORAGE_REGION", "STORAGE_ACCESS_KEY", "STORAGE_SECRET_KEY", "STORAGE_BUCKET_NAME",
	}

	// Check for required variables
	for _, v := range requiredVars {
		if value := os.Getenv(v); value == "" {
			errList = append(errList, fmt.Sprintf("%s is required but not set", v))
		}
	}

	// Parse DB_MAX_RETRIES, defaulting to 5 if not set or invalid
	dbMaxRetries := 5
	if dbMaxRetriesStr := os.Getenv("DB_MAX_RETRIES"); dbMaxRetriesStr != "" {
		if val, err := strconv.Atoi(dbMaxRetriesStr); err == nil {
			dbMaxRetries = val
		} else {
			errList = append(errList, fmt.Sprintf("Invalid DB_MAX_RETRIES value '%s', defaulting to 5", dbMaxRetriesStr))
		}
	}

	// Parse DB_RETRY_DELAY as a time.Duration, defaulting to 2s if invalid
	dbRetryDelay := 2 * time.Second
	dbRetryDelayStr := getEnv("DB_RETRY_DELAY", "2s")
	if parsedDelay, err := time.ParseDuration(dbRetryDelayStr); err == nil {
		dbRetryDelay = parsedDelay
	} else {
		errList = append(errList, fmt.Sprintf("Invalid DB_RETRY_DELAY value '%s', defaulting to 2s", dbRetryDelayStr))
	}

	// If any errors were encountered, print them and return an empty AppConfig
	if len(errList) > 0 {
		for _, e := range errList {
			fmt.Println("Error:", e)
		}
		return &AppConfig{}
	}

	// Return the fully initialized configuration
	return &AppConfig{
		// Application
		AppName:      getEnv("APP_NAME", "horizon-corp"),
		AppVersion:   getEnv("APP_VERSION", "0.0.0"),
		AppEnv:       getEnv("APP_ENV", ""),
		AppClientUrl: getEnv("APP_CLIENT_URL", "http://client:80"),
		AppPort:      getEnv("APP_PORT", "8080"),
		AppSeeder:    getEnv("APP_SEEDER", "horizon-corp-seed"),
		AppTokenName: getEnv("APP_TOKEN_NAME", "horizon-corp"),
		AppToken:     []byte(os.Getenv("APP_TOKEN")),
		AppLogo:      getEnv("APP_LOGO", "https://s3.ap-southeast-2.amazonaws.com/horizon.assets/ecoop-logo.png"),
		LogLevel:     getEnv("LOG_LEVEL", "info"),

		// AWS
		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AWSRegion:          os.Getenv("AWS_REGION"),

		// Database
		DBUsername:   os.Getenv("DB_USERNAME"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBHost:       os.Getenv("DB_HOST"),
		DBPort:       os.Getenv("DB_PORT"),
		DBName:       os.Getenv("DB_NAME"),
		DBCharset:    os.Getenv("DB_CHARSET"),
		DBParseTime:  os.Getenv("DB_PARSE_TIME"),
		DBLoc:        os.Getenv("DB_LOC"),
		DBMaxRetries: dbMaxRetries,
		DBRetryDelay: dbRetryDelay,

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

		// Contact
		ContactNumber: os.Getenv("CONTACT_NUMBER"),

		// Accounts
		AdminEmail:       os.Getenv("ADMIN_EMAIL"),
		AdminPassword:    os.Getenv("ADMIN_PASSWORD"),
		OwnerEmail:       os.Getenv("OWNER_EMAIL"),
		OwnerPassword:    os.Getenv("OWNER_PASSWORD"),
		MemberEmail:      os.Getenv("MEMBER_EMAIL"),
		MemberPassword:   os.Getenv("MEMBER_PASSWORD"),
		EmployeeEmail:    os.Getenv("EMPLOYEE_EMAIL"),
		EmployeePassword: os.Getenv("EMPLOYEE_PASSWORD"),
	}
}

// getEnv returns the value of the environment variable named by the key.
// If the variable is not present, it returns defaultValue.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
