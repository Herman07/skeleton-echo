package models

type MasterDataProvinsi struct {
	ID        string    `gorm:"column:id_prov" json:"id_prov"`
	Provinsi  string    `gorm:"column:nama_prov" json:"nama_provinsi"`
}

func (c *MasterDataProvinsi) TableName() string {
	return "provinsi"
}

