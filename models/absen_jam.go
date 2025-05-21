package models

import (
	"time"
)

type AbsenJam struct {
	IdAbsenJam uint      `gorm:"primaryKey;column:IdAbsenJam" json:"IdAbsenJam"`
	IdKaryawan string    `gorm:"column:IdKaryawan" json:"IdKaryawan"`
	IdBagian   string    `gorm:"column:IdBagian" json:"IdBagian"`
	TglAbsen   time.Time `gorm:"column:TglAbsen" json:"TglAbsen"`
	JamMasuk   time.Time `gorm:"column:JamMasuk" json:"JamMasuk"`
	JamPulang  time.Time `gorm:"column:JamPulang" json:"JamPulang"`
	JumlahJam  string    `gorm:"column:JumlahJam" json:"JumlahJam"`
	JumlahJamX string    `gorm:"column:JumlahJamX" json:"JumlahJamX"`
	Proses     string    `gorm:"column:Proses" json:"Proses"`
	Shift      string    `gorm:"column:Shift" json:"Shift"`

	Karyawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
}

func (AbsenJam) TableName() string {
	return "tabsen_jam"
}
