package config

import (
	"fmt"
	"horizon-core/helpers"
	"os"
	"sync"

	"github.com/gin-contrib/cors"
	"github.com/spf13/viper"
)

type Config struct {
	DB      DBConfig    `mapstructure:"db"`
	App     AppConfig   `mapstructure:"app"`
	Log     LogConfig   `mapstructure:"log"`
	Api     cors.Config `mapstructure:"api"`
	Storage Storage     `mapstructure:"storage"`
}

type DBConfig struct {
	Username  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	Host      string `mapstructure:"host"`
	Port      string `mapstructure:"port"`
	Name      string `mapstructure:"name"`
	Charset   string `mapstructure:"charset"`
	ParseTime string `mapstructure:"parse_time"`
	Loc       string `mapstructure:"loc"`
}

type AppConfig struct {
	Port  string `mapstructure:"port"`
	Token []byte `mapstructure:"token"`
}

type LogConfig struct {
	Level string `mapstructure:"level"`
}

type Storage struct {
	Endpoint    string `mapstructure:"endpoint"`
	Region      string `mapstructure:"region"`
	AccessKey   string `mapstructure:"access_key"`
	SecretKey   string `mapstructure:"secret_key"`
	BucketName  string `mapstructure:"bucket_name"`
	MaxfileSize int64  `mapstructure:"max_file_size"`
}

var (
	config *Config
	once   sync.Once
)

// GetConfig returns the singleton configuration instance
func GetConfig() *Config {
	once.Do(func() {
		config = &Config{}
		if err := loadConfig(); err != nil {
			panic(fmt.Errorf("fatal error loading config: %w", err))
		}
	})
	return config
}

func loadConfig() error {
	// Set default values
	viper.SetDefault("app.port", "8080")
	viper.SetDefault("log.level", "debug")
	viper.SetDefault("db.charset", "utf8")
	viper.SetDefault("db.parse_time", "True")
	viper.SetDefault("db.loc", "Local")
	viper.SetDefault("storage.max_file_size", helpers.FileSizeStringToInt64("10MB"))

	// Read configuration from environment variables
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP") // Use a prefix for environment variables
	viper.BindEnv("db.username")
	viper.BindEnv("db.password")
	viper.BindEnv("db.host")
	viper.BindEnv("db.port")
	viper.BindEnv("db.name")
	viper.BindEnv("app.token")
	viper.BindEnv("storage.access_key")
	viper.BindEnv("storage.secret_key")

	// Optionally read from .env file if it exists
	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err != nil {
			return fmt.Errorf("error reading .env file: %w", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return fmt.Errorf("unable to decode into struct: %w", err)
	}

	// Convert token to byte slice
	config.App.Token = []byte(viper.GetString("app.token"))

	return nil
}
