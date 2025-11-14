package main

import (
	"log"
	"net/http"

	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/router"
	"github.com/joho/godotenv"
)

func main() {
	err:= godotenv.Load()
	if err != nil {
        log.Println("No .env file found")
    }
	router := router.SetupRoutes()

	log.Fatal(http.ListenAndServe(":80", router))
}