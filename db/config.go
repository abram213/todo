package db

import (
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURI string
	PathToFile  string
}

func InitConfig(dataMode string) (*Config, error) {
	config := &Config{
		DatabaseURI: viper.GetString("DatabaseURI"),
		PathToFile:  viper.GetString("PathToFile"),
	}
	if config.DatabaseURI == "" {
		switch dataMode {
		case "sql":
			config.DatabaseURI = ":memory:"
		case "mongo":
			config.DatabaseURI = "mongodb+srv://{your_username}:{your_password}@{cluster_name}.0fggh.mongodb.net/{table_name}?retryWrites=true&w=majority" //change to pass tests
		}
	}
	if config.PathToFile == "" {
		config.PathToFile = "todos.txt"
	}
	return config, nil
}
