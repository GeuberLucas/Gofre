package router

import (
	"github.com/GeuberLucas/Gofre/backend/pkg/routes"
	"github.com/GeuberLucas/Gofre/backend/services/investments/internal/handler"
	"github.com/gorilla/mux"
)

func getTransactionRoutes(hd handler.IHandlerService) []routes.Route {
	return []routes.Route{
		{Path: "/", Method: "POST", HandlerFunc: hd.AddInvestmentHandler()},
		{Path: "/", Method: "GET", HandlerFunc: hd.GetInvestmentHandler()},
		{Path: "/{idInvestment}", Method: "GET", HandlerFunc: hd.GetByIdInvestmentHandler()},
		{Path: "/{idInvestment}", Method: "PUT", HandlerFunc: hd.UpdateInvestmentHandler()},
		{Path: "/{idInvestment}", Method: "DELETE", HandlerFunc: hd.DeleteInvestmentHandler()},
	}
}

// SetupRoutes initializes the router and defines the routes
func SetupRoutes(hd handler.IHandlerService) *mux.Router {
	r := mux.NewRouter()

	return routes.ConfigureRoutes(r, getTransactionRoutes(hd))
}
