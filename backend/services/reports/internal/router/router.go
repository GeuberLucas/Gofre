package router

import (
	"github.com/GeuberLucas/Gofre/backend/pkg/routes"
	"github.com/GeuberLucas/Gofre/backend/services/reports/internal/handler"
	"github.com/gorilla/mux"
)

func getTransactionRoutes(hd handler.IHandlerService) []routes.Route {
	return []routes.Route{}
}

// SetupRoutes initializes the router and defines the routes
func SetupRoutes(hd handler.IHandlerService) *mux.Router {
	r := mux.NewRouter()

	return routes.ConfigureRoutes(r, getTransactionRoutes(hd))
}
