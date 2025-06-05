package models

import (
	"time"
)

type AbsenJam struct {
	IdAbsenJam int       `gorm:"column:IdAbsenJam;primaryKey" json:"IdAbsenJam"`
	IdKaryawan *int      `gorm:"column:IdKaryawan" json:"IdKaryawan"`
	IdBagian   *int      `gorm:"column:IdBagian" json:"IdBagian"`
	TglAbsen   time.Time `gorm:"column:TglAbsen" json:"TglAbsen"`
	JamMasuk   string    `gorm:"column:JamMasuk" json:"JamMasuk"`
	JamPulang  string    `gorm:"column:JamPulang" json:"JamPulang"`
	JumlahJam  *float64  `gorm:"column:JumlahJam" json:"JumlahJam"`
	JumlahJamX string    `gorm:"column:JumlahJamX" json:"JumlahJamX"`
	Proses     string    `gorm:"column:Proses" json:"Proses"`
	Shift      string    `gorm:"column:Shift" json:"Shift"`

	DetailKaryawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
	Bagian         Bagian   `gorm:"foreignKey:IdBagian;references:IdBagian" json:"Bagian"`
}

func (AbsenJam) TableName() string {
	return "tabsen_jam"
}
