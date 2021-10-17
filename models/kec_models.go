package models

type MasterDataKec struct {
	ID        string    `gorm:"column:id_kec" json:"id_kec"`
	Kecamatan  string    `gorm:"column:nama_kecamatan" json:"nama_kecamatan"`
}

func (c *MasterDataKec) TableName() string {
	return "kecamatan"
}