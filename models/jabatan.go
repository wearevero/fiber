package models

type Jabatan struct {
	IdJabatan uint   `gorm:"primaryKey;column:IdJabatan" json:"IdJabatan"`
	Jabatan   string `gorm:"type:varchar(300);column:Jabatan" json:"Jabatan"`
}

func (Jabatan) TableName() string {
	return "tjabatan"
}
