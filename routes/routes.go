package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/controllers/Laporan/absenhariancontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenjamcontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenlemburcontroller"
	"github.com/wearevero/fiber/controllers/MasterData/bagiancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/jabatancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/karyawancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/usercontroller"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/hello", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello from Fiber API v1!",
		})
	})

	// Master Data Routes
	master_data := v1.Group("/master-data")
	bagian := master_data.Group("/bagian")
	jabatan := master_data.Group("/jabatan")
	karyawan := master_data.Group("/karyawan")
	user := master_data.Group("/user")

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

	// Laporan Routes
	laporan := v1.Group("/laporan")
	jam := laporan.Group("/absen-jam")
	harian := laporan.Group("/absen-harian")
	lembur := laporan.Group("/absen-lembur")

	// Define laporan/absen-jam routes
	jam.Get("/:IdBagian/:TglAbsen", absenjamcontroller.Index)
	jam.Get("/detail/:IdAbsenJam", absenjamcontroller.Show)

	// Define laporan/absen-harian routes
	harian.Get("/:IdBagian/:TglAbsen", absenhariancontroller.Index)

	// Define laporan/absen-lembur routes
	lembur.Get("/:IdBagian/:TglAbsen", absenlemburcontroller.Index)
}
