package auth

import (
	"github.com/gofiber/fiber"
)

func SetupRoutes(app *fiber.App) {
	route := app.Group("/auth")
	route.Post("/login", LoginHandler)
	app.Post("/register", RegisterHandler)
	app.Get("/isAuthenticated", IsAuthenticatedHandler)
	app.Get("/profile/:userId", ProfileHandler)
	app.Post("/fogort-passord", ForgotPasswordHandler)
	app.Post("/reset-password/:HashEncoded", ResetPasswordHandler)

}
