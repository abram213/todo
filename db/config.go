package db

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURI string
	PathToFile  string
}

func InitConfig() (*Config, error) {
	config := &Config{
		DatabaseURI: viper.GetString("DatabaseURI"),
		PathToFile:  viper.GetString("PathToFile"),
	}
	if config.DatabaseURI == "" {
		config.DatabaseURI = ":memory:"
	}
	if config.PathToFile == "" {
		config.PathToFile = "todos.txt"
	}
	return config, nil
}
