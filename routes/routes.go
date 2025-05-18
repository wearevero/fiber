package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/controllers/bagiancontroller"
	"github.com/wearevero/fiber/controllers/jabatancontroller"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	master_data := v1.Group("/master-data")
	bagian := master_data.Group("/bagian")
	jabatan := master_data.Group("jabatan")

	// Define routes bagian
	bagian.Get("/", bagiancontroller.Index)
	bagian.Get("/:IdBagian", bagiancontroller.Show)
	bagian.Post("/", bagiancontroller.Create)
	bagian.Put("/:IdBagian", bagiancontroller.Update)
	bagian.Delete("/:IdBagian", bagiancontroller.Delete)

	// Define routes jabatan
	jabatan.Get("/", jabatancontroller.Index)
	jabatan.Get("/:IdJabatan", jabatancontroller.Show)
	jabatan.Post("/", jabatancontroller.Create)
	jabatan.Put("/:IdJabatan", jabatancontroller.Update)
	jabatan.Delete("/:IdJabatan", jabatancontroller.Delete)
}
