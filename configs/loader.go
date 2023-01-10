package configs

import (
	"log"

	env "github.com/joho/godotenv"
)

func LoadEnv() {
	err := env.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
