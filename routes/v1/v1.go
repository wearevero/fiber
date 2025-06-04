package v1

import "github.com/gofiber/fiber/v2"

func RegisterV1Routes(api fiber.Router) {
	v1 := api.Group("/v1")

	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello from Fiber API v1!"})
	})

	registerMasterDataRoutes(v1)
	registerLaporanRoutes(v1)
}
