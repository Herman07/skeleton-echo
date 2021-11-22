package services

import (
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
	"strings"
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
		ID:           request.ID,
		Nama:         request.Nama,
		JenisKelamin: request.JenisKelamin,
		NoTlp:        request.NoTlp,
		Email:        request.Email,
		Alamat:       request.Alamat,
		TglLahir:     request.TglLahir,
	}
	data, err := s.UsersDataRepository.Create(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *UsersDataService) CreateAkun(request request.UsersReq, name string, user string) (*models.Akun, error) {
	entity := models.Akun{
		ID:       user,
		Username: request.Username,
		Password: request.Password,
		Foto:     name,
		TypeUser: request.TypeUser,
	}
	data, err := s.UsersDataRepository.CreateAkun(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *UsersDataService) QueryDatatable(searchValue string, orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.User, err error) {
	recordTotal, err = s.UsersDataRepository.Count()
	strings.ToLower(searchValue)
	if searchValue != "" {
		recordFiltered, err = s.UsersDataRepository.CountWhere("or", map[string]interface{}{

			"nama LIKE ?": "%" + searchValue + "%",
		})

		data, err = s.UsersDataRepository.FindAllWhere("or", orderType, "nama", limit, offset, map[string]interface{}{
			"nama LIKE ?":    "%" + searchValue + "%",
			"id_user LIKE ?": "%" + searchValue + "%",
		})
		return recordTotal, recordFiltered, data, err
	}
	recordFiltered, err = s.UsersDataRepository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.UsersDataRepository.FindAllWhere("or", orderType, "id_user", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}

func (s *UsersDataService) FindById(id string) (*models.UserUpdate, error) {
	data, err := s.UsersDataRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *UsersDataService) Delete(id string) error {
	entity := models.User{
		ID: id,
	}
	err := s.UsersDataRepository.DeleteUser(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s *UsersDataService) UpdateUser(id string, dto request.UsersReq) (*models.User, error) {
	entity := models.User{
		ID:           id,
		Nama:         dto.Nama,
		Alamat:       dto.Alamat,
		TglLahir:     dto.TglLahir,
		Email:        dto.Email,
		JenisKelamin: dto.JenisKelamin,
		NoTlp:        dto.NoTlp,
	}

	data, err := s.UsersDataRepository.UpdateUser(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *UsersDataService) UpdateAkun(id string, dto request.UsersReq, nama string) (*models.Akun, error) {

	entity := models.Akun{
		ID:       id,
		Username: dto.Username,
		Password: dto.Password,
		TypeUser: dto.TypeUser,
	}
	if nama != "" {
		entity = models.Akun{
			ID:       id,
			Username: dto.Username,
			Password: dto.Password,
			TypeUser: dto.TypeUser,
			Foto:     nama,
		}
	}

	data, err := s.UsersDataRepository.UpdateAkun(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}
