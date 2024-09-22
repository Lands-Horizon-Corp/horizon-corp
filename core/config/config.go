package config

import (
	"horizon-core/helpers"

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

func ProvideConfig() (*Config, error) {
	// Set default values
	viper.SetDefault("app.port", "8080")
	viper.SetDefault("log.level", "debug")
	viper.SetDefault("db.charset", "utf8")
	viper.SetDefault("db.parse_time", "True")
	viper.SetDefault("db.loc", "Local")
	viper.SetDefault("storage.max_file_size", helpers.FileSizeStringToInt64("10MB"))

	// Read .env file
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	// Convert token to byte slice
	config.App.Token = []byte(viper.GetString("app.token"))

	return config, nil
}
