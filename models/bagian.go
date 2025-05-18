package models

import "database/sql"

type Bagian struct {
	IdBagian      uint           `gorm:"primaryKey;column:IdBagian" json:"IdBagian"`
	NamaBagian    string         `gorm:"column:NamaBagian" json:"NamaBagian"`
	Yonma         string         `gorm:"column:Yonma" json:"Yonma"`
	UrutanBagian  string         `gorm:"column:UrutanBagian" json:"UrutanBagian"`
	NamaBagianAlt sql.NullString `gorm:"column:nama_bagian" json:"nama_bagian"`
}

func (Bagian) TableName() string {
	return "tbagian"
}
