package usercontroller

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/wearevero/fiber/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetDB(database *gorm.DB) {
	DB = database
}

// helper response for consistency
func respond(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func Index(c *fiber.Ctx) error {
	var tuser []models.User

	if err := models.DB.Find(&tuser).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", tuser)
}

func Show(c *fiber.Ctx) error {
	IdUser := c.Params("IdUser")
	var tuser models.User

	if err := models.DB.First(&tuser, "IdUser = ?", IdUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil)
		}
		return respond(c, http.StatusInternalServerError, "Terjadi kesalahan saat mengambil data", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", tuser)
}

func Create(c *fiber.Ctx) error {
	var tuser models.User

	if err := c.BodyParser(&tuser); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	if tuser.Spesial != "Y" && tuser.Spesial != "N" {
		return respond(c, http.StatusBadRequest, "Nilai Spesialis harus Y atau N", nil)
	}

	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(tuser.Password), bcrypt.DefaultCost)
	if err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengenkripsi password", err.Error())
	}
	tuser.Password = string(hashedPassword)

	if err := models.DB.Create(&tuser).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menyimpan data", err.Error())
	}

	return respond(c, http.StatusCreated, "Data berhasil ditambahkan", tuser)
}

func Update(c *fiber.Ctx) error {
	IdUser := c.Params("IdUser")
	var tuser models.User

	if err := c.BodyParser(&tuser); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	result := models.DB.Model(&models.User{}).Where("IdUser = ?", IdUser).Updates(tuser)
	if result.RowsAffected == 0 {
		return respond(c, http.StatusNotFound, "Data tidak ditemukan atau tidak ada perubahan", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil diupdate", tuser)
}

func Delete(c *fiber.Ctx) error {
	IdUser := c.Params("IdUser")

	result := models.DB.Delete(&models.User{}, "IdUser = ?", IdUser)
	if result.RowsAffected == 0 {
		return respond(c, http.StatusNotFound, "Tidak dapat menghapus data / data tidak ditemukan", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil dihapus", nil)
}
