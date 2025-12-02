package routes

import (
	"net/http"

	"github.com/GeuberLucas/Gofre/backend/pkg/middlewares"
	"github.com/gorilla/mux"
)

type Route struct {
	Path        string
	Method      string
	HandlerFunc http.HandlerFunc
}

func ConfigureRoutes(r *mux.Router, routes []Route) *mux.Router {
	for _, route := range routes {

		r.HandleFunc(route.Path, middlewares.Logger(route.HandlerFunc)).Methods(route.Method)

	}
	return r
}
