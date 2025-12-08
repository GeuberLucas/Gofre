package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/GeuberLucas/Gofre/backend/pkg/config"
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/router"
)

func main() {
	config.LoadEnv()
	messagingService, _ := messaging.NewNATSMessaging()

	router := router.SetupRoutes(messagingService)

	var portApi string = ":50728"
	if os.Getenv("Enviroment") != "Development" {
		portApi = ":80"
	}
	serverConfig := http.Server{
		Addr:    portApi,
		Handler: router,
	}
	if err := serverConfig.ListenAndServe(); err != nil {
		fmt.Printf("Api stopped with error: %v\n", err)
	}

}
