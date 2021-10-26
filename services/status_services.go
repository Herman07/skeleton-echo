package services

import (
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
	"strings"
)

type StatusDataService struct {
	StatusDataRepository repository.StatusDataRepository
}

func NewStatusDataService(repository repository.StatusDataRepository) *StatusDataService {
	return &StatusDataService{
		StatusDataRepository: repository,
	}
}

func (s *StatusDataService) QueryDatatable(searchValue string, orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.StatusLegal, err error) {
	recordTotal, err = s.StatusDataRepository.Count()
	strings.ToLower(searchValue)
	if searchValue != "" {
		recordFiltered, err = s.StatusDataRepository.CountWhere("or", map[string]interface{}{

			"id_status_legal LIKE ?":   "%" + searchValue + "%",
			"tahun_pembentukan LIKE ?": "%" + searchValue + "%",
			"sk_bupati LIKE ?":         "%" + searchValue + "%",
			"akte_notaris LIKE ?":      "%" + searchValue + "%",
			"no_pendaftaran LIKE ?":    "%" + searchValue + "%",
		})

		data, err = s.StatusDataRepository.FindAllWhere("or", orderType, "id_status_legal", limit, offset, map[string]interface{}{
			"id_status_legal LIKE ?":   "%" + searchValue + "%",
			"tahun_pembentukan LIKE ?": "%" + searchValue + "%",
			"sk_bupati LIKE ?":         "%" + searchValue + "%",
			"akte_notaris LIKE ?":      "%" + searchValue + "%",
			"no_pendaftaran LIKE ?":    "%" + searchValue + "%",
		})
		return recordTotal, recordFiltered, data, err
	}
	recordFiltered, err = s.StatusDataRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.StatusDataRepository.FindAllWhere("or", orderType, "id_status_legal", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
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
