package router

import (
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/GeuberLucas/Gofre/backend/pkg/routes"
	route "github.com/GeuberLucas/Gofre/backend/pkg/routes"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/handler"
	"github.com/gorilla/mux"
)

func getTransactionRoutes(broker *messaging.NATSMessaging) []routes.Route {
	return []routes.Route{
		{Path: "/login", Method: "POST", HandlerFunc: handler.LoginHandler(broker), NeedsAuth: false},
		{Path: "/register", Method: "POST", HandlerFunc: handler.RegisterHandler(broker), NeedsAuth: false},
		{Path: "/profile/{userId}", Method: "GET", HandlerFunc: handler.ProfileHandler(broker), NeedsAuth: true},
		{Path: "/fogort-passord", Method: "POST", HandlerFunc: handler.ForgotPasswordHandler(broker), NeedsAuth: false},
		{Path: "/reset-password/{HashEncoded}", Method: "POST", HandlerFunc: handler.ResetPasswordHandler(broker), NeedsAuth: false},
	}
}

// SetupRoutes initializes the router and defines the routes
func SetupRoutes(broker *messaging.NATSMessaging) *mux.Router {
	r := mux.NewRouter()

	return route.ConfigureRoutes(r, getTransactionRoutes(broker))
}
