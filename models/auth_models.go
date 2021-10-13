package models

import "time"

type Users struct {
	ID        string    `gorm:"column:id;AUTO_INCREMENT" json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	Email     string    `gorm:"column:email" json:"email"`
	TypeUser  string    `gorm:"column:type_user" json:"type_user"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (c *Users) TableName() string {
	return "users"
}
