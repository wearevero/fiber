package models

import "time"

type KaryawanKeluar struct {
	IdKeluar       int        `gorm:"primaryKey;column:IdKeluar" json:"IdKeluar"`
	IdKaryawan     *int       `gorm:"column:IdKaryawan;not null" json:"IdKaryawan"`
	TglPengajuan   *time.Time `gorm:"column:TglPengajuan" json:"TglPengajuan"`
	TglKeluar      *time.Time `gorm:"column:TglKeluar" json:"TglKeluar"`
	Alasan         *string    `gorm:"column:Alasan;type:varchar(255)" json:"Alasan"`
	DapatSeragam   *string    `gorm:"column:DapatSeragam;type:varchar(255)" json:"DapatSeragam"`
	KembaliSeragam *time.Time `gorm:"column:KembaliSeragam" json:"KembaliSeragam"`

	DetailKaryawan *Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
}

func (KaryawanKeluar) TableName() string {
	return "tkaryawan_keluar"
}
