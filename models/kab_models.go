package models

type MasterDataKab struct {
	ID        string `gorm:"column:id_kab" json:"id_kab"`
	IDProv    string `gorm:"column:id_provkb" json:"id_provkb"`
	Kabupaten string `gorm:"column:nama_kab" json:"nama_kab"`
}

func (c *MasterDataKab) TableName() string {
	return "kabupaten"
}
