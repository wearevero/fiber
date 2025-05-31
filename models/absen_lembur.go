package models

type AbsenLembur struct {
	IdAbsenLembur int      `gorm:"column:IdAbsenLembur;primaryKey" json:"IdAbsenLembur"`
	IdKaryawan    int      `gorm:"column:IdKaryawan" json:"IdKaryawan"`
	IdBagian      *int     `gorm:"column:IdBagian" json:"IdBagian"`
	TglAbsen      string   `gorm:"column:TglAbsen;index" json:"TglAbsen"`
	JamMasuk      string   `gorm:"column:JamMasuk" json:"JamMasuk"`
	JamPulang     string   `gorm:"column:JamPulang" json:"JamPulang"`
	JumlahJam     *float64 `gorm:"column:JumlahJam" json:"JumlahJam"`
	JumlahJamx    string   `gorm:"column:JumlahJamx" json:"JumlahJamx"`
	Proses        string   `gorm:"column:Proses" json:"Proses"`
	Shift         string   `gorm:"column:Shift" json:"Shift"`
	Keterangan    string   `gorm:"column:Keterangan" json:"Keterangan"`

	Karyawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
}

func (AbsenLembur) TableName() string {
	return "tabsen_lembur"
}
