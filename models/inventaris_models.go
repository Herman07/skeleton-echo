package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Inventaris struct {
	ID        string    `gorm:"column:id" json:"id"`
	Provinsi  string    `gorm:"column:provinsi" json:"provinsi"`
	Kecamatan string    `gorm:"column:kecamatan" json:"kecamatan"`
	Daerah    string    `gorm:"column:daerah" json:"daerah"`
	Luas      string    `gorm:"column:luas" json:"luas"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (c *Inventaris) TableName() string {
	return "inventaris"
}

// BeforeCreate - Lifecycle callback - Generate UUID before persisting
func (c *Inventaris) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()

	return
}
