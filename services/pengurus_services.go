package services

import (
	"skeleton-echo/repository"
)

type PengurusDataService struct {
	PengurusDataRepository repository.PengurusDataRepository
}

func NewPengurusDataService(repository repository.PengurusDataRepository) *PengurusDataService {
	return &PengurusDataService{
		PengurusDataRepository: repository,
	}
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
//func (s *MasterDataService) FindById(id string) (*models.MasterDataProvinsi, error) {
//	data, err := s.MasterDataRepository.FindById(id)
//	if err != nil {
//		return nil, err
//	}
//	return data, err
//}
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