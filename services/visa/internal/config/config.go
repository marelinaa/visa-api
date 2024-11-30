package config

import (
	"os"
)

type AppConfig struct {
	DatabaseURL string
}

func Load() AppConfig {
	return AppConfig{
		DatabaseURL: getEnv("DB_URL", "postgres://postgres:p1111@localhost:5432/visa-api?sslmode=disable"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
