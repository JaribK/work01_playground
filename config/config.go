package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	JWT_SECRET string
}

func ReadInConfig() Config {
	return Config{
		JWT_SECRET: viper.GetString("JWT_SECRET"),
	}
}

func LoadConfig() error {
	viper.SetConfigFile("env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
