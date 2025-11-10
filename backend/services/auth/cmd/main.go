package main

import (
	"log"
	"net/http"

	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/router"
)

func main() {
	router := router.SetupRoutes()

	log.Fatal(http.ListenAndServe(":8081", router))
}