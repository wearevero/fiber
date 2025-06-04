package absenlemburcontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(database *gorm.DB) {
	DB = database
}

// Fungsi response standar
func respond(c *fiber.Ctx, status int, message string, data interface{}, count int) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"count":   count,
		"data":    data,
	})
}

// Endpoint utama
func Index(c *fiber.Ctx) error {
	idBagian := c.Params("IdBagian")
	tglAbsen := c.Params("TglAbsen")

	var absenLembur []models.AbsenLembur
	var bagianList []models.Bagian

	// Ambil semua bagian
	if err := DB.Find(&bagianList).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data bagian", nil, 0)
	}

	query := DB.Where("TglAbsen = ?", tglAbsen)

	if idBagian != "ALL" && idBagian != "all" {
		query = query.Where("IdBagian = ?", idBagian)
	}

	if err := query.
		Preload("DetailKaryawan").
		Preload("Bagian").
		Find(&absenLembur).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data absen lembur", nil, 0)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     http.StatusOK,
		"message":    "Data berhasil ditemukan",
		"count":      len(absenLembur),
		"data":       absenLembur,
		"bagianList": bagianList,
	})
}
