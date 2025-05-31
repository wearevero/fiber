package models

type Karyawan struct {
	IdKaryawan       int      `gorm:"primaryKey;column:IdKaryawan" json:"IdKaryawan"`
	NikKaryawan      *string  `gorm:"column:NikKaryawan" json:"NikKaryawan"`
	TglMasuk         *string  `gorm:"column:TglMasuk" json:"TglMasuk"`
	NamaLengkap      *string  `gorm:"column:NamaLengkap" json:"NamaLengkap"`
	Jabatan          *string  `gorm:"column:Jabatan" json:"Jabatan"`
	NamaPanggilan    *string  `gorm:"column:NamaPanggilan" json:"NamaPanggilan"`
	NikKtp           *string  `gorm:"column:NikKtp" json:"NikKtp"`
	TempatLahir      *string  `gorm:"column:TempatLahir" json:"TempatLahir"`
	TglLahir         string   `gorm:"column:TglLahir" json:"TglLahir"`
	JenisKelamin     *string  `gorm:"column:JenisKelamin" json:"JenisKelamin"`
	GolDarah         *string  `gorm:"column:GolDarah" json:"GolDarah"`
	Agama            *string  `gorm:"column:Agama" json:"Agama"`
	Kewarganegaraan  *string  `gorm:"column:Kewarganegaraan" json:"Kewarganegaraan"`
	TinggiBadan      *float64 `gorm:"column:TinggiBadan" json:"TinggiBadan"`
	BeratBadan       *float64 `gorm:"column:BeratBadan" json:"BeratBadan"`
	StatusDomisili   *string  `gorm:"column:StatusDomisili" json:"StatusDomisili"`
	AlamatKtp        *string  `gorm:"column:AlamatKtp" json:"AlamatKtp"`
	KecKtp           *string  `gorm:"column:KecKtp" json:"KecKtp"`
	KabKtp           *string  `gorm:"column:KabKtp" json:"KabKtp"`
	NoHp1            *string  `gorm:"column:NoHp1" json:"NoHp1"`
	NoHp2            *string  `gorm:"column:NoHp2" json:"NoHp2"`
	Email            *string  `gorm:"column:Email" json:"Email"`
	TinggalDi        *string  `gorm:"column:TinggalDi" json:"TinggalDi"`
	AlamatTinggal    *string  `gorm:"column:AlamatTinggal" json:"AlamatTinggal"`
	StatusPerkawinan *string  `gorm:"column:StatusPerkawinan" json:"StatusPerkawinan"`
	Photo            *string  `gorm:"column:Photo" json:"Photo"`
	IdBagian         *int64   `gorm:"column:IdBagian" json:"IdBagian"`
	GajiPokok        *float64 `gorm:"column:GajiPokok" json:"GajiPokok"`
	TunPribadi       *float64 `gorm:"column:TunPribadi" json:"TunPribadi"`
	Retensi          *string  `gorm:"column:Retensi" json:"Retensi"`
	Keterangan       *string  `gorm:"column:Keterangan" json:"Keterangan"`
	Aktif            *string  `gorm:"column:Aktif" json:"Aktif"`
}

func (Karyawan) TableName() string {
	return "tkaryawan_pribadi"
}
