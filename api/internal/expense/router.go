package expense

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app fiber.Router, hd IHandlerExpense) {
	route := app.Group("/expense")
	route.Get("/", hd.GetExpenseHandler)
	app.Post("/", hd.AddExpenseHandler)
	app.Get("/:idExpense", hd.GetByIdExpenseHandler)
	app.Put("/:idExpense", hd.UpdateExpenseHandler)
	app.Delete("/:idExpense", hd.DeleteExpenseHandler)
	app.Patch("/:idExpense/update-status", hd.UpdateIsPaidHandler)
}
