package auth

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app fiber.Router) {
	route := app.Group("/auth")
	route.Post("/login", LoginHandler)
	app.Post("/register", RegisterHandler)
	app.Get("/isAuthenticated", IsAuthenticatedHandler)
	app.Get("/profile/:userId", ProfileHandler)
	app.Post("/fogort-passord", ForgotPasswordHandler)
	app.Post("/reset-password/:HashEncoded", ResetPasswordHandler)

}
