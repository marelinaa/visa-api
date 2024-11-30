package config

import (
	"os"
)

type AppConfig struct {
	APIPort        string
	AuthServiceURL string
	VisaServiceURL string
}

func Load() AppConfig {
	return AppConfig{
		APIPort:        getEnv("API_PORT", "8088"),
		AuthServiceURL: getEnv("AUTH_URL", "localhost:8082"),
		VisaServiceURL: getEnv("VISA_URL", "localhost:8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
