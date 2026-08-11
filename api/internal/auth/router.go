package auth

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router, hd IHandlerAuth) {
	route := app.Group("/auth")
	route.Post("/login", hd.LoginHandler)
	app.Post("/register", hd.RegisterHandler)
	app.Get("/isAuthenticated", hd.IsAuthenticatedHandler)
	app.Get("/profile/:userId", hd.ProfileHandler)
	app.Post("/fogort-passord", hd.ForgotPasswordHandler)
	app.Post("/reset-password/:HashEncoded", hd.ResetPasswordHandler)

}
