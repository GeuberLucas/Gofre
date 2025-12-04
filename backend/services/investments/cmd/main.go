package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GeuberLucas/Gofre/backend/pkg/config"
	"github.com/GeuberLucas/Gofre/backend/pkg/db"
	gracefulshutdown "github.com/GeuberLucas/Gofre/backend/pkg/graceful_shutdown"
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/handler"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/repository"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/router"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/service"
)

func main() {
	config.LoadEnv()
	dbConn, err := db.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Connecting database: %v", err)
	}

	messagingService := messaging.NewNATSMessaging()
	messagingService.ConnectToBroker()

	portfolioRepository := repository.NewPortfolioRepository(dbConn)
	portfolioService := service.NewPortfolioService(portfolioRepository, messagingService)
	handlerService := handler.NewHandlerService(portfolioService)

	routers := router.SetupRoutes(handlerService)
	var portApi string = ":50728"
	if os.Getenv("Enviroment") != "Development" {
		portApi = ":80"
	}
	serverConfig := http.Server{
		Addr:    portApi,
		Handler: routers,
	}

	shutdownManager := gracefulshutdown.NewGracefulShutdown(dbConn, messagingService, &serverConfig)

	go func() {
		if err := serverConfig.ListenAndServe(); err != nil {
			fmt.Printf("Api stopped with error: %v\n", err)
		}
	}()

	shutdownManager.ListenSignals()
}
