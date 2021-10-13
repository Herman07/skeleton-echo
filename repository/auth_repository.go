package repository

import (
	"gorm.io/gorm"
	"skeleton-echo/models"
)

type AuthRepository interface {
	Login( username string) (*models.Users, error)
	DbInstance() *gorm.DB
}

type authRepository struct {
	*gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{
		DB: db,
	}
}

func (r *authRepository)Login( username string) (*models.Users,error) {
	var entity models.Users
	err := r.DB.Table("users").Where("username = ?", username).First(&entity).Error

	return &entity, err

}

func (r *authRepository) DbInstance() *gorm.DB {
	return r.DB
}
