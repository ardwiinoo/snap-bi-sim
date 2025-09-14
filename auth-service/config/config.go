package config

import "os"

type AppConfig struct {
	Port string

	// database
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string

	// cache
	RedisHost string
	RedisPort string
	RedisPass string
}

func LoadConfig() *AppConfig {

	cfg := &AppConfig{
		Port: getEnv("APP_PORT", "8080"),

		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "user"),
		DBPass:    getEnv("DB_PASS", "password"),
		DBName:    getEnv("DB_NAME", "snapbidev"),
		RedisHost: getEnv("REDIS_HOST", "localhost"),
		RedisPort: getEnv("REDIS_PORT", "6379"),
		RedisPass: getEnv("REDIS_PASS", ""),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
