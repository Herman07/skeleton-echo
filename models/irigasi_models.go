package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeknikIrigasi struct {
	ID           string `gorm:"column:id_t_irigasi" json:"id_t_irigasi"`
	Operasi      string `gorm:"column:operasi" json:"operasi"`
	Partisipatif string `gorm:"column:partisipatif" json:"partisipatif"`
}

func (c *TeknikIrigasi) TableName() string {
	return "teknik_irigasi"
}
func (c *TeknikIrigasi) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()

	return
}