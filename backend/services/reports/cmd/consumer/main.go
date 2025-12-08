package main

import (
	"encoding/json"
	"log"

	gracefulshutdown "github.com/GeuberLucas/Gofre/backend/pkg/graceful_shutdown"
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/nats-io/nats.go"
)

func main() {
	messagingService, _ := messaging.NewNATSMessaging()

	messagingService.SubscribeToSubject("finance.INVESTMENT.CREATE", HandlerInvestmentInsert)

	gracefulshutdown.NewGracefulShutdown(nil, messagingService, nil)

}

func HandlerInvestmentInsert(m *nats.Msg) {
	var msg messaging.MessagingDto
	err := json.Unmarshal(m.Data, &msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(msg)

}
