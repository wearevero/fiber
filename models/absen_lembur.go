package models

import (
	"time"
)

type AbsenLembur struct {
	IdAbsenLemburd uint      `gorm:"primaryKey;column:IdAbsenHarian" json:"IdAbsenHarian"`
	IdKaryawan     string    `gorm:"column:IdKaryawan;size:255" json:"IdKaryawan"`  // VARCHAR(255)
	IdBagian       string    `gorm:"column:IdBagian;type:longtext" json:"IdBagian"` // longtext
	TglAbsen       time.Time `gorm:"column:TglAbsen;type:date" json:"TglAbsen"`     // date
	JamMasuk       time.Time `gorm:"column:JamMasuk;type:date" json:"JamMasuk"`
	JamPulang      time.Time `gorm:"column:JamPulang;type:date" json:"JamPulang"`
	JamPulangX     time.Time `gorm:"column:JamPulangX;type:date" json:"JamPulangX"`
	Proses         string    `gorm:"column:Proses;type:date" json:"Proses"`
	Shift          time.Time `gorm:"column:Shift;type:date" json:"Shift"`
	Keterangan     time.Time `gorm:"column:Keterangan;type:date" json:"Keterangan"`

	Karyawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
}

func (AbsenLembur) TableName() string {
	return "tabsen_lembur"
}
