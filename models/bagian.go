package models

type Bagian struct {
	IdBagian     int    `gorm:"primaryKey;column:IdBagian" json:"IdBagian"`
	NamaBagian   string `gorm:"column:NamaBagian" json:"NamaBagian"`
	Yonma        string `gorm:"column:Yonma" json:"Yonma"`
	UrutanBagian string `gorm:"column:UrutanBagian" json:"UrutanBagian"`
}

func (Bagian) TableName() string {
	return "tbagian"
}
