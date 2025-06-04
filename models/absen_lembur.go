package models

import "time"

type AbsenLembur struct {
	IdAbsenLembur int       `gorm:"column:IdAbsenLembur;primaryKey" json:"IdAbsenLembur"`
	IdKaryawan    *int      `gorm:"column:IdKaryawan" json:"IdKaryawan"`
	IdBagian      *int      `gorm:"column:IdBagian" json:"IdBagian"`
	TglAbsen      time.Time `gorm:"column:TglAbsen;index" json:"TglAbsen"`
	JamMasuk      string    `gorm:"column:JamMasuk" json:"JamMasuk"`
	JamPulang     string    `gorm:"column:JamPulang" json:"JamPulang"`
	JumlahJam     *float64  `gorm:"column:JumlahJam" json:"JumlahJam"`
	JumlahJamx    string    `gorm:"column:JumlahJamx" json:"JumlahJamx"`
	Proses        string    `gorm:"column:Proses" json:"Proses"`
	Shift         string    `gorm:"column:Shift" json:"Shift"`
	Keterangan    string    `gorm:"column:Keterangan" json:"Keterangan"`

	DetailKaryawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
	Bagian         Bagian   `gorm:"foreignKey:IdBagian;references:IdBagian" json:"Bagian"`
}

func (AbsenLembur) TableName() string {
	return "tabsen_lembur"
}
