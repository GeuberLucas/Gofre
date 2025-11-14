package router

import (
	route "github.com/GeuberLucas/Gofre/backend/pkg/routes"
	"github.com/gorilla/mux"
)

// SetupRoutes initializes the router and defines the routes
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	return route.ConfigureRoutes(r, authRoutes)
}