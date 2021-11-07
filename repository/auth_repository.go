package repository

import (
	"gorm.io/gorm"
	"skeleton-echo/models"
	"skeleton-echo/request"
)

type AuthRepository interface {
	Login( req request.LoginRequest) (*models.Users, error)
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

func (r *authRepository)Login( req request.LoginRequest) (*models.Users,error) {
	var entity models.Users
	err := r.DB.Table("user_group").Where("username = ? AND password = ?", req.Username, req.Password).First(&entity).Error

	return &entity, err

}

func (r *authRepository) DbInstance() *gorm.DB {
	return r.DB
}
