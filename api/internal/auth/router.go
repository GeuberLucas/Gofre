package auth

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app fiber.Router, hd IHandlerAuth) {
	route := app.Group("/auth")
	route.Post("/login", hd.LoginHandler)
	route.Post("/register", hd.RegisterHandler)
	route.Get("/profile/:userId", hd.ProfileHandler)
	route.Post("/fogort-passord", hd.ForgotPasswordHandler)
	route.Post("/reset-password/:HashEncoded", hd.ResetPasswordHandler)

}
