package reports

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app fiber.Router) {
	app.Get("/reports/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
