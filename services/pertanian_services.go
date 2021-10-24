package services

import (
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
	"strings"
)

type TaniDataService struct {
	TaniDataRepository repository.TaniDataRepository
}

func NewTaniDataService(repository repository.TaniDataRepository) *TaniDataService {
	return &TaniDataService{
		TaniDataRepository: repository,
	}
}

func (s *TaniDataService) QueryDatatable(searchValue string,orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.TeknikPertanian, err error) {
	recordTotal, err = s.TaniDataRepository.Count()
	strings.ToLower(searchValue)
	if searchValue != "" {
		recordFiltered, err = s.TaniDataRepository.CountWhere("or", map[string]interface{}{
			"pola_tanam LIKE ?": "%" + searchValue + "%",
			"usaha_tani LIKE ?": "%" + searchValue + "%",
			"id_t_pertanian LIKE ?": "%" + searchValue + "%",
		})

		data, err = s.TaniDataRepository.FindAllWhere("or", orderType, "id_t_pertanian", limit, offset, map[string]interface{}{
			"pola_tanam LIKE ?": "%" + searchValue + "%",
			"usaha_tani LIKE ?": "%" + searchValue + "%",
			"id_t_pertanian LIKE ?": "%" + searchValue + "%",

		})
		return recordTotal, recordFiltered, data, err
	}
	recordFiltered, err = s.TaniDataRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.TaniDataRepository.FindAllWhere("or", orderType, "id_t_pertanian", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
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