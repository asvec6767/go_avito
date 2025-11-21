package config

import "github.com/spf13/viper"

type Config struct {
	Port        string `mapstructure:"PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	LogLevel    string `mapstructure:"LOG_LEVEL"`
}

func Load() *Config {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("LOG_LEVEL", "info")

	viper.AutomaticEnv()

	var config Config
	viper.Unmarshal(&config)

	return &config
}
