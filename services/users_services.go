package services

import (
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
)

type UsersDataService struct {
	UsersDataRepository repository.UsersDataRepository
}

func NewUsersDataService(repository repository.UsersDataRepository) *UsersDataService {
	return &UsersDataService{
		UsersDataRepository: repository,
	}
}

func (s *UsersDataService) Create(request request.UsersReq) (*models.User, error) {
	entity := models.User{
		ID: request.ID,
		Nama: request.Nama,
		JenisKelamin: request.JenisKelamin,
		NoTlp: request.NoTlp,
		Email: request.Email,
		Alamat: request.Alamat,
		TglLahir: request.TglLahir,
	}
	data, err := s.UsersDataRepository.Create(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *UsersDataService) CreateAkun(request request.UsersReq, name string,user string) (*models.Akun, error) {
	entity := models.Akun{
		ID: user,
		Username: request.Username,
		Password: request.Password,
		Foto: name,
		TypeUser: request.TypeUser,
	}
	data, err := s.UsersDataRepository.CreateAkun(entity)

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