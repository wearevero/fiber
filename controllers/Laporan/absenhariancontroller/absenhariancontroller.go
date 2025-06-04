package absenhariancontroller

import (
	"net/http"
	"strconv"
	"time"

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

func Index(c *fiber.Ctx) error {
	idBagianStr := c.Params("IdBagian")
	bulanStr := c.Params("bulan")
	tahunStr := c.Params("tahun")

	// Validate month
	bulan, err := strconv.Atoi(bulanStr)
	if err != nil || bulan < 1 || bulan > 12 {
		return respond(c, http.StatusBadRequest, "Bulan tidak valid", nil, 0)
	}

	// Validate year
	tahun, err := strconv.Atoi(tahunStr)
	if err != nil || tahun < 2000 {
		return respond(c, http.StatusBadRequest, "Tahun tidak valid", nil, 0)
	}

	// Calculate date range for the month
	startDate := time.Date(tahun, time.Month(bulan), 1, 0, 0, 0, 0, time.UTC)
	endDate := startDate.AddDate(0, 1, -1) // Last day of the month

	var data []models.AbsenHarian
	query := DB.Preload("DetailKaryawan").Preload("Bagian").Model(&models.AbsenHarian{})

	// Filter by date range
	query = query.Where("TglAbsen BETWEEN ? AND ?", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))

	// Filter by department if not "ALL"
	if idBagianStr != "ALL" && idBagianStr != "all" {
		idBagian, err := strconv.Atoi(idBagianStr)
		if err != nil {
			return respond(c, http.StatusBadRequest, "IdBagian tidak valid", nil, 0)
		}
		query = query.Where("IdBagian = ?", idBagian)
	}

	// Execute query
	if err := query.Find(&data).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil, 0)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", data, len(data))
}
