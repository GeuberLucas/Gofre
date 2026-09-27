package profile

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app fiber.Router, hd IProfileHandler) {
	app.RouteChain("/profile").
		Get(hd.GetProfileHandler).
		Post(hd.AddProfileHandler).
		Patch(hd.UpdateProfileHandler).
		Delete(hd.DeleteProfileHandler)

}
