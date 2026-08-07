package investments

import "github.com/gofiber/fiber/v2"

// SetupRoutes initializes the router and defines the routes
func SetupRoutes(app fiber.Router, hd IHandlerService) {
	route := app.Group("/investments")
	route.Get("/", hd.GetInvestmentHandler)
	app.Post("/", hd.AddInvestmentHandler)
	app.Get("/:idInvestment", hd.GetByIdInvestmentHandler)
	app.Put("/:idInvestment", hd.UpdateInvestmentHandler)
	app.Delete("/:idInvestment", hd.DeleteInvestmentHandler)
	app.Patch("/:idInvestment/update-status", hd.UpdateIsDoneInvestmentHandler)
}
