package config

import "os"

type Config struct {
	ServerAddr  string
	DatabaseURL string
}

func Load() *Config {
	return &Config{
		ServerAddr:  getEnv("SERVER_ADDR", ":8080"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@db:5432/filestore?sslmode=disable"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
