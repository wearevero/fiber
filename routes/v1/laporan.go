package v1

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/controllers/Laporan/absenhariancontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenhariinicontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenijincontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenjamcontroller"
	"github.com/wearevero/fiber/controllers/Laporan/absenlemburcontroller"
)

func registerLaporanRoutes(v1 fiber.Router) {
	laporan := v1.Group("/laporan")

	laporan.Get("/cek-absen-bagian/:IdBagian/:TglAbsen", absenhariinicontroller.CekAbsenBagian)

	jam := laporan.Group("/absen-jam")
	jam.Get("/:IdBagian/:TglAbsen", absenjamcontroller.Index)

	harian := laporan.Group("/absen-harian")
	harian.Get("/:IdBagian/:bulan/:tahun", absenhariancontroller.Index)

	lembur := laporan.Group("/absen-lembur")
	lembur.Get("/:IdBagian/:TglAbsen", absenlemburcontroller.Index)

	ijin := laporan.Group("/absen-ijin")
	ijin.Get("/:IdBagian/:TglIjin", absenijincontroller.Index)
}
