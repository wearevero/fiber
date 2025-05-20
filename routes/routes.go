package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/controllers/MasterData/bagiancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/jabatancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/karyawancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/usercontroller"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	master_data := v1.Group("/master-data")
	bagian := master_data.Group("/bagian")
	jabatan := master_data.Group("jabatan")
	karyawan := master_data.Group("karyawan")
	user := master_data.Group("user")

	// Define master-data/bagian routes
	bagian.Get("/", bagiancontroller.Index)
	bagian.Get("/:IdBagian", bagiancontroller.Show)
	bagian.Post("/", bagiancontroller.Create)
	bagian.Patch("/:IdBagian", bagiancontroller.Update)
	bagian.Delete("/:IdBagian", bagiancontroller.Delete)

	// Define master-data/jabatan routes
	jabatan.Get("/", jabatancontroller.Index)
	jabatan.Get("/:IdJabatan", jabatancontroller.Show)
	jabatan.Post("/", jabatancontroller.Create)
	jabatan.Patch("/:IdJabatan", jabatancontroller.Update)
	jabatan.Delete("/:IdJabatan", jabatancontroller.Delete)

	// Define master-data/karyawan routes
	karyawan.Get("/", karyawancontroller.Index)
	karyawan.Get("/:IdKaryawan", karyawancontroller.Show)
	karyawan.Post("/", karyawancontroller.Create)
	karyawan.Patch("/:IdKaryawan", karyawancontroller.Update)
	karyawan.Delete("/:IdKaryawan", karyawancontroller.Delete)

	// Define master-data/user routes
	user.Get("/", usercontroller.Index)
	user.Get("/:IdUser", usercontroller.Show)
	user.Post("/", usercontroller.Create)
	user.Patch("/:IdUser", usercontroller.Update)
	user.Delete("/:IdUser", usercontroller.Delete)
}
