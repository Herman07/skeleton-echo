package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeknikPertanian struct {
	ID        string `gorm:"column:id_t_pertanian" json:"id_t_pertanian"`
	PolaTanam string `gorm:"column:pola_tanam" json:"pola_tanam"`
	UsahaTani string `gorm:"column:usaha_tani" json:"usaha_tani"`
}

func (c *TeknikPertanian) TableName() string {
	return "teknik_pertanian"
}
func (c *TeknikPertanian) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()

	return
}
