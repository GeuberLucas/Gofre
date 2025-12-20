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
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/repository"
)

func main() {
	config.LoadEnv()
	dbConn, err := db.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Connecting database: %v", err)
	}
	messagingService, err := messaging.NewNATSMessaging()
	if err != nil {
		log.Fatal(err)
	}
	//_, err = messagingService.SubscribeToSubject("finance.INVESTMENT.CREATE", HandlerInvestmentInsert)
	if err != nil {
		log.Println(err)
		return
	}
	aggregatedRepository := repository.NewAggregatedRepository(dbConn)
	eventTrackRepository := repository.NewEventTrackRepository(dbConn)
	expenseRepository := repository.NewExpensesRepository(dbConn)
	investmentsRepository := repository.NewInvestmentsRepository(dbConn)
	revenueRepository := repository.NewRevenueRepository(dbConn)
	Service := service.NewPortfolioService(portfolioRepository)
	// handlerService := handler.NewHandlerService()

	// routers := router.SetupRoutes(handlerService)
	var portApi string = ":50728"
	if os.Getenv("Enviroment") != "Development" {
		portApi = ":80"
	}
	serverConfig := http.Server{
		Addr: portApi,
		//Handler: routers,
	}

	shutdownManager := gracefulshutdown.NewGracefulShutdown(dbConn, messagingService, &serverConfig)

	go func() {
		if err := serverConfig.ListenAndServe(); err != nil {
			fmt.Printf("Api stopped with error: %v\n", err)
		}
	}()

	shutdownManager.ListenSignals()
}
