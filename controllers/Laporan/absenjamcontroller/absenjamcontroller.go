package absenjamcontroller

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
	idBagian := c.Params("IdBagian")
	tglAbsen := c.Params("TglAbsen")

	var data []models.AbsenJam
	query := DB.Where("TglAbsen = ?", tglAbsen)

	if idBagian != "ALL" && idBagian != "all" {
		query = query.Where("IdBagian = ?", idBagian)
	}

	if err := query.Preload("Karyawan").Find(&data).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil, 0)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", data, len(data))
}

func Show(c *fiber.Ctx) error {
	idStr := c.Params("IdAbsenJam")
	IdAbsenJam64, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return respond(c, http.StatusBadRequest, "IdAbsenJam tidak valid", nil, 0)
	}
	IdAbsenJam := uint(IdAbsenJam64)

	var data models.AbsenJam

	columns := []string{
		"IdAbsenJam",
		"IdKaryawan",
		"IdBagian",
		"TglAbsen",
		"JamMasuk",
		"JamPulang",
		"JumlahJam",
	}

	err = DB.Select(columns).First(&data, "IdAbsenJam = ?", IdAbsenJam).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil, 0)
		}
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil, 0)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", data, 1)
}

func Create(c *fiber.Ctx) error {
	data := new(models.AbsenJam)

	if err := c.BodyParser(data); err != nil {
		return respond(c, http.StatusBadRequest, "Request tidak valid", nil, 0)
	}

	if err := DB.Create(&data).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal membuat data", nil, 0)
	}

	return respond(c, http.StatusCreated, "Data berhasil dibuat", data, 1)
}

func Update(c *fiber.Ctx) error {
	IdAbsenJam := c.Params("IdAbsenJam")
	data := new(models.AbsenJam)

	if err := DB.First(&models.AbsenJam{}, "IdAbsenJam = ?", IdAbsenJam).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil, 0)
		}
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil, 0)
	}

	if err := c.BodyParser(data); err != nil {
		return respond(c, http.StatusBadRequest, "Request tidak valid", nil, 0)
	}

	idUint, _ := strconv.ParseUint(IdAbsenJam, 10, 64)
	data.IdAbsenJam = int(idUint)

	if err := DB.Save(&data).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal memperbarui data", nil, 0)
	}

	return respond(c, http.StatusOK, "Data berhasil diperbarui", data, 1)
}

func Delete(c *fiber.Ctx) error {
	IdAbsenJam := c.Params("IdAbsenJam")

	if err := DB.Delete(&models.AbsenJam{}, "IdAbsenJam = ?", IdAbsenJam).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menghapus data", nil, 0)
	}

	return respond(c, http.StatusOK, "Data berhasil dihapus", nil, 0)
}
