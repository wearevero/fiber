package karyawancontroller

import (
	"math"
	"net/http"
	"strconv"
	"strings"

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

	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 || limit > 100 {
		limit = 10
	}

	aktif := c.Query("aktif", "Ya")
	search := strings.TrimSpace(c.Query("search", ""))

	offset := (page - 1) * limit

	var query *gorm.DB
	if aktif == "semua" {

		query = DB.Model(&models.Karyawan{})
	} else {

		if aktif != "Ya" && aktif != "Tidak" {
			aktif = "Ya"
		}
		query = DB.Model(&models.Karyawan{}).Where("Aktif = ?", aktif)
	}

	if search != "" {

		query = query.Where(
			"NikKaryawan LIKE ? OR NamaLengkap LIKE ? OR NamaPanggilan LIKE ? OR Jabatan LIKE ?",
			"%"+search+"%", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
	}

	var activeCount, inactiveCount, allCount int64

	if err := DB.Model(&models.Karyawan{}).Where("Aktif = ?", "Ya").Count(&activeCount).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menghitung karyawan aktif", nil)
	}

	if err := DB.Model(&models.Karyawan{}).Where("Aktif = ?", "Tidak").Count(&inactiveCount).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menghitung karyawan tidak aktif", nil)
	}

	if err := DB.Model(&models.Karyawan{}).Count(&allCount).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menghitung semua karyawan", nil)
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	if page > totalPages && totalPages > 0 {
		return respond(c, http.StatusNotFound, "Halaman tidak ditemukan", nil)
	}

	err = query.
		Preload("MasterJabatan").
		Preload("MasterBagian").
		Limit(limit).
		Offset(offset).
		Order("IdKaryawan ASC").
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
		"has_previous": page > 1,
		"has_next":     page < totalPages,
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":     http.StatusOK,
		"message":    "Data berhasil ditemukan",
		"data":       data,
		"pagination": pagination,
		"statistics": fiber.Map{
			"active_count":   activeCount,
			"inactive_count": inactiveCount,
			"all_count":      allCount,
		},
		"current_filter": aktif,
		"search_term":    search,
	})
}

func CountActive(c *fiber.Ctx) error {
	var count int64
	if err := DB.Model(&models.Karyawan{}).
		Where("Aktif = ?", "Ya").
		Count(&count).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menghitung data aktif", nil)
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":  http.StatusOK,
		"message": "Jumlah karyawan aktif berhasil dihitung",
		"data":    count,
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
