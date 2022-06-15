package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// DatabaseConfig is the database configuration
type DatabaseConfig struct {
	DataSourceName  string
	Name            string
	MigrationFolder string
	MaxOpen         int
	MaxIdle         int
	MaxLifetime     int
}

// LoggerConfig is the logger configuration
type LoggerConfig struct {
	Development bool
	Encoding    string
	Level       string
}

//
type ServerConfig struct {
	AppVersion       string
	Mode             string
	RoutePrefix      string
	Debug            bool
	Port             string
	TimeoutSecs      int64
	ReadTimeoutSecs  int64
	WriteTimeoutSecs int64
}

type Config struct {
	ServerConfig ServerConfig
	DBConfig     DatabaseConfig
	LoggerConfig LoggerConfig
}

// LoadConfig loads the configuration from the given file.
func LoadConfig(file string) (*Config, error) {
	v := viper.New()

	v.SetConfigName(file)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	fmt.Fprintf(os.Stderr, "Loading configuration from %s\n", file)

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}

		return nil, err
	}

	var c Config
	err := v.Unmarshal(&c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
