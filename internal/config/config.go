package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PostgresDSN string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found, using defaults")
	}

	return Config{
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
	}
}
