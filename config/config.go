package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

// LoadConfig loads environment variables
func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(" No .env file found, using system environment variables")
	}

	return Config{
		DatabaseURL: os.Getenv("DATABASE_URL"),
	}
}
