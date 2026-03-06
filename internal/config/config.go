package config

import (
	"errors"
	"os"
)

type Config struct {
	DATABASE_URL string
	PORT         string
}

func Load() (*Config, error) {
	if os.Getenv("DATABASE_URL") == "" {
		return nil, errors.New("DATABASE_URL environment variable is not set")
	}
	if os.Getenv("PORT") == "" {
		return nil, errors.New("PORT environment variable is not set")
	}
	return &Config{
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		PORT:         os.Getenv("PORT"),
	}, nil
}
