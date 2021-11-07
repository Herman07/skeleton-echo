package request

type PengurusReq struct {
	Ketua                  string `gorm:"column:ketua" json:"ketua"`
	Wakil                  string `gorm:"column:wakil" json:"wakil"`
	Sekretaris             string `gorm:"column:sekretaris" json:"sekretaris"`
	Bendahara              string `gorm:"column:bendahara" json:"bendahara"`
	SekTeknik              string `gorm:"column:sek_teknik" json:"sek_teknik"`
	SekOP                  string `gorm:"column:sek_op" json:"sek_op"`
	SekBisnis              string `gorm:"column:sek_bisnis" json:"sek_bisnis"`
	JumlahAnggota          string `gorm:"column:jumlah_anggota" json:"jumlah_anggota"`
	NoADRT                 string `gorm:"column:no_ad_rt" json:"no_ad_rt"`
	LampiranADRT           string `gorm:"column:lampiran_ad_rt" json:"lampiran_ad_rt"`
	Sekretariat            string `gorm:"column:sekretariat" json:"sekretariat"`
	LampiranSekretariat    string `gorm:"column:lampiran_sekretariat" json:"lampiran_sekretariat"`
	PresentasiPerempuanP3A string `gorm:"column:presentase_perempuan_p3a" json:"presentase_perempuan_p3a"`
	ArealTersier           string `gorm:"column:area_tersier" json:"areal_tersier"`
	PengisianBuku          string `gorm:"column:pengisian_buku" json:"pengisian_buku"`
	Iuran                  string `gorm:"column:iuran" json:"iuran"`
}