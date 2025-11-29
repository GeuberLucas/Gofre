package main

import (
	"log"
	"net/http"
	"os"

	"github.com/GeuberLucas/Gofre/backend/pkg/config"
	"github.com/GeuberLucas/Gofre/backend/pkg/db"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/repository"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/router"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/service"
)

func main() {
	config.LoadEnv()
	dbConn, err := db.ConnectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()
	revenueRepository := repository.NewRevenueRepository(dbConn)
	expenseRepository := repository.NewExpenseRepository(dbConn)
	transactionService := service.NewTransactionService(revenueRepository, expenseRepository)
	router := router.SetupRoutes(transactionService)

	var portApi string = ":50728"
	if os.Getenv("Enviroment") != "Development" {
		portApi = ":80"
	}

	log.Fatal(http.ListenAndServe(portApi, router))
}
