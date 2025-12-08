package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/GeuberLucas/Gofre/backend/pkg/config"
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/nats-io/nats.go"
)

func main() {
	config.LoadEnv()
	messagingService, err := messaging.NewNATSMessaging()
	if err != nil {
		log.Fatal(err)
	}
	_, err = messagingService.SubscribeToSubject("finance.INVESTMENT.CREATE", HandlerInvestmentInsert)
	if err != nil {
		log.Println(err)
		return
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

}

func HandlerInvestmentInsert(m *nats.Msg) {
	err := m.Ack()

	if err != nil {
		log.Println("Unable to Ack", err)
		return
	}
	var msg messaging.MessagingDto
	err = json.Unmarshal(m.Data, &msg)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(m.Header)

}
