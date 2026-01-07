package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

	aggregatedRepository := repository.NewAggregatedRepository(dbConn)
	eventTrackRepository := repository.NewEventTrackRepository(dbConn)
	//expenseRepository := repository.NewExpensesRepository(dbConn)
	//investmentsRepository := repository.NewInvestmentsRepository(dbConn)
	//revenueRepository := repository.NewRevenueRepository(dbConn)
	aggregatedService := service.NewService(aggregatedRepository)
	//expenseService := service.NewExpenseService(expenseRepository)
	eventTrackService := service.NewEventTrackService(eventTrackRepository)
	consumerName := "Reports"
	messagingService.SubscribeToSubject(consumerName, "finance.>", func(msg *nats.Msg) {
		ctx := context.Background()
		eventID := msg.Header.Get(nats.MsgIdHdr)
		transac, _ := dbConn.Begin()
		if eventID == "" {
			log.Println("Aviso: Mensagem sem Nats-Msg-Id, descartando")
			msg.Term()
			return
		}

		exists, _ := eventTrackService.IsEventProcessed(ctx, eventID, consumerName)
		if exists {
			log.Printf("Evento %s já processado. Ignorando.", eventID)
			msg.Ack()
			return
		}
		var dto messaging.MessagingDto
		err = json.Unmarshal(msg.Data, &dto)
		if err != nil {
			errorMsg := fmt.Errorf("decode error mesage %s", eventID)
			log.Println(errorMsg)
			msg.Term()
			return
		}

		switch dto.Movement {
		case messaging.TypeExpense:
			err = aggregatedService.RegisterEventExpense(transac, dto)
		case messaging.TypeIncome:
			err = aggregatedService.RegisterEventRevenue(transac, dto)
		case messaging.TypeInvestment:
			err = aggregatedService.RegisterEventInvestment(transac, dto)
		}

		if err != nil {
			transac.Rollback()
			msg.NakWithDelay(5 * time.Second)
			errorMsg := fmt.Errorf("processing error mesage %s %v", eventID, err)
			log.Println(errorMsg)
			return
		}

		eventTrackService.MarkEventAsProcessed(ctx, eventID, dto.UserId, consumerName)
		transac.Commit()
		msg.Ack()
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
