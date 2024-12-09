package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	JWT_SECRET       string
	PUBLIC_KEY       string
	PRIVATE_KEY      string
	MINIO_ACCESS_KEY string
	MINIO_SECRET_KEY string
}

func ReadInConfig() Config {
	return Config{
		JWT_SECRET:       viper.GetString("JWT_SECRET"),
		PUBLIC_KEY:       viper.GetString("PUBLIC_KEY_PATH"),
		PRIVATE_KEY:      viper.GetString("PRIVATE_KEY_PATH"),
		MINIO_ACCESS_KEY: viper.GetString("MINIO_ACCESS_KEY"),
		MINIO_SECRET_KEY: viper.GetString("MINIO_SECRET_KEY"),
	}
}

func LoadConfig() error {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
