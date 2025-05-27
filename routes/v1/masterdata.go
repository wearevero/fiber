package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/controllers/MasterData/bagiancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/jabatancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/karyawancontroller"
	"github.com/wearevero/fiber/controllers/MasterData/karyawankeluarcontroller"
	"github.com/wearevero/fiber/controllers/MasterData/usercontroller"
)

func registerMasterDataRoutes(v1 fiber.Router) {
	master := v1.Group("/master-data")

	bagian := master.Group("/bagian")
	bagian.Get("/", bagiancontroller.Index)
	bagian.Get("/:IdBagian", bagiancontroller.Show)
	bagian.Post("/", bagiancontroller.Create)
	bagian.Patch("/:IdBagian", bagiancontroller.Update)
	bagian.Delete("/:IdBagian", bagiancontroller.Delete)

	jabatan := master.Group("/jabatan")
	jabatan.Get("/", jabatancontroller.Index)
	jabatan.Get("/:IdJabatan", jabatancontroller.Show)
	jabatan.Post("/", jabatancontroller.Create)
	jabatan.Patch("/:IdJabatan", jabatancontroller.Update)
	jabatan.Delete("/:IdJabatan", jabatancontroller.Delete)

	karyawan := master.Group("/karyawan")
	karyawan.Get("/", karyawancontroller.Index)
	karyawan.Get("/:IdKaryawan", karyawancontroller.Show)
	karyawan.Post("/", karyawancontroller.Create)
	karyawan.Patch("/:IdKaryawan", karyawancontroller.Update)
	karyawan.Delete("/:IdKaryawan", karyawancontroller.Delete)

	user := master.Group("/user")
	user.Get("/", usercontroller.Index)
	user.Get("/:IdUser", usercontroller.Show)
	user.Post("/", usercontroller.Create)
	user.Patch("/:IdUser", usercontroller.Update)
	user.Delete("/:IdUser", usercontroller.Delete)

	karyawanKeluar := master.Group("/karyawan-keluar")
	karyawanKeluar.Get("/", karyawankeluarcontroller.Index)
	karyawanKeluar.Get("/:IdKeluar", karyawankeluarcontroller.Show)
	karyawanKeluar.Post("/", karyawankeluarcontroller.Create)
	karyawanKeluar.Patch("/:IdKeluar", karyawankeluarcontroller.Update)
	karyawanKeluar.Delete("/:IdKeluar", karyawankeluarcontroller.Delete)
}
