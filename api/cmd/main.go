package main

import (
	"fmt"
	"log"
	"os"

	"github.com/GeuberLucas/Gofre/api/internal/auth"
	"github.com/GeuberLucas/Gofre/api/internal/expense"
	"github.com/GeuberLucas/Gofre/api/internal/investments"
	"github.com/GeuberLucas/Gofre/api/internal/revenue"
	"github.com/GeuberLucas/Gofre/api/pkg/config"
	"github.com/GeuberLucas/Gofre/api/pkg/db"
	gracefulshutdown "github.com/GeuberLucas/Gofre/api/pkg/graceful_shutdown"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	configApp := fiber.Config{
		StrictRouting: true,
	}
	app := fiber.New(configApp)
	app.Use(logger.New())

	api := app.Group("/api")
	config.LoadEnv()
	dbConn, err := db.ConnectToDatabase()
	if err != nil {
		log.Fatalf("Connecting database: %v", err)
	}
	//Auth Module
	repoAuth := auth.NewAuthRepository(dbConn)
	serviceAuth := auth.NewAuthService(repoAuth)
	handlerAuth := auth.NewHandlerService(serviceAuth)
	auth.SetupRoutes(api, handlerAuth)
	protectedApi := api.Group("", handlerAuth.IsAuthenticatedMiddleware)

	//Investments module
	portRepo := investments.NewPortfolioRepository(dbConn)
	portSvc := investments.NewPortfolioService(portRepo)
	portHandler := investments.NewHandlerService(portSvc)
	investments.SetupRoutes(protectedApi, portHandler)

	//Expense module
	expRepo := expense.NewExpenseRepository(dbConn)
	expSvc := expense.NewExpenseService(expRepo)
	expHandler := expense.NewHandlerService(expSvc)
	expense.SetupRoutes(protectedApi, expHandler)

	//Investments module
	revRepo := revenue.NewRevenueRepository(dbConn)
	revSvc := revenue.NewRevenueService(revRepo)
	revHandler := revenue.NewHandlerService(revSvc)
	revenue.SetupRoutes(protectedApi, revHandler)

	var portApi string = ":50728"
	if os.Getenv("Enviroment") != "Development" {
		portApi = ":80"
	}

	shutdownManager := gracefulshutdown.NewGracefulShutdown(dbConn, app)

	go func() {
		if err := app.Listen(portApi); err != nil {
			fmt.Printf("Api stopped with error: %v\n", err)
		}
	}()

	// Fica bloqueado aguardando os sinais do sistema operacional
	shutdownManager.ListenSignals()

}
