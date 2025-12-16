package models

import (
	"os"
)

type Config struct {
	Email    string
	Password string
	Headless bool
}

func LoadConfig() *Config {
	return &Config{
		Email:    os.Getenv("LINKEDIN_EMAIL"),
		Password: os.Getenv("LINKEDIN_PASSWORD"),
		Headless: os.Getenv("HEADLESS") == "true",
	}
}
