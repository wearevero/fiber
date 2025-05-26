package models

type User struct {
	IdUser     uint   `gorm:"primaryKey;column:IdUser" json:"IdUser"`
	IdKaryawan string `gorm:"column:IdKaryawan" json:"IdKaryawan"`
	Username   string `gorm:"column:Username" json:"Username"`
	Password   string `gorm:"column:Password" json:"Password"`
	HakAkses   string `gorm:"column:HakAkses" json:"HakAkses"`
	Spesial    string `gorm:"column:Spesial" json:"Spesial"`

	Karyawan Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan" json:"Karyawan"`
}

func (User) TableName() string {
	return "tuser"
}
