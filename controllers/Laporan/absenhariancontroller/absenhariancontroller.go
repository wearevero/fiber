package absenhariancontroller

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
	IdBagian := c.Params("IdBagian")
	TglAbsen := c.Params("TglAbsen")

	var data []models.AbsenHarian
	query := DB.Where("TglAbsen = ?", TglAbsen)

	if IdBagian != "ALL" && IdBagian != "all" {
		query = query.Where("IdBagian = ?", IdBagian)
	}

	if err := query.Find(&data).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil, 0)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", data, len(data))
}
