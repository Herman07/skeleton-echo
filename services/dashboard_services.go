package services

import (
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
func (s *DashboardService) QueryDatatable(searchValue string, orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.Inventaris, err error) {
	recordTotal, err = s.DashboardRepository.Count()
	strings.ToLower(searchValue)
	if searchValue != "" {
		recordFiltered, err = s.DashboardRepository.CountWhere("or", map[string]interface{}{
			"ID LIKE ?":             "%" + searchValue + "%",
			"NomorUrut LIKE ?":      "%" + searchValue + "%",
			"NamaP3A LIKE ?":        "%" + searchValue + "%",
			"JumlahP3A LIKE ?":      "%" + searchValue + "%",
			"DaerahIrigasi LIKE ?":  "%" + searchValue + "%",
			"LuasWilayah LIKE ?":    "%" + searchValue + "%",
			"LuasLayananP3A LIKE ?": "%" + searchValue + "%",
			"Keterangan LIKE ?":     "%" + searchValue + "%",
		})

		data, err = s.DashboardRepository.FindAllWhere("or", orderType, "NomorUrut", limit, offset, map[string]interface{}{
			"ID LIKE ?":             "%" + searchValue + "%",
			"NomorUrut LIKE ?":      "%" + searchValue + "%",
			"NamaP3A LIKE ?":        "%" + searchValue + "%",
			"JumlahP3A LIKE ?":      "%" + searchValue + "%",
			"DaerahIrigasi LIKE ?":  "%" + searchValue + "%",
			"LuasWilayah LIKE ?":    "%" + searchValue + "%",
			"LuasLayananP3A LIKE ?": "%" + searchValue + "%",
			"Keterangan LIKE ?":     "%" + searchValue + "%",
		})
		return recordTotal, recordFiltered, data, err
	}
	recordFiltered, err = s.DashboardRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.DashboardRepository.FindAllWhere("or", orderType, "id_p3a", limit, offset, map[string]interface{}{
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
		ID:             id,
		NoUrut:         dto.NoUrut,
		NamaP3A:        dto.NamaP3A,
		JumlahP3A:      dto.JumlahP3A,
		DaerahIrigasi:  dto.DaerahIrigasi,
		LuasWilayah:    dto.LuasWilayah,
		LuasLayananP3A: dto.LuasLayananP3A,
	}

	data, err := s.DashboardRepository.UpdateById(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *DashboardService) Create(request request.RequestInventaris) (*models.Inventaris, error) {
	entity := models.Inventaris{
		IDProv:         request.IDProv,
		IDKab:          request.IDKab,
		IDKec:          request.IDKec,
		IDStatusLegal:  request.IDStatusLegal,
		IDPengurus:     request.IDPengurus,
		IDIrigasi:      request.IDIrigasi,
		IDPertanian:    request.IDPertanian,
		NoUrut:         request.NoUrut,
		NamaP3A:        request.NamaP3A,
		JumlahP3A:      request.JumlahP3A,
		DaerahIrigasi:  request.DaerahIrigasi,
		LuasWilayah:    request.LuasWilayah,
		LuasLayananP3A: request.LuasLayananP3A,
	}
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
