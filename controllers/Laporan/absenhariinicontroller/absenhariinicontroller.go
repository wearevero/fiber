package absenhariinicontroller

import (
	"fmt"
	"math"
	"net/http"
	"time"

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

func CekAbsenBagian(c *fiber.Ctx) error {
	idBagian := c.Params("IdBagian")
	tglAbsen := c.Params("TglAbsen")

	fmt.Printf("DEBUG - IdBagian: %s, TglAbsen: %s\n", idBagian, tglAbsen)

	if _, err := time.Parse("2006-01-02", tglAbsen); err != nil {
		return respond(c, http.StatusBadRequest, "Format tanggal tidak valid. Gunakan format YYYY-MM-DD", nil, 0)
	}

	var bagianList []models.Bagian
	if err := DB.Find(&bagianList).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data bagian", nil, 0)
	}

	var semuaKaryawanAktif []models.Karyawan
	karyawanQuery := DB.Where("Aktif = ?", "Ya").Preload("MasterBagian")
	if idBagian != "ALL" && idBagian != "all" {
		karyawanQuery = karyawanQuery.Where("IdBagian = ?", idBagian)
	}

	if err := karyawanQuery.Find(&semuaKaryawanAktif).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data karyawan aktif", nil, 0)
	}

	totalKaryawanAktif := int64(len(semuaKaryawanAktif))

	var absenJam []models.AbsenJam
	absenQuery := DB.Where("DATE(TglAbsen) = ?", tglAbsen).
		Preload("DetailKaryawan").
		Preload("Bagian")

	if idBagian != "ALL" && idBagian != "all" {
		absenQuery = absenQuery.Where("IdBagian = ?", idBagian)
	}

	if err := absenQuery.Find(&absenJam).Error; err != nil {
		return respond(c, http.StatusInternalServerError, "Gagal mengambil data absen jam", nil, 0)
	}

	karyawanSudahAbsenMap := make(map[int]bool)
	var karyawanSudahAbsen int64 = 0

	for _, absen := range absenJam {

		if absen.IdKaryawan != nil && absen.JamMasuk != "" && absen.JamMasuk != "00:00:00" {
			if !karyawanSudahAbsenMap[*absen.IdKaryawan] {
				karyawanSudahAbsenMap[*absen.IdKaryawan] = true
				karyawanSudahAbsen++
			}
		}
	}

	var listKaryawanBelumAbsen []models.Karyawan
	for _, karyawan := range semuaKaryawanAktif {
		if !karyawanSudahAbsenMap[karyawan.IdKaryawan] {
			listKaryawanBelumAbsen = append(listKaryawanBelumAbsen, karyawan)
		}
	}

	karyawanBelumAbsen := int64(len(listKaryawanBelumAbsen))

	fmt.Printf("DEBUG - Total Karyawan: %d, Sudah Absen: %d, Belum Absen: %d\n",
		totalKaryawanAktif, karyawanSudahAbsen, karyawanBelumAbsen)

	var persentaseKehadiran float64 = 0
	if totalKaryawanAktif > 0 {
		persentaseKehadiran = (float64(karyawanSudahAbsen) / float64(totalKaryawanAktif)) * 100
	}

	var namaBagian string = "Semua Bagian"
	if idBagian != "ALL" && idBagian != "all" {
		for _, bagian := range bagianList {
			if fmt.Sprintf("%d", bagian.IdBagian) == idBagian {
				namaBagian = bagian.NamaBagian
				break
			}
		}
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"status":               http.StatusOK,
		"message":              "Data berhasil ditemukan",
		"total_karyawan":       totalKaryawanAktif,
		"karyawan_sudah_absen": karyawanSudahAbsen,
		"karyawan_belum_absen": karyawanBelumAbsen,
		"persentase_kehadiran": math.Round(persentaseKehadiran*100) / 100,
		"bagian_list":          bagianList,
		"tanggal_absen":        tglAbsen,
		"id_bagian":            idBagian,
		"nama_bagian":          namaBagian,
		"detail_belum_absen":   listKaryawanBelumAbsen,
	})
}
