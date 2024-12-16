package config

import (
	"fmt"
	"os"

	"github.com/marelinaa/visa-api/services/visa/internal/domain"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DatabaseURL string
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
	SSLMode  string
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

type Config struct {
	API   APIConfig
	DB    DBConfig
	Redis RedisConfig
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

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, domain.ErrLoadEnvVars
	}

	requiredEnvs := []string{
		"DB_HOST", "DB_USER", "DB_PASSWORD",
		"DB_NAME", "DB_PORT", "DB_SSL", "API_PORT",
		"REDIS_HOST", "REDIS_PORT",
	}

	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			return nil, fmt.Errorf("environment variable `%s` is not set or is empty", env)
		}
	}

	config := Config{
		API: APIConfig{
			Port: os.Getenv("API_PORT"),
		},
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
			SSLMode:  os.Getenv("DB_SSL"),
		},
		Redis: RedisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: "",
			DB:       0,
		},
	}

	return &config, nil
}
