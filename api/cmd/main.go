package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GeuberLucas/Gofre/api/internal/auth"
	"github.com/GeuberLucas/Gofre/api/pkg/config"
	"github.com/GeuberLucas/Gofre/api/pkg/db"
	gracefulshutdown "github.com/GeuberLucas/Gofre/api/pkg/graceful_shutdown"
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()
	config.LoadEnv()
	dbConn, err := db.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Connecting database: %v", err)
	}
	//Auth Module
	auth.SetupRoutes(app)

	var portApi string = ":50728"
	if os.Getenv("Enviroment") != "Development" {
		portApi = ":80"
	}

	shutdownManager := gracefulshutdown.NewGracefulShutdown(dbConn, app)

	go func() {
		if err := app.Listen(portApi); err != nil {
			fmt.Printf("Api stopped with error: %v\n", err)
		}
	}()

	// Fica bloqueado aguardando os sinais do sistema operacional
	shutdownManager.ListenSignals()

}
