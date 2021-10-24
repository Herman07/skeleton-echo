package request

import "time"

type RequestInventaris struct {
	ID             string    `json:"id_p3a"`
	IDProv         string    `json:"id_prov_fk"`
	IDKab          string    `json:"id_kab_fk"`
	IDKec          string    `json:"id_kec_fk"`
	IDStatusLegal  string    `json:"id_status_legal_fk"`
	IDPengurus     string    `json:"id_kepengurusan_fk"`
	IDIrigasi      string    `json:"id_teknik_irigasi_fk"`
	IDPertanian    string    `json:"id_teknik_pertanian_fk"`
	NoUrut         string    `json:"no_urut"`
	NamaP3A        string    `json:"nama_p3a"`
	JumlahP3A      string    `json:"jumlah_p3a"`
	DaerahIrigasi  string    `json:"nama_daerah_irigasi"`
	LuasWilayah    string    `json:"luas_wilayah"`
	LuasLayananP3A string    `json:"luas_layanan_p3a"`
	Keterangan     string    `json:"keterangan"`
	CreatedAt      time.Time `json:"created_at"`
}
