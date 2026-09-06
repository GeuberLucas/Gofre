package expense

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app fiber.Router, hd IHandlerExpense) {
	app.RouteChain("/expense/").
		Get(hd.GetExpenseHandler).
		Post(hd.AddExpenseHandler)

	expenseID := app.Group("/expense/:idExpense")

	expenseID.Get("", hd.GetByIdExpenseHandler)
	expenseID.Put("", hd.UpdateExpenseHandler)
	expenseID.Delete("", hd.DeleteExpenseHandler)
	expenseID.Patch("/update-status", hd.UpdateIsPaidHandler)
}
