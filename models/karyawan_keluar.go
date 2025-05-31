package models

import "time"

type KaryawanKeluar struct {
	IdKeluar       int       `gorm:"primaryKey;column:IdKeluar" json:"IdKeluar"`
	IdKaryawan     int       `gorm:"type:int;column:IdKaryawan" json:"IdKaryawan"`
	TglPengajuan   time.Time `gorm:"type:timestamp;column:TglPengajuan" json:"TglPengajuan"`
	TglKeluar      time.Time `gorm:"type:timestamp;column:TglKeluar" json:"TglKeluar"`
	Alasan         string    `gorm:"type:varchar(255);column:Alasan" json:"Alasan"`
	DapatSeragam   string    `gorm:"type:varchar(255);column:DapatSeragam" json:"DapatSeragam"`
	KembaliSeragam time.Time `gorm:"type:timestamp;column:KembaliSeragam" json:"KembaliSeragam"`

	Karyawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
}

func (KaryawanKeluar) TableName() string {
	return "tkaryawan_keluar"
}
