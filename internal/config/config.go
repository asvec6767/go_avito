package config

import (
	"os"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	Environment string `mapstructure:"ENVIRONMENT"`

	DatabaseURL string `mapstructure:"DATABASE_URL"`
	DBHost      string `mapstructure:"DB_HOST"`
	DBPort      string `mapstructure:"DB_PORT"`
	DBUser      string `mapstructure:"DB_USER"`
	DBPassword  string `mapstructure:"DB_PASSWORD"`
	DBName      string `mapstructure:"DB_NAME"`
	DBSSLMode   string `mapstructure:"DB_SSL_MODE"`

	LogLevel string `mapstructure:"LOG_LEVEL"`
}

func Load() *Config {
	// if os.Getenv("ENVIRONMENT") != "production" {
	// 	err := godotenv.Load()
	// 	if err != nil {
	// 		println("WARNING: .env не найден, используются системные переменные")
	// 	}
	// }

	return &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),

		DatabaseURL: getEnv("DATABASE_URL", ""),
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "5432"),
		DBUser:      getEnv("DB_USER", "postgres"),
		DBPassword:  getEnv("DB_PASSWORD", "password"),
		DBName:      getEnv("DB_NAME", "myapp"),
		DBSSLMode:   getEnv("DB_SSL_MODE", "disable"),

		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func (c *Config) GetDataBaseURL() string {
	if c.DatabaseURL != "" {
		return c.DatabaseURL
	}

	return "host=" + c.DBHost +
		" port=" + c.DBPort +
		" user=" + c.DBUser +
		" password=" + c.DBPassword +
		" dbname=" + c.DBName +
		" sslmode=" + c.DBSSLMode
}

func (c *Config) IsProduction() bool {
	return c.Environment == "production"
}
