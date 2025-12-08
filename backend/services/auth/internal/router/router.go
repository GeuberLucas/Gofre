package router

import (
	"github.com/GeuberLucas/Gofre/backend/pkg/messaging"
	"github.com/GeuberLucas/Gofre/backend/pkg/routes"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/handler"
	"github.com/gorilla/mux"
)

func getTransactionRoutes(broker messaging.IMessaging) []routes.Route {
	return []routes.Route{
		{Path: "/login", Method: "POST", HandlerFunc: handler.LoginHandler(broker)},
		{Path: "/register", Method: "POST", HandlerFunc: handler.RegisterHandler(broker)},
		{Path: "/isAuthenticated", Method: "GET", HandlerFunc: handler.IsAuthenticatedHandler()},
		{Path: "/profile/{userId}", Method: "GET", HandlerFunc: handler.ProfileHandler(broker)},
		{Path: "/fogort-passord", Method: "POST", HandlerFunc: handler.ForgotPasswordHandler(broker)},
		{Path: "/reset-password/{HashEncoded}", Method: "POST", HandlerFunc: handler.ResetPasswordHandler(broker)},
	}
}

// SetupRoutes initializes the router and defines the routes
func SetupRoutes(broker messaging.IMessaging) *mux.Router {
	r := mux.NewRouter()

	return routes.ConfigureRoutes(r, getTransactionRoutes(broker))
}
