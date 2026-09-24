package investments

import "github.com/gofiber/fiber/v3"

// SetupRoutes initializes the router and defines the routes
func SetupRoutes(app fiber.Router, hd IHandlerService) {
	app.RouteChain("/investments/").
		Get(hd.GetInvestmentHandler).
		Post(hd.AddInvestmentHandler)

	expenseID := app.Group("/investments/:idInvestment")
	expenseID.Get("", hd.GetByIdInvestmentHandler)
	expenseID.Put("", hd.UpdateInvestmentHandler)
	expenseID.Delete("", hd.DeleteInvestmentHandler)

	expenseID.Patch("/update-status", hd.UpdateIsDoneInvestmentHandler)
}
