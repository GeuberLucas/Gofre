package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil && os.Getenv("Enviroment") == "Development" {
		log.Println("No .env file found")
	}
}
