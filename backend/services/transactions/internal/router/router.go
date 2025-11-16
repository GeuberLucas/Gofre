package router

import (
	routes "github.com/GeuberLucas/Gofre/backend/pkg/routes"
)

var transactionRoutes []routes.Route{

}

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	return routes.ConfigureRoutes(r, transactionRoutes)
}