package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	JWT JWTConfig
}

type JWTConfig struct {
	Secret string
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		config = &Config{
			JWT: JWTConfig{
				Secret: viper.GetString("jwt.secret"),
			},
		}
	}
	return config
}

func Init() error {
	viper.SetDefault("jwt.secret", "your-256-bit-secret")

	viper.AutomaticEnv()
	return nil
}
