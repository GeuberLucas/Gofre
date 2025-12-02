package router

import (
	"github.com/GeuberLucas/Gofre/backend/pkg/routes"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/handler"
	"github.com/GeuberLucas/Gofre/backend/services/transaction/internal/service"
	"github.com/gorilla/mux"
)

func getTransactionRoutes(s *service.TransactionService) []routes.Route {
	return []routes.Route{
		//expense paths
		{Path: "/expense", Method: "POST", HandlerFunc: handler.AddExpenseHandler(s)},
		{Path: "/expense/{idTransaction}", Method: "GET", HandlerFunc: handler.GetByIdExpenseHandler(s)},
		{Path: "/expense", Method: "GET", HandlerFunc: handler.GetByIdUserExpenseHandler(s)},
		{Path: "/expense/{idTransaction}", Method: "PUT", HandlerFunc: handler.UpdateExpenseHandler(s)},
		{Path: "/expense/{idTransaction}", Method: "DELETE", HandlerFunc: handler.DeleteExpenseHandler(s)},

		//revenue paths
		{Path: "/revenue", Method: "POST", HandlerFunc: handler.AddRevenueHandler(s)},
		{Path: "/revenue/{idTransaction}", Method: "GET", HandlerFunc: handler.GetByIdRevenueHandler(s)},
		{Path: "/revenue", Method: "GET", HandlerFunc: handler.GetByIdUserRevenueHandler(s)},
		{Path: "/revenue/{idTransaction}", Method: "PUT", HandlerFunc: handler.UpdateRevenueHandler(s)},
		{Path: "/revenue/{idTransaction}", Method: "DELETE", HandlerFunc: handler.DeleteRevenueHandler(s)},
	}
}

func SetupRoutes(s *service.TransactionService) *mux.Router {
	r := mux.NewRouter()
	transactionRoutes := getTransactionRoutes(s)
	return routes.ConfigureRoutes(r, transactionRoutes)
}
