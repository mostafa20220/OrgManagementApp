// config/config.go
package config

import (
	"github.com/spf13/viper"
)

// AppConfig represents the application configuration structure.
type AppConfig struct {
	Server   ServerConfig
	Database DatabaseConfig
}

// ServerConfig represents the server configuration structure.
type ServerConfig struct {
	Port int
}

// DatabaseConfig represents the database configuration structure.
type DatabaseConfig struct {
	MongoURI string
}

// LoadConfig loads the application configuration from the provided YAML file.
func LoadConfig() *AppConfig {
	viper.SetConfigFile("./app-config.yaml")
	// viper.SetConfigType("./database-config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config AppConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
