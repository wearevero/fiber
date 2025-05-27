package models

import (
	"time"
)

type AbsenLembur struct {
	IdAbsenLembur int       `gorm:"column:IdAbsenLembur;primaryKey" json:"IdAbsenLembur"`
	IdKaryawan    string    `gorm:"column:IdKaryawan" json:"IdKaryawan"`
	IdBagian      string    `gorm:"column:IdBagian" json:"IdBagian"`
	TglAbsen      time.Time `gorm:"column:TglAbsen" json:"TglAbsen"`
	JamMasuk      time.Time `gorm:"column:JamMasuk" json:"JamMasuk"`
	JamPulang     time.Time `gorm:"column:JamPulang" json:"JamPulang"`
	JumlahJam     string    `gorm:"column:JumlahJam" json:"JumlahJam"`
	JumlahJamx    time.Time `gorm:"column:JumlahJamx" json:"JumlahJamx"`
	Proses        string    `gorm:"column:Proses" json:"Proses"`
	Shift         string    `gorm:"column:Shift" json:"Shift"`
	Keterangan    string    `gorm:"column:Keterangan" json:"Keterangan"`

	Karyawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
}

func (AbsenLembur) TableName() string {
	return "tabsen_lembur"
}
