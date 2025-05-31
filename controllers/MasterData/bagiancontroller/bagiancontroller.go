package bagiancontroller

import (
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
	var data []models.Bagian
	if err := DB.Find(&data).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil, 0)
	}
	return respond(c, http.StatusOK, "Data berhasil ditemukan", data, len(data))
}

func Show(c *fiber.Ctx) error {
	IdBagian := c.Params("IdBagian")
	var data models.Bagian

	if err := DB.First(&data, "IdBagian = ?", IdBagian).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil, 0)
		}
		return respond(c, http.StatusInternalServerError, "Terjadi kesalahan saat mengambil data", nil, 0)
	}
	return respond(c, http.StatusOK, "Data berhasil ditemukan", data, 1)
}

func Create(c *fiber.Ctx) error {
	data := new(models.Bagian)

	if err := c.BodyParser(data); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", nil, 0)
	}

	if err := DB.Create(&data).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menyimpan data", nil, 0)
	}

	return respond(c, http.StatusCreated, "Data berhasil ditambahkan", data, 1)
}

func Update(c *fiber.Ctx) error {
	IdBagian := c.Params("IdBagian")
	data := new(models.Bagian)

	if err := DB.First(&models.Bagian{}, "IdBagian = ?", IdBagian).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil, 0)
		}
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil, 0)
	}

	if err := c.BodyParser(data); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", nil, 0)
	}

	idUint, _ := strconv.ParseUint(IdBagian, 10, 64)
	data.IdBagian = int(idUint)

	if err := DB.Save(&data).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal memperbarui data", nil, 0)
	}

	return respond(c, http.StatusOK, "Data berhasil diupdate", data, 1)
}

func Delete(c *fiber.Ctx) error {
	IdBagian := c.Params("IdBagian")

	if err := DB.Delete(&models.Bagian{}, "IdBagian = ?", IdBagian).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menghapus data", nil, 0)
	}

	return respond(c, http.StatusOK, "Data berhasil dihapus", nil, 0)
}
