// models/absenharian.go
package models

import (
	"time"
)

type AbsenHarian struct {
	IdAbsenHarian int       `gorm:"primaryKey;column:IdAbsenHarian" json:"IdAbsenHarian"`
	IdKaryawan    *int      `gorm:"column:IdKaryawan" json:"IdKaryawan"`
	IdBagian      *int      `gorm:"column:IdBagian" json:"IdBagian"`
	TglAbsen      time.Time `gorm:"column:TglAbsen;type:date" json:"TglAbsen"`
	Absen         *string   `gorm:"column:Absen" json:"Absen"`

	DetailKaryawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
	Bagian         Bagian   `gorm:"foreignKey:IdBagian;references:IdBagian" json:"Bagian"`
}

func (AbsenHarian) TableName() string {
	return "tabsen_harian"
}
