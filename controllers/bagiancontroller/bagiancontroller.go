package bagiancontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/models"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var bagian []models.Bagian

	if err := models.DB.Find(&bagian).Error; err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal mengambil data",
		})
	}

	return c.JSON(bagian)
}

func Show(c *fiber.Ctx) error {
	IdBagian := c.Params("IdBagian")
	var tbagian models.Bagian

	if err := models.DB.First(&tbagian, "IdBagian = ?", IdBagian).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Data tidak ditemukan",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Terjadi kesalahan saat mengambil data",
		})
	}

	return c.JSON(tbagian)
}

func Create(c *fiber.Ctx) error {
	var bagian models.Bagian

	if err := c.BodyParser(&bagian); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format data tidak valid",
			"error":   err.Error(),
		})
	}

	if err := models.DB.Create(&bagian).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Gagal menyimpan data",
			"error":   err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(bagian)
}

func Update(c *fiber.Ctx) error {
	IdBagian := c.Params("IdBagian")

	var bagian models.Bagian
	if err := c.BodyParser(&bagian); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Format data tidak valid",
			"error":   err.Error(),
		})
	}

	// Update berdasarkan IdBagian
	if models.DB.Model(&models.Bagian{}).Where("IdBagian = ?", IdBagian).Updates(bagian).RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Data tidak ditemukan atau tidak ada perubahan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil diupdate",
	})
}

func Delete(c *fiber.Ctx) error {
	IdBagian := c.Params("IdBagian")

	if models.DB.Delete(&models.Bagian{}, "IdBagian = ?", IdBagian).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Tidak dapat menghapus data / data tidak ditemukan",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data berhasil dihapus",
	})
}
