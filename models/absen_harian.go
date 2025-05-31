package models

import (
	"time"
)

type AbsenHarian struct {
	IdAbsenHarian int       `gorm:"primaryKey;column:IdAbsenHarian" json:"IdAbsenHarian"`
	IdKaryawan    int       `gorm:"column:IdKaryawan;size:255" json:"IdKaryawan"`
	IdBagian      string    `gorm:"column:IdBagian;type:longtext" json:"IdBagian"`
	TglAbsen      time.Time `gorm:"column:TglAbsen;type:date" json:"TglAbsen"`
	Absen         string    `gorm:"column:Absen;type:longtext" json:"Absen"`

	Karyawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
}

func (AbsenHarian) TableName() string {
	return "tabsen_harian"
}
