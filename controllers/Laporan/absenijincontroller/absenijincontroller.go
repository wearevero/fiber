package absenijincontroller

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

func respond(c *fiber.Ctx, status int, message string, data interface{}, count int) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"count":   count,
		"data":    data,
	})
}

func Index(c *fiber.Ctx) error {
	idBagian := c.Params("IdBagian")
	tglIjin := c.Params("TglIjin")

	var absenIjin []models.AbsenIjin
	var bagianList []models.Bagian

	if err := DB.Find(&bagianList).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data bagian", nil, 0)
	}

	query := DB.Where("TglIjin = ?", tglIjin)

	if idBagian != "ALL" && idBagian != "all" {
		query = query.Where("IdBagian = ?", idBagian)
	}

	if err := query.
		Preload("DetailKaryawan").
		Preload("Bagian").
		Find(&bagianList).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data absen ijin", nil, 0)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     http.StatusOK,
		"message":    "Data berhasil ditemukan",
		"count":      len(absenIjin),
		"data":       absenIjin,
		"bagianList": bagianList,
	})
}
