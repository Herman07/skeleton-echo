package services

import (
	"skeleton-echo/models"
	"skeleton-echo/repository"
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

func (s *StatusDataService) QueryDatatable(searchValue string,orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.StatusLegal, err error) {
	recordTotal, err = s.StatusDataRepository.Count()
	strings.ToLower(searchValue)
	if searchValue != "" {
		recordFiltered, err = s.StatusDataRepository.CountWhere("or", map[string]interface{}{

			"id_status_legal LIKE ?": "%" + searchValue + "%" ,
			"tahun_pembentukan LIKE ?": "%" + searchValue + "%" ,
			"sk_bupati LIKE ?": "%" + searchValue + "%" ,
			"akte_notaris LIKE ?": "%" + searchValue + "%" ,
			"no_pendaftaran LIKE ?": "%" + searchValue + "%" ,

		})

		data, err = s.StatusDataRepository.FindAllWhere("or", orderType, "id_status_legal", limit, offset, map[string]interface{}{
			"id_status_legal LIKE ?": "%" + searchValue + "%" ,
			"tahun_pembentukan LIKE ?": "%" + searchValue + "%" ,
			"sk_bupati LIKE ?": "%" + searchValue + "%" ,
			"akte_notaris LIKE ?": "%" + searchValue + "%" ,
			"no_pendaftaran LIKE ?": "%" + searchValue + "%" ,

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

//func (s *MasterDataService) Create(request request.ProvinsiReq) (*models.MasterDataProvinsi, error) {
//	entity := models.MasterDataProvinsi{
//		Provinsi:  request.Nama,
//		ID: request.ID,
//	}
//	fmt.Println("Isinya ", entity)
//	data, err := s.MasterDataRepository.Create(entity)
//
//	if err != nil {
//		return nil, err
//	}
//	return data, err
//}
//func (s *MasterDataService) UpdateById(id string, dto request.ProvinsiReq) (*models.MasterDataProvinsi, error) {
//	entity := models.MasterDataProvinsi{
//		ID:        id,
//		Provinsi:  dto.Nama,
//	}
//
//	data, err := s.MasterDataRepository.UpdateById(entity)
//
//	if err != nil {
//		return nil, err
//	}
//	return data, err
//}
func (s *StatusDataService) FindById(id string) (*models.StatusLegal, error) {
	data, err := s.StatusDataRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return data, err
}
//
//func (s *MasterDataService) Delete(id string) error {
//	entity := models.MasterDataProvinsi{
//		ID: id,
//	}
//	err := s.MasterDataRepository.Delete(entity)
//	if err != nil {
//		return err
//	} else {
//		return nil
//	}
//}