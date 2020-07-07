package app

import (
	"github.com/spf13/viper"
)

type Config struct {
	DataMode string
}

func InitConfig() (*Config, error) {
	config := &Config{
		DataMode: viper.GetString("DataMode"),
	}
	if config.DataMode == "" {
		config.DataMode = "sql"
	}
	return config, nil
}
