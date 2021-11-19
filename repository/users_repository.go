package repository

import (
	"gorm.io/gorm"
	"skeleton-echo/models"
)

type UsersDataRepository interface {
	Create(entity models.User) (*models.User, error)
	CreateAkun(entity models.Akun) (*models.Akun, error)
	DbInstance() *gorm.DB
}


type usersdataRepository struct {
	*gorm.DB
}

func NewUsersDataRepository(db *gorm.DB) UsersDataRepository {
	return &usersdataRepository{
		DB: db,
	}
}
func (r usersdataRepository) Create(entity models.User) (*models.User, error) {
	err := r.DB.Table("user").Create(&entity).Error
	return &entity, err
}
func (r usersdataRepository) CreateAkun(entity models.Akun) (*models.Akun, error) {
	err := r.DB.Table("user_group").Create(&entity).Error
	return &entity, err
}
func (r *usersdataRepository) DbInstance() *gorm.DB {
	return r.DB
}

