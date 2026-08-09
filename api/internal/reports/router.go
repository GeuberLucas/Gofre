package reports

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app fiber.Router) {
	route := app.Group("/expense")
	route.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

}
