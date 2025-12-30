package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/GeuberLucas/Gofre/backend/pkg/config"
	"github.com/GeuberLucas/Gofre/backend/pkg/db"
	gracefulshutdown "github.com/GeuberLucas/Gofre/backend/pkg/graceful_shutdown"
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	service "github.com/GeuberLucas/Gofre/backend/services/reports/internal/Service"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/repository"
	"github.com/nats-io/nats.go"
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
	//eventTrackRepository := repository.NewEventTrackRepository(dbConn)
	expenseRepository := repository.NewExpensesRepository(dbConn)
	investmentsRepository := repository.NewInvestmentsRepository(dbConn)
	revenueRepository := repository.NewRevenueRepository(dbConn)
	aggregatedService := service.NewService(aggregatedRepository, revenueRepository, investmentsRepository)
	expenseService := service.NewExpenseService(expenseRepository)
	// // handlerService := handler.NewHandlerService()
	messagingService.SubscribeToSubject("finance.>", func(msg *nats.Msg) {
		var dto messaging.MessagingDto
		err := json.Unmarshal(msg.Data, &dto)
		if err != nil {
			log.Fatalln("decode error mesage")
		}

		switch dto.Movement {
		case messaging.TypeExpense:
			err = expenseService.RegisterEvent(dto)
			err = aggregatedService.RegisterEventExpense(dto)
		}

	})
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
