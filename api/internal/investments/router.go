package investments

import "github.com/gofiber/fiber/v3"

// SetupRoutes initializes the router and defines the routes
func SetupRoutes(app fiber.Router, hd IHandlerService) {
	app.RouteChain("/investments/").
		Get(hd.GetInvestmentHandler).
		Post(hd.AddInvestmentHandler)

	app.RouteChain("/investments/:idInvestment").
		Get(hd.GetByIdInvestmentHandler).
		Put(hd.UpdateInvestmentHandler).
		Delete(hd.DeleteInvestmentHandler)

	app.Patch("/investments/:idInvestment/update-status", hd.UpdateIsDoneInvestmentHandler)
}
