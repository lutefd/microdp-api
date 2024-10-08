package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBPath string
	Port   int
}

func Load() (*Config, error) {
	godotenv.Load()

	config := &Config{}

	config.DBPath = os.Getenv("DB_PATH")
	if config.DBPath == "" {
		config.DBPath = "test.db"
	}

	portStr := os.Getenv("PORT")
	if portStr != "" {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("invalid PORT: %w", err)
		}
		config.Port = port
	} else {
		config.Port = 8080
	}

	return config, nil
}
