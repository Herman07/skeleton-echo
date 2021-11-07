package services

import (
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
)

type AuthService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(repositoris repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: repositoris,
	}
}
func (s *AuthService) Login(request request.LoginRequest) (*models.Users, error) {
	data, err := s.authRepository.Login(request)
	if err != nil{
		return nil, err
	} else {
	}
	return data, err
}

