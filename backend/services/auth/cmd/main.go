package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GeuberLucas/Gofre/backend/pkg/config"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/router"
)

func main() {
	config.LoadEnv()
	router := router.SetupRoutes()

	var portApi string = ":50728"
	if os.Getenv("Enviroment") != "Development" {
		portApi = ":80"
	}

	log.Fatal(http.ListenAndServe(portApi, router))
}
