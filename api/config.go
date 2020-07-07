package api

import "github.com/spf13/viper"

type Config struct {
	Port int
}

func InitConfig() (*Config, error) {
	config := &Config{
		Port: viper.GetInt("Port"),
	}
	if config.Port == 0 {
		config.Port = 4545
	}
	return config, nil
}
