package router

import (
	"github.com/GeuberLucas/Gofre/backend/pkg"
	"github.com/GeuberLucas/Gofre/backend/services/auth/internal/handler"
)


var routes = []pkg.Route{
	{Path: "/login", Method: "POST", HandlerFunc: handler.LoginHandler, NeedsAuth: false},
	{Path: "/register", Method: "POST", HandlerFunc: handler.RegisterHandler, NeedsAuth: false},
	{Path: "/profile", Method: "GET", HandlerFunc: handler.ProfileHandler, NeedsAuth: true},
	{Path: "/forgot-passord",Method: "POST", HandlerFunc: handler.ForgotPasswordHandler, NeedsAuth: false},
	{Path: "/reset-password/{HashEncoded}",Method: "POST", HandlerFunc: handler.ResetPasswordHandler, NeedsAuth: false},
}