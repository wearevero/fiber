package models

type User struct {
	IdUser     int     `gorm:"primaryKey;column:IdUser" json:"IdUser"`
	IdKaryawan *int    `gorm:"column:IdKaryawan" json:"IdKaryawan"`
	Username   *string `gorm:"column:Username" json:"Username"`
	Password   *string `gorm:"column:Password" json:"Password"`
	HakAkses   *string `gorm:"column:HakAkses" json:"HakAkses"`
	Spesial    *string `gorm:"column:Spesial" json:"Spesial"`

	DetailKaryawan *Karyawan `gorm:"foreignKey:IdKaryawan;references:IdKaryawan;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"Karyawan"`
}

func (User) TableName() string {
	return "tuser"
}
