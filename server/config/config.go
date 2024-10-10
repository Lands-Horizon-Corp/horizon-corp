package config

import (
	"os"
)

type AppConfig struct {
	AppPort            string
	AppToken           []byte
	AppForwardPort     []byte
	AppAdminToken      []byte
	AppOwnerToken      []byte
	AppEmployeeToken   []byte
	AppMemberToken     []byte
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
	appToken, err := Encrypt([]byte(os.Getenv("APP_TOKEN")), []byte(os.Getenv("APP_NAME")))
	if err != nil {
		return nil, err
	}
	appAdminToken, err := Encrypt([]byte(os.Getenv("APP_ADMIN_TOKEN")), appToken)
	if err != nil {
		return nil, err
	}
	appOwnerToken, err := Encrypt([]byte(os.Getenv("APP_OWNER_TOKEN")), appToken)
	if err != nil {
		return nil, err
	}
	appEmployeeToken, err := Encrypt([]byte(os.Getenv("APP_EMPLOYEE_TOKEN")), appToken)
	if err != nil {
		return nil, err
	}
	appMemberToken, err := Encrypt([]byte(os.Getenv("APP_MEMBER_TOKEN")), appToken)
	if err != nil {
		return nil, err
	}
	config := AppConfig{
		AppPort:            getEnv("APP_PORT", "8080"),
		AppToken:           appToken,
		AppAdminToken:      appAdminToken,
		AppOwnerToken:      appOwnerToken,
		AppEmployeeToken:   appEmployeeToken,
		AppMemberToken:     appMemberToken,
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
