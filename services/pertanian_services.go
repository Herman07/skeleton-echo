package services

import (
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
)

type TaniDataService struct {
	TaniDataRepository repository.TaniDataRepository
}

func NewTaniDataService(repository repository.TaniDataRepository) *TaniDataService {
	return &TaniDataService{
		TaniDataRepository: repository,
	}
}

func (s *TaniDataService) Create(request request.TeknikTaniReq) (*models.TeknikPertanian, error) {
	entity := models.TeknikPertanian{
		PolaTanam:  request.PolaTanam,
		UsahaTani: request.UsahaTani,
		ID: request.ID,
	}
	data, err := s.TaniDataRepository.Create(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *TaniDataService) UpdateById(id string, dto request.TeknikTaniReq) (*models.TeknikPertanian, error) {
	entity := models.TeknikPertanian{
		ID:        id,
		PolaTanam:  dto.PolaTanam,
		UsahaTani: dto.UsahaTani,
	}

	data, err := s.TaniDataRepository.UpdateById(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *TaniDataService) FindById(id string) (*models.TeknikPertanian, error) {
	data, err := s.TaniDataRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *TaniDataService) Delete(id string) error {
	entity := models.TeknikPertanian{
		ID: id,
	}
	err := s.TaniDataRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}