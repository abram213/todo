package db

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURI string
}

func InitConfig() (*Config, error) {
	config := &Config{
		DatabaseURI: viper.GetString("DatabaseURI"),
	}
	if config.DatabaseURI == "" {
		config.DatabaseURI = ":memory:"
	}
	return config, nil
}
