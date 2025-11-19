package router

import (
	"github.com/GeuberLucas/Gofre/backend/pkg/routes"
	"github.com/gorilla/mux"
)

var transactionRoutes = []routes.Route{
	
}



func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	return routes.ConfigureRoutes(r, transactionRoutes)
}