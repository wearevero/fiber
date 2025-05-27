package absenlemburcontroller

import (
	"fmt"
	"net/http"
	"strconv"

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

	var data []models.AbsenLembur
	query := DB.Where("TglAbsen = ?", TglAbsen)

	if IdBagian != "ALL" && IdBagian != "all" {
		query = query.Where("IdBagian = ?", IdBagian)
	}

	if err := query.Preload("Karyawan").Find(&data).Error; err != nil {
		fmt.Println("Query error:", err)
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil, 0)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", data, len(data))
}

func Show(c *fiber.Ctx) error {
	idStr := c.Params("IdAbsenLembur")
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return respond(c, http.StatusBadRequest, "IdAbsenLembur tidak valid", nil, 0)
	}
	id := uint(idUint64)

	var data models.AbsenLembur

	err = DB.First(&data, "IdAbsenLembur = ?", id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil, 0)
		}
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil, 0)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", data, 1)
}

func Create(c *fiber.Ctx) error {
	data := new(models.AbsenLembur)

	if err := c.BodyParser(data); err != nil {
		return respond(c, http.StatusBadRequest, "Request tidak valid", nil, 0)
	}

	if err := DB.Create(&data).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal membuat data", nil, 0)
	}

	return respond(c, http.StatusCreated, "Data berhasil dibuat", data, 1)
}

func Delete(c *fiber.Ctx) error {
	idStr := c.Params("IdAbsenLembur")
	idUint64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return respond(c, http.StatusBadRequest, "IdAbsenLembur tidak valid", nil, 0)
	}
	id := uint(idUint64)

	if err := DB.Delete(&models.AbsenLembur{}, "IdAbsenLembur = ?", id).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menghapus data", nil, 0)
	}

	return respond(c, http.StatusOK, "Data berhasil dihapus", nil, 0)
}
