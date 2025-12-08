package main

import (
	"encoding/json"
	"log"

	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/nats-io/nats.go"
)

func main() {
	messagingService := messaging.NewNATSMessaging()

	messagingService.SubscribeToSubject("finance.INVESTMENT.CREATE", HandlerInvestmentInsert)
}

func HandlerInvestmentInsert(m *nats.Msg) {
	var msg messaging.MessagingDto
	err := json.Unmarshal(m.Data, &msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(msg)

}
