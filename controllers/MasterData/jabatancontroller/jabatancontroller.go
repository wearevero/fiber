package jabatancontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/models"
	"gorm.io/gorm"
)

func respond(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func Index(c *fiber.Ctx) error {
	var tjabatan []models.Jabatan

	if err := models.DB.Find(&tjabatan).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", tjabatan)
}

func Show(c *fiber.Ctx) error {
	IdJabatan := c.Params("IdJabatan")
	var tjabatan models.Jabatan

	if err := models.DB.First(&tjabatan, "IdJabatan = ?", IdJabatan).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil)
		}
		return respond(c, http.StatusInternalServerError, "Terjadi kesalahan saat mengambil data", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", tjabatan)
}

func Create(c *fiber.Ctx) error {
	var tjabatan models.Jabatan

	if err := c.BodyParser(&tjabatan); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	if err := models.DB.Create(&tjabatan).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menyimpan data", err.Error())
	}

	return respond(c, http.StatusCreated, "Data berhasil ditambahkan", tjabatan)
}

func Update(c *fiber.Ctx) error {
	IdJabatan := c.Params("IdJabatan")
	var tjabatan models.Jabatan

	if err := c.BodyParser(&tjabatan); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	result := models.DB.Model(&models.Jabatan{}).Where("IdJabatan = ?", IdJabatan).Updates(tjabatan)
	if result.RowsAffected == 0 {
		return respond(c, http.StatusNotFound, "Data tidak ditemukan atau tidak ada perubahan", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil diupdate", tjabatan)
}

func Delete(c *fiber.Ctx) error {
	IdJabatan := c.Params("IdJabatan")

	result := models.DB.Delete(&models.Jabatan{}, "IdJabatan = ?", IdJabatan)
	if result.RowsAffected == 0 {
		return respond(c, http.StatusNotFound, "Tidak dapat menghapus data / data tidak ditemukan", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil dihapus", nil)
}
