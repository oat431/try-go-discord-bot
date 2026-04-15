package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Loading ENV for discord bot")
}

func LoadEnvConfig() {
	failedToLoadEnv := godotenv.Load(".env")
	if failedToLoadEnv != nil {
		fmt.Println("Failed to load environment variables from .env file")
	}
}
