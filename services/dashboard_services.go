package services

import (
	"fmt"
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
	"strings"
)

type DashboardService struct {
	DashboardRepository repository.DashboardRepository
}

func NewDashboardService(repository repository.DashboardRepository) *DashboardService {
	return &DashboardService{
		DashboardRepository: repository,
	}
}
func (s *DashboardService) QueryDatatable(searchValue string,orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.Inventaris, err error) {
	recordTotal, err = s.DashboardRepository.Count()
	strings.ToLower(searchValue)
	if searchValue != "" {
		recordFiltered, err = s.DashboardRepository.CountWhere("or", map[string]interface{}{
			"provinsi LIKE ?": "%" + searchValue + "%" ,
			"kecamatan LIKE ?":   "%" + searchValue + "%",
			"daerah LIKE ?":   "%" + searchValue + "%",
			"luas LIKE ?":   "%" + searchValue + "%",

		})

		data, err = s.DashboardRepository.FindAllWhere("or", orderType, "provinsi", limit, offset, map[string]interface{}{
			"provinsi LIKE ?": "%" + searchValue + "%",
			"kecamatan LIKE ?":   "%" + searchValue + "%",
			"daerah LIKE ?":   "%" + searchValue + "%",
			"luas LIKE ?":   "%" + searchValue + "%",

		})
		return recordTotal, recordFiltered, data, err
	}
	recordFiltered, err = s.DashboardRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.DashboardRepository.FindAllWhere("or", orderType, "created_at", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}
func (s *DashboardService) GetAll(dataReq models.Inventaris) (*models.Inventaris, error) {
	data, err := s.DashboardRepository.GetData(dataReq)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *DashboardService) FindById(id string) (*models.Inventaris, error) {
	data, err := s.DashboardRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *DashboardService) UpdateById(id string, dto request.RequestInventaris) (*models.Inventaris, error) {
	entity := models.Inventaris{
		ID:        id,
		Kecamatan: dto.Kecamatan,
		Provinsi:  dto.Provinsi,
		Daerah:    dto.Daerah,
		Luas:      dto.Luas,
	}
	fmt.Println("Isinya ", entity)

	data, err := s.DashboardRepository.UpdateById(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *DashboardService) Create(request request.RequestInventaris) (*models.Inventaris, error) {
	entity := models.Inventaris{
		Provinsi:  request.Provinsi,
		Kecamatan: request.Kecamatan,
		Daerah:    request.Daerah,
		Luas:      request.Luas,
	}
	fmt.Println("Isinya ", entity)
	data, err := s.DashboardRepository.Create(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *DashboardService) Delete(id string) error {
	entity := models.Inventaris{
		ID: id,
	}
	err := s.DashboardRepository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}
