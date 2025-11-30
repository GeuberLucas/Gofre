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
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/repository"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/router"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/service"
)

func main() {
	config.LoadEnv()
	dbConn, err := db.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Connecting database: %v", err)
	}

	messagingService := messaging.NewNATSMessaging()
	messagingService.ConnectToBroker()
	revenueRepository := repository.NewRevenueRepository(dbConn)
	expenseRepository := repository.NewExpenseRepository(dbConn)
	transactionService := service.NewTransactionService(revenueRepository, expenseRepository)

	routers := router.SetupRoutes(transactionService)
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
