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
func (s *DashboardService) Create1(request request.RequestInventaris) (*models.Pengurus, error) {
	entity := models.Pengurus{
		ID: request.IDPengurus,
		Ketua: request.Ketua,
		Wakil: request.Wakil,
		Bendahara: request.Bendahara,
		Sekretaris: request.Sekretaris,
		SekTeknik: request.SekTeknik,
		SekBisnis: request.SekBisnis,
		SekOP: request.SekOP,
		JumlahAnggota: request.JumlahAnggota,
		Sekretariat: request.Sekretariat,
		LampiranSekretariat: request.LampiranSekretariat,
		ArealTersier: request.ArealTersier,
		NoADRT: request.NoADRT,
		LampiranADRT: request.LampiranADRT,
		PresentasiPerempuanP3A: request.PresentasiPerempuanP3A,
		PengisianBuku: request.PengisianBuku,
		Iuran: request.Iuran,

	}
	data, err := s.DashboardRepository.Create2(entity)
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
func (s *DashboardService) Create2(request request.RequestInventaris) (*models.StatusLegal, error) {
	entity := models.StatusLegal{
		ID: request.IDStatus,
		TahunPembentukan: request.TahunPembentukan,
		LamTahunPembentukan: request.LamTahunPembentukan,
		LamKplDesa: request.LamKplDesa,
		SKBupati: request.SKBupati,
		LamSKBupati: request.LamSKBupati,
		AkteNotaris: request.AkteNotaris,
		LamAkteNotaris: request.LamAkteNotaris,
		NoPendaftaran: request.NoPendaftaran,
		LamPendaftaran: request.LamPendaftaran,
	}

	data, err := s.DashboardRepository.Create3(entity)
	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *DashboardService) Create3(request request.RequestInventaris) (*models.TeknikIrigasi, error) {
	entity := models.TeknikIrigasi{
		ID: request.IDIrigasi,
		Operasi: request.Operasi,
		Partisipatif: request.Partisipatif,
	}
	data, err := s.DashboardRepository.Create4(entity)
	if err != nil {
		return nil, err
	}
	return data, err
}
func (s *DashboardService) Create4(request request.RequestInventaris) (*models.TeknikPertanian, error) {
	entity := models.TeknikPertanian{
		ID: request.IDPertanian,
		PolaTanam: request.PolaTanam,
		UsahaTani: request.UsahaTani,
	}
	data, err := s.DashboardRepository.Create5(entity)
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
