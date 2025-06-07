package models

import (
	"time"
)

type AbsenIjin struct {
	IdIjinKeluar int       `gorm:"column:IdIjinKeluar;primaryKey" json:"IdIjinKeluar"`
	IdKaryawan   *int      `gorm:"column:IdKaryawan" json:"IdKaryawan"`
	IdBagian     *int      `gorm:"column:IdBagian" json:"IdBagian"`
	TglIjin      time.Time `gorm:"column:TglIjin" json:"TglIjin"`
	JamMasuk     string    `gorm:"column:JamMasuk" json:"JamMasuk"`
	JamKeluar    string    `gorm:"column:JamPulang" json:"JamKeluar"`
	JumIjin      *float64  `gorm:"column:JumIjin" json:"JumIjinX"`
	JumIjinX     string    `gorm:"column:JumJamX" json:"JumJamX"`
	Proses       string    `gorm:"column:Proses" json:"Proses"`

	DetailKaryawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
	Bagian         Bagian   `gorm:"foreignKey:IdBagian;references:IdBagian" json:"Bagian"`
}

func (AbsenIjin) TableName() string {
	return "tabsen_ijinkeluar"
}
