package karyawancontroller

import (
	"math"
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

func respond(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func Index(c *fiber.Ctx) error {
	var data []models.Karyawan

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
	if err := DB.Model(&models.Karyawan{}).Count(&total).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	if page > totalPages && totalPages != 0 {
		return respond(c, http.StatusNotFound, "Halaman tidak ditemukan", nil)
	}

	err := DB.
		Limit(limit).
		Offset(offset).
		Order("IdKaryawan DESC").
		Find(&data).Error

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

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     http.StatusOK,
		"message":    "Data berhasil ditemukan",
		"data":       data,
		"pagination": pagination,
	})
}

func Show(c *fiber.Ctx) error {
	IdKaryawan := c.Params("IdKaryawan")
	var data models.Karyawan

	if err := DB.First(&data, "IdKaryawan = ?", IdKaryawan).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil)
		}
		return respond(c, http.StatusInternalServerError, "Terjadi kesalahan saat mengambil data", nil)
	}
	return respond(c, http.StatusOK, "Data berhasil ditemukan", data)
}

func Create(c *fiber.Ctx) error {
	data := new(models.Karyawan)

	if err := c.BodyParser(data); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", nil)
	}

	if err := DB.Create(&data).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menyimpan data", nil)
	}

	return respond(c, http.StatusCreated, "Data berhasil ditambahkan", data)
}

func Update(c *fiber.Ctx) error {
	IdKaryawan := c.Params("IdKaryawan")
	data := new(models.Karyawan)

	if err := DB.First(&models.Karyawan{}, "IdKaryawan = ?", IdKaryawan).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil)
		}
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
	}

	if err := c.BodyParser(data); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", nil)
	}

	idInt, _ := strconv.ParseUint(IdKaryawan, 10, 64)
	data.IdKaryawan = int(idInt)

	if err := DB.Save(&data).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal memperbarui data", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil diupdate", data)
}

func Delete(c *fiber.Ctx) error {
	IdKaryawan := c.Params("IdKaryawan")

	if err := DB.Delete(&models.Karyawan{}, "IdKaryawan = ?", IdKaryawan).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menghapus data", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil dihapus", nil)
}
