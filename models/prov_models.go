package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MasterDataProvinsi struct {
	ID        string    `gorm:"column:id_prov" json:"id_prov"`
	Provinsi  string    `gorm:"column:nama_prov" json:"nama_provinsi"`
}

func (c *MasterDataProvinsi) TableName() string {
	return "provinsi"
}

func (c *MasterDataProvinsi) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()
	return
}
