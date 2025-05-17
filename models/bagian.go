package models

type Bagian struct {
	IdBagian   int64  `gorm:"primaryKey;column:IdBagian" json:"IdBagian"`
	NamaBagian string `gorm:"type:varchar(300);column:NamaBagian" json:"NamaBagian"`
	Yonma      string `gorm:"type:text;column:Yonma" json:"Yonma"`
}

func (Bagian) TableName() string {
	return "tbagian"
}
