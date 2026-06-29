package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort    string
	DBDriver   string
	DBPath     string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	JWTSecret  string
}

func Load() *Config {
	_ = godotenv.Load()
	return &Config{
		AppPort:    getEnv("APP_PORT", "8080"),
		DBDriver:   getEnv("DB_DRIVER", "postgres"),
		DBPath:     getEnv("DB_PATH", "database.db"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "telego"),
		DBPassword: getEnv("DB_PASSWORD", "telego1234"),
		DBName:     getEnv("DB_NAME", "telego"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		JWTSecret:  getEnv("JWT_SECRET", "default-secret-change-in-production"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvDuration(key string, fallback time.Duration) time.Duration {
	if value, ok := os.LookupEnv(key); ok {
		if d, err := time.ParseDuration(value); err == nil {
			return d
		}
	}
	return fallback
}
