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

// Helper response
func respond(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func Index(c *fiber.Ctx) error {
	var users []models.User

	if err := DB.Preload("Karyawan").Find(&users).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", users)
}

func Show(c *fiber.Ctx) error {
	IdUser := c.Params("IdUser")
	var user models.User

	if err := DB.Preload("Karyawan").First(&user, "IdUser = ?", IdUser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return respond(c, http.StatusNotFound, "Data tidak ditemukan", nil)
		}
		return respond(c, http.StatusInternalServerError, "Terjadi kesalahan saat mengambil data", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil ditemukan", user)
}

func Create(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	if user.Username == nil || user.Password == nil || user.IdKaryawan == nil || *user.IdKaryawan == 0 {
		return respond(c, http.StatusBadRequest, "Username, Password, dan IdKaryawan wajib diisi", nil)
	}

	if user.Spesial != nil && *user.Spesial != "Y" && *user.Spesial != "N" {
		return respond(c, http.StatusBadRequest, "Nilai Spesial harus Y atau N", nil)
	}

	// Hashing password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengenkripsi password", err.Error())
	}
	hashed := string(hashedPassword)
	user.Password = &hashed

	if err := DB.Create(&user).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal menyimpan data", err.Error())
	}

	return respond(c, http.StatusCreated, "Data berhasil ditambahkan", user)
}

func Update(c *fiber.Ctx) error {
	IdUser := c.Params("IdUser")
	var userInput models.User

	if err := c.BodyParser(&userInput); err != nil {
		return respond(c, http.StatusBadRequest, "Format data tidak valid", err.Error())
	}

	// Cek apakah password ingin diupdate
	if userInput.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*userInput.Password), bcrypt.DefaultCost)
		if err != nil {
			return respond(c, http.StatusInternalServerError, "Gagal mengenkripsi password", err.Error())
		}
		hashed := string(hashedPassword)
		userInput.Password = &hashed
	}

	result := DB.Model(&models.User{}).Where("IdUser = ?", IdUser).Updates(userInput)
	if result.RowsAffected == 0 {
		return respond(c, http.StatusNotFound, "Data tidak ditemukan atau tidak ada perubahan", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil diupdate", userInput)
}

func Delete(c *fiber.Ctx) error {
	IdUser := c.Params("IdUser")

	result := DB.Delete(&models.User{}, "IdUser = ?", IdUser)
	if result.RowsAffected == 0 {
		return respond(c, http.StatusNotFound, "Tidak dapat menghapus data / data tidak ditemukan", nil)
	}

	return respond(c, http.StatusOK, "Data berhasil dihapus", nil)
}
