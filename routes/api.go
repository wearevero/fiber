package routes

import (
	"github.com/gofiber/fiber/v2"
	v1 "github.com/wearevero/fiber/routes/v1"
)

func RegisterAPIRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1.RegisterV1Routes(api)
}
