package response

import (
	"github.com/gofiber/fiber"
)

// JSONResponse sends a JSON response with the given status code and data
func JSONResponse(c *fiber.Ctx, statusCode int, data interface{}) error {

	return c.Status(statusCode).JSON(data)
}

func ErrorResponse(c *fiber.Ctx, statusCode int, err error) error {

	return JSONResponse(c, statusCode, fiber.Map{
		"erro": err.Error(),
	})
}
