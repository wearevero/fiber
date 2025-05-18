package karyawancontroller

import (
	"math"
	"net/http"
	"strconv"

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
	var karyawan []models.Karyawan

	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit

	var total int64
	models.DB.Model(&models.Karyawan{}).Count(&total)

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	if page > totalPages && totalPages != 0 {
		return respond(c, http.StatusNotFound, "Halaman tidak ditemukan", nil)
	}

	err := models.DB.
		Limit(limit).
		Offset(offset).
		Order("IdKaryawan DESC").
		Find(&karyawan).Error

	if err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
	}

	pagination := fiber.Map{
		"total":       total,
		"limit":       limit,
		"page":        page,
		"total_pages": totalPages,
		"previous_page": func() int {
			if page > 1 {
				return page - 1
			}
			return 0
		}(),
		"next_page": func() int {
			if page < totalPages {
				return page + 1
			}
			return 0
		}(),
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", fiber.Map{
		"data":       karyawan,
		"pagination": pagination,
	})
}

func Show(c *fiber.Ctx) error {
	IdKaryawan := c.Params("IdKaryawan")
	var tkaryawan_pribadi []models.Karyawan

	if err := models.DB.First(&tkaryawan_pribadi, "IdKaryawan = ?", IdKaryawan).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil)
		}
		return respond(c, http.StatusInternalServerError, "Terjadi kesalahan saat mengambil data", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", tkaryawan_pribadi)
}

func Create(c *fiber.Ctx) error {
	var tkaryawan_pribadi models.Karyawan

	if err := c.BodyParser(&tkaryawan_pribadi); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	if err := models.DB.Create(&tkaryawan_pribadi).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menyimpan data", err.Error())
	}

	return respond(c, http.StatusCreated, "Data berhasil ditambahkan", tkaryawan_pribadi)
}

func Update(c *fiber.Ctx) error {
	IdKaryawan := c.Params("IdKaryawan")

	var input map[string]interface{}
	if err := c.BodyParser(&input); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	delete(input, "IdKaryawan")

	var existing models.Karyawan
	if err := models.DB.First(&existing, "IdKaryawan = ?", IdKaryawan).Error; err != nil {
		return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil)
	}

	result := models.DB.Model(&existing).Updates(input)
	if result.Error != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengupdate data", result.Error.Error())
	}
	if result.RowsAffected == 0 {
		return respond(c, http.StatusOK, "Tidak ada perubahan data", nil)
	}

	models.DB.First(&existing, "IdKaryawan = ?", IdKaryawan)

	return respond(c, http.StatusOK, "Data berhasil diupdate", existing)
}

func Delete(c *fiber.Ctx) error {
	IdKaryawan := c.Params("IdKaryawan")

	result := models.DB.Delete(&models.Karyawan{}, "IdKaryawan = ?", IdKaryawan)
	if result.RowsAffected == 0 {
		return respond(c, http.StatusNotFound, "Tidak dapat menghapus data / data tidak ditemukan", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil dihapus", nil)
}
