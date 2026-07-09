package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string

	GitHubToken string

	DatabaseURL string

	PollInterval time.Duration
}

func Load() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found. Using system environment variables.")
	}

	pollInterval, err := time.ParseDuration(getEnv("POLL_INTERVAL", "5m"))

	if err != nil {
		log.Fatal("Invalid POLL_INTERVAL")
	}

	cfg := &Config{

		Port: getEnv("PORT", "8080"),

		GitHubToken: getEnv("GITHUB_TOKEN", ""),

		DatabaseURL: getEnv("DATABASE_URL", ""),

		PollInterval: pollInterval,
	}

	validate(cfg)

	return cfg
}

func validate(cfg *Config) {

	if cfg.GitHubToken == "" {
		log.Fatal("GITHUB_TOKEN is missing")
	}

	if cfg.DatabaseURL == "" {
		log.Fatal("DATABASE_URL is missing")
	}
}

func getEnv(key string, fallback string) string {

	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}