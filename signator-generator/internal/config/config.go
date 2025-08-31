package config

import "os"

type AppConfig struct {
	AppPort     	string
	PrivateKeyPath 	string
}

func LoadConfig() *AppConfig {

	cfg := &AppConfig{
		AppPort:      	getEnv("APP_PORT", "8080"),
		PrivateKeyPath: getEnv("PRIVATE_KEY_PATH", "./private.pem"),
	}

	return cfg
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
