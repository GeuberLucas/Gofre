package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	fmt.Println(os.Getenv("DB_HOST"))
	if err != nil && os.Getenv("Enviroment") == "Development" {
		fmt.Println("No .env file found")
	}
}
