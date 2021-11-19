package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           string `gorm:"column:id_user" json:"id_user"`
	Nama         string `gorm:"column:nama" json:"nama"`
	TglLahir     string `gorm:"column:tgl_lahir" json:"tgl_lahir"`
	Alamat       string `gorm:"column:alamat" json:"alamat"`
	JenisKelamin string `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	NoTlp        string `gorm:"column:no_telepon" json:"no_telepon"`
	Email        string `gorm:"column:email" json:"email"`
}
func (c *User) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New().String()

	return
}