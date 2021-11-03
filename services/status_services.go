package services

import (
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
)

type StatusDataService struct {
	StatusDataRepository repository.StatusDataRepository
}

func NewStatusDataService(repository repository.StatusDataRepository) *StatusDataService {
	return &StatusDataService{
		StatusDataRepository: repository,
	}
}

func (s *StatusDataService) Create(request request.StatusLegalReq) (*models.StatusLegal, error) {
	entity := models.StatusLegal{
		ID:                  request.ID,
		TahunPembentukan:    request.TahunPembentukan,
		LamTahunPembentukan: request.LamTahunPembentukan,
		LamKplDesa:          request.LamKplDesa,
		SKBupati:            request.SKBupati,
		LamSKBupati:         request.LamSKBupati,
		AkteNotaris:         request.AkteNotaris,
		LamAkteNotaris:      request.LamAkteNotaris,
		NoPendaftaran:       request.NoPendaftaran,
		LamPendaftaran:      request.LamPendaftaran,
	}
	data, err := s.StatusDataRepository.Create(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *StatusDataService) UpdateById(id string, dto request.StatusLegalReq) (*models.StatusLegal, error) {
	entity := models.StatusLegal{
		ID:                  id,
		TahunPembentukan:    dto.TahunPembentukan,
		LamTahunPembentukan: dto.LamTahunPembentukan,
		LamKplDesa:          dto.LamKplDesa,
		SKBupati:            dto.SKBupati,
		LamSKBupati:         dto.LamSKBupati,
		AkteNotaris:         dto.AkteNotaris,
		LamAkteNotaris:      dto.LamAkteNotaris,
		NoPendaftaran:       dto.NoPendaftaran,
		LamPendaftaran:      dto.LamPendaftaran,
	}

	data, err := s.StatusDataRepository.UpdateById(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *StatusDataService) FindById(id string) (*models.StatusLegal, error) {
	data, err := s.StatusDataRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *StatusDataService) Delete(id string) error {
	entity := models.StatusLegal{
		ID: id,
	}
	err := s.StatusDataRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}
