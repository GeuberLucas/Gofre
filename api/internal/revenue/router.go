package revenue

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app fiber.Router, hd IRevenueHandler) {
	app.RouteChain("/revenue/").
		Get(hd.GetRevenueHandler).
		Post(hd.AddRevenueHandler)

	app.RouteChain("/revenue/:idRevenue").
		Get(hd.GetByIdRevenueHandler).
		Put(hd.UpdateRevenueHandler).
		Delete(hd.DeleteRevenueHandler)

	app.Patch("/revenue/:idRevenue/update-status", hd.UpdateIsRecievedHandler)
}
