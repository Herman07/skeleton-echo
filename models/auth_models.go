package models

type Users struct {
	ID       string `gorm:"column:id_usergroup;AUTO_INCREMENT" json:"id_usergroup"`
	Username string `gorm:"column:username" json:"username"`
	Password string `gorm:"column:password" json:"password"`
	Foto     string `gorm:"column:foto" json:"foto"`
	TypeUser string `gorm:"column:user_type" json:"user_type"`
}

func (c *Users) TableName() string {
	return "user_group"
}
