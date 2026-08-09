package revenue

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app fiber.Router, hd IRevenueHandler) {
	route := app.Group("/expense")
	route.Get("/", hd.GetRevenueHandler)
	app.Post("/", hd.AddRevenueHandler)
	app.Get("/:idRevenue", hd.GetByIdRevenueHandler)
	app.Put("/:idRevenue", hd.UpdateRevenueHandler)
	app.Delete("/:idRevenue", hd.DeleteRevenueHandler)
	app.Patch("/:idRevenue/update-status", hd.UpdateIsRecievedHandler)
}
