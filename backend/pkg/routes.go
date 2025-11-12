package pkg

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Path        string
	Method      string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
	NeedsAuth   bool
}


func ConfigureRoutes(r *mux.Router, routes []Route) *mux.Router {
	for _, route := range routes {
		if route.NeedsAuth{
			r.HandleFunc(route.Path,middlewares.Authenticate)
		}
		r.HandleFunc(route.Path, route.HandlerFunc).Methods(route.Method)
	}
	return r
}