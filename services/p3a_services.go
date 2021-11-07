package services

import (
	"skeleton-echo/models"
	"skeleton-echo/repository"
	"skeleton-echo/request"
	"strings"
)

type P3Service struct {
	P3Repository repository.P3Repository
}

func NewP3Service(repository repository.P3Repository) *P3Service {
	return &P3Service{
		P3Repository: repository,
	}
}
func (s *P3Service) QueryDatatable(searchValue string, orderType string, orderBy string, limit int, offset int) (
	recordTotal int64, recordFiltered int64, data []models.Inventaris, err error) {
	recordTotal, err = s.P3Repository.Count()
	strings.ToLower(searchValue)
	if searchValue != "" {
		recordFiltered, err = s.P3Repository.CountWhere("or", map[string]interface{}{
			"ID LIKE ?":             "%" + searchValue + "%",
			"NomorUrut LIKE ?":      "%" + searchValue + "%",
			"NamaP3A LIKE ?":        "%" + searchValue + "%",
			"JumlahP3A LIKE ?":      "%" + searchValue + "%",
			"DaerahIrigasi LIKE ?":  "%" + searchValue + "%",
			"LuasWilayah LIKE ?":    "%" + searchValue + "%",
			"LuasLayananP3A LIKE ?": "%" + searchValue + "%",
			"Keterangan LIKE ?":     "%" + searchValue + "%",
		})

		data, err = s.P3Repository.FindAllWhere("or", orderType, "NomorUrut", limit, offset, map[string]interface{}{
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
	recordFiltered, err = s.P3Repository.CountWhere("or", map[string]interface{}{
		"1 =?": 1,
	})

	data, err = s.P3Repository.FindAllWhere("or", orderType, "id_p3a", limit, offset, map[string]interface{}{
		"1= ?": 1,
	})
	return recordTotal, recordFiltered, data, err
}
//func (s *P3Service) GetAll(dataReq models.Inventaris) (*models.Inventaris, error) {
//	data, err := s.P3Repository.GetData(dataReq)
//	if err != nil {
//		return nil, err
//	}
//	return data, err
//}

//func (s *P3Service) FindById(id string) (*models.Inventaris, error) {
//	data, err := s.P3Repository.FindById(id)
//	if err != nil {
//		return nil, err
//	}
//	return data, err
//}

//func (s *P3Service) UpdateById(id string, dto request.RequestInventaris) (*models.Inventaris, error) {
//	entity := models.Inventaris{
//		ID:             id,
//		NoUrut:         dto.NoUrut,
//		NamaP3A:        dto.NamaP3A,
//		JumlahP3A:      dto.JumlahP3A,
//		DaerahIrigasi:  dto.DaerahIrigasi,
//		LuasWilayah:    dto.LuasWilayah,
//		LuasLayananP3A: dto.LuasLayananP3A,
//	}
//
//	data, err := s.P3Repository.UpdateById(entity)
//
//	if err != nil {
//		return nil, err
//	}
//	return data, err
//}

func (s *P3Service) CreateStatusLegal(request request.RequestInventaris) (*models.StatusLegal, error) {
	entity := models.StatusLegal{
		ID:                  request.IDStatusLegal,
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
	data, err := s.P3Repository.CreateStatusLegal(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}


func (s *P3Service) CreatePengurus(request request.RequestInventaris) (*models.Pengurus, error) {
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

	data, err := s.P3Repository.CreatePengurus(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) CreateIrigasi(request request.RequestInventaris) (*models.TeknikIrigasi, error) {
	entity := models.TeknikIrigasi{
		Operasi:      request.Operasi,
		Partisipatif: request.Partisipatif,
		ID:           request.IDIrigasi,
	}
	data, err := s.P3Repository.CreateIrigasi(entity)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) CreatePertanian(request request.RequestInventaris) (*models.TeknikPertanian, error) {
	entity := models.TeknikPertanian{
		PolaTanam:  request.PolaTanam,
		UsahaTani: request.UsahaTani,
		ID: request.IDPertanian,
	}
	data, err := s.P3Repository.CreatePertanian(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) CreateDataP3a(request request.RequestInventaris,idStatusLegal string, idPengurus string, idIrigasi string, idPertanian string) (*models.Inventaris, error) {
	entity := models.Inventaris{
		IDProv:         request.IDProv,
		IDKab:          request.IDKab,
		IDKec:          request.IDKec,
		IDStatusLegal:  idStatusLegal,
		IDPengurus:     idPengurus,
		IDIrigasi:      idIrigasi,
		IDPertanian:    idPertanian,
		NoUrut:         request.NoUrut,
		NamaP3A:        request.NamaP3A,
		JumlahP3A:      request.JumlahP3A,
		DaerahIrigasi:  request.DaerahIrigasi,
		LuasWilayah:    request.LuasWilayah,
		LuasLayananP3A: request.LuasLayananP3A,
		Keterangan: request.Keterangan,
	}
	data, err := s.P3Repository.Create(entity)
	if err != nil {
		return nil, err
	}
	return data, err
}

//func (s *P3Service) Delete(id string) error {
//	entity := models.Inventaris{
//		ID: id,
//	}
//	err := s.P3Repository.Delete(entity)
//	if err != nil {
//		return err
//	} else {
//		return nil
//	}
//}
