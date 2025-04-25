package config

import "os"

type Config struct {
	ServerPort  string
	PostgresDSN string
	Environment string
}

func Load() *Config {
	return &Config{
		ServerPort:  getEnvOrDefault("SERVER_PORT", "8080"),
		PostgresDSN: getEnvOrDefault("PostgresDSN", "host=localhost user=postgres password=postgres dbname=url_shortener port=5432 sslmode=disable"),
		Environment: getEnvOrDefault("Env", "development"),
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
