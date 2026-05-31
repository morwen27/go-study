package config

import (
	"os"

	"github.com/lpernett/godotenv"
)

type Config struct {
	key string
}

func InitConfig() *Config {
	envLoadError := godotenv.Load()
	if envLoadError != nil {
		panic("Environment variables were not loaded")
	}

	key := os.Getenv("KEY")
	if key == "" {
		panic("There is not KEY")
	}

	return &Config{
		key,
	}
}
