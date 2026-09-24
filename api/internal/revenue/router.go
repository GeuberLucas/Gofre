package revenue

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app fiber.Router, hd IRevenueHandler) {

	app.RouteChain("/revenue/").
		Get(hd.GetRevenueHandler).
		Post(hd.AddRevenueHandler)

	revenueID := app.Group("/revenue/:idRevenue")
	revenueID.Get("", hd.GetByIdRevenueHandler)
	revenueID.Put("", hd.UpdateRevenueHandler)
	revenueID.Delete("", hd.DeleteRevenueHandler)
	revenueID.Patch("/update-status", hd.UpdateIsRecievedHandler)
}
