package services

import (
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
)

type PengurusDataService struct {
	PengurusDataRepository repository.PengurusDataRepository
}

func NewPengurusDataService(repository repository.PengurusDataRepository) *PengurusDataService {
	return &PengurusDataService{
		PengurusDataRepository: repository,
	}
}

func (s *PengurusDataService) Create(request request.PengurusReq) (*models.Pengurus, error) {
	entity := models.Pengurus{
		Ketua:                  request.Ketua,
		Wakil:                  request.Wakil,
		Sekretaris:             request.Sekretaris,
		Bendahara:              request.Bendahara,
		SekTeknik:              request.SekTeknik,
		SekOP:                  request.SekOP,
		SekBisnis:              request.SekBisnis,
		JumlahAnggota:          request.JumlahAnggota,
		NoADRT:                 request.NoADRT,
		LampiranADRT:           request.LampiranADRT,
		Sekretariat:            request.Sekretariat,
		LampiranSekretariat:    request.LampiranSekretariat,
		PresentasiPerempuanP3A: request.PresentasiPerempuanP3A,
		ArealTersier:           request.ArealTersier,
		PengisianBuku:          request.PengisianBuku,
		Iuran:                  request.Iuran,
	}

	data, err := s.PengurusDataRepository.Create(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
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