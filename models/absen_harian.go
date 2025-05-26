package models

import (
	"time"
)

type AbsenHarian struct {
	IdAbsenHarian uint      `gorm:"primaryKey;column:IdAbsenHarian" json:"IdAbsenHarian"`
	IdKaryawan    string    `gorm:"column:IdKaryawan;size:255" json:"IdKaryawan"`  // VARCHAR(255)
	IdBagian      string    `gorm:"column:IdBagian;type:longtext" json:"IdBagian"` // longtext
	TglAbsen      time.Time `gorm:"column:TglAbsen;type:date" json:"TglAbsen"`     // date
	Absen         string    `gorm:"column:Absen;type:longtext" json:"Absen"`

	Karyawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
}

func (AbsenHarian) TableName() string {
	return "tabsen_harian"
}
