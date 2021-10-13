package services

import (
	"skeleton-echo/models"
	"skeleton-echo/repository"
)

type AuthService struct {
	authRepository repository.AuthRepository
}

func NewAuthService(repositoris repository.AuthRepository) *AuthService {
	return &AuthService{
		authRepository: repositoris,
	}
}
func (s *AuthService) Login(username string) (*models.Users, error) {
	data, err := s.authRepository.Login(username)
	if err != nil{
		return nil, err
	} else {
	}
	return data, err
}

