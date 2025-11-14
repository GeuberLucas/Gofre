package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	err:= godotenv.Load()
	if err != nil {
        log.Println("No .env file found")
    }
	router := router.SetupRoutes()

	var portApi string = ":50728"
	if os.Getenv("Enviroment") != "Development"{
		portApi=":80"
	}

	log.Fatal(http.ListenAndServe(portApi, router))
}