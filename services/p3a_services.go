package services

import (
	"fmt"
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

func (s *P3Service) FindById(id string) (*models.P3AModels, error) {
	data, err := s.P3Repository.FindById(id)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) UpdateById(id string, dto request.UpdateInventaris) (*models.Inventaris, error) {
	entity := models.Inventaris{
		ID:             id,
		NoUrut:         dto.NoUrut,
		NamaP3A:        dto.NamaP3A,
		JumlahP3A:      dto.JumlahP3A,
		DaerahIrigasi:  dto.DaerahIrigasi,
		LuasWilayah:    dto.LuasWilayah,
		LuasLayananP3A: dto.LuasLayananP3A,
	}

	data, err := s.P3Repository.UpdateById(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) CreateStatusLegal(request request.RequestInventaris, namaFile []string) (*models.StatusLegal, error) {
	entity := models.StatusLegal{
		TahunPembentukan:    request.TahunPembentukan,
		LamTahunPembentukan: namaFile[0],
		LamKplDesa:          namaFile[1],
		SKBupati:            request.SKBupati,
		LamSKBupati:         namaFile[2],
		AkteNotaris:         request.AkteNotaris,
		LamAkteNotaris:      namaFile[3],
		NoPendaftaran:       request.NoPendaftaran,
		LamPendaftaran:      namaFile[4],
	}
	data, err := s.P3Repository.CreateStatusLegal(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) CreatePengurus(request request.RequestInventaris, namaFile []string) (*models.Pengurus, error) {
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
		LampiranADRT:           namaFile[5],
		Sekretariat:            request.Sekretariat,
		LampiranSekretariat:    namaFile[6],
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
		PolaTanam: request.PolaTanam,
		UsahaTani: request.UsahaTani,
		ID:        request.IDPertanian,
	}
	data, err := s.P3Repository.CreatePertanian(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) CreateDataP3a(request request.RequestInventaris, idStatusLegal string, idPengurus string, idIrigasi string, idPertanian string) (*models.Inventaris, error) {
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
		Keterangan:     request.Keterangan,
	}
	data, err := s.P3Repository.Create(entity)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) Delete(id string) error {
	data, err := s.P3Repository.FindById(id)

	// Delete Status Legal
	err = s.P3Repository.DeleteStatusLegal(data.IDStatus)

	// Delete Irigasi
	err = s.P3Repository.DeleteIrigasi(data.IDIrig)

	// Delete Pengurusan
	err = s.P3Repository.DeletePengurusan(data.IDPengurusan)

	// Delete Pertanian
	err = s.P3Repository.DeletePertanian(data.IDTani)



	entity := models.Inventaris{
		ID: id,
	}
	err = s.P3Repository.Delete(entity)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (s *P3Service) UpdateStatusLegal(id string, dto request.UpdateInventaris) (*models.StatusLegal, error) {
	entity := models.StatusLegal{
		ID:                  id,
		TahunPembentukan:    dto.TahunPembentukan,
		LamTahunPembentukan: *dto.LamTahunPembentukan,
		LamKplDesa:          *dto.LamKplDesa,
		SKBupati:            dto.SKBupati,
		LamSKBupati:         *dto.LamSKBupati,
		AkteNotaris:         dto.AkteNotaris,
		LamAkteNotaris:      *dto.LamAkteNotaris,
		NoPendaftaran:       dto.NoPendaftaran,
		LamPendaftaran:      *dto.LamPendaftaran,
	}
	fmt.Println("Entity Services : ", entity.ID)
	data, err := s.P3Repository.UpdateStatusLegal(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) UpdatePengurus(id string, dto request.UpdateInventaris) (*models.Pengurus, error) {
	entity := models.Pengurus{
		ID:                     id,
		Ketua:                  dto.Ketua,
		Wakil:                  dto.Wakil,
		Sekretaris:             dto.Sekretaris,
		Bendahara:              dto.Bendahara,
		SekTeknik:              dto.SekTeknik,
		SekOP:                  dto.SekOP,
		SekBisnis:              dto.SekBisnis,
		JumlahAnggota:          dto.JumlahAnggota,
		NoADRT:                 dto.NoADRT,
		LampiranADRT:           *dto.LampiranADRT,
		Sekretariat:            dto.Sekretariat,
		LampiranSekretariat:    *dto.LampiranSekretariat,
		PresentasiPerempuanP3A: dto.PresentasiPerempuanP3A,
		ArealTersier:           dto.ArealTersier,
		PengisianBuku:          dto.PengisianBuku,
		Iuran:                  dto.Iuran,
	}

	data, err := s.P3Repository.UpdatePengurus(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) UpdateIrigasi(id string, dto request.UpdateInventaris) (*models.TeknikIrigasi, error) {
	entity := models.TeknikIrigasi{
		ID:           id,
		Operasi:      dto.Operasi,
		Partisipatif: dto.Partisipatif,
	}
	data, err := s.P3Repository.UpdateIrigasi(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) UpdatePertanian(id string, dto request.UpdateInventaris) (*models.TeknikPertanian, error) {
	entity := models.TeknikPertanian{
		ID:        id,
		PolaTanam: dto.PolaTanam,
		UsahaTani: dto.UsahaTani,
	}

	data, err := s.P3Repository.UpdatePertanian(entity)

	if err != nil {
		return nil, err
	}
	return data, err
}

func (s *P3Service) GetDataExport() ([]models.P3AModels, error) {
	data, err := s.P3Repository.ExportExcel()
	if err != nil {
		return nil, err
	}
	return *data, nil
}