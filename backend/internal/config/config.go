package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	Database string
}

func LoadEnv() *Config {
	// env vars might be set via Docker/K8s without a .env file.
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, fetching from system environment")
	}

	return &Config{
		Port:     getenv("PORT", "8000"),
		Database: getenv("DATABASE_URL", ""),
	}
}

func getenv(key, defaultval string) string {
	if value, exist := os.LookupEnv(key); exist {
		return value
	}
	return defaultval
}
