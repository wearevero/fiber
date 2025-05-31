package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/controllers/Laporan/absenhariancontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenjamcontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenlemburcontroller"
)

func registerLaporanRoutes(v1 fiber.Router) {
	laporan := v1.Group("/laporan")

	jam := laporan.Group("/absen-jam")

	jam.Get("/:IdBagian/:TglAbsen", absenjamcontroller.Index)
	jam.Get("/detail/:IdAbsenJam", absenjamcontroller.Show)

	harian := laporan.Group("/absen-harian")
	harian.Get("/:IdBagian/:TglAbsen", absenhariancontroller.Index)

	lembur := laporan.Group("/absen-lembur")
	lembur.Get("/:IdBagian/:TglAbsen", absenlemburcontroller.Index)
}
