package bagiancontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/models"
	"gorm.io/gorm"
)

// helper response for consistency
func respond(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func Index(c *fiber.Ctx) error {
	var tbagian []models.Bagian

	if err := models.DB.Find(&tbagian).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", tbagian)
}

func Show(c *fiber.Ctx) error {
	IdBagian := c.Params("IdBagian")
	var tbagian models.Bagian

	if err := models.DB.First(&tbagian, "IdBagian = ?", IdBagian).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil)
		}
		return respond(c, http.StatusInternalServerError, "Terjadi kesalahan saat mengambil data", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", tbagian)
}

func Create(c *fiber.Ctx) error {
	var tbagian models.Bagian

	if err := c.BodyParser(&tbagian); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	if err := models.DB.Create(&tbagian).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menyimpan data", err.Error())
	}

	return respond(c, http.StatusCreated, "Data berhasil ditambahkan", tbagian)
}

func Update(c *fiber.Ctx) error {
	IdBagian := c.Params("IdBagian")
	var tbagian models.Bagian

	if err := c.BodyParser(&tbagian); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	result := models.DB.Model(&models.Bagian{}).Where("IdBagian = ?", IdBagian).Updates(tbagian)
	if result.RowsAffected == 0 {
		return respond(c, http.StatusNotFound, "Data tidak ditemukan atau tidak ada perubahan", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil diupdate", tbagian)
}

func Delete(c *fiber.Ctx) error {
	IdBagian := c.Params("IdBagian")

	result := models.DB.Delete(&models.Bagian{}, "IdBagian = ?", IdBagian)
	if result.RowsAffected == 0 {
		return respond(c, http.StatusNotFound, "Tidak dapat menghapus data / data tidak ditemukan", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil dihapus", nil)
}
