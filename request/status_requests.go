package request

type StatusLegalReq struct {
	ID                  string `json:"id_status_legal" form:"id_status_legal"`
	TahunPembentukan    string `json:"tahun_pembentukan" form:"tahun_pembentukan"`
	LamTahunPembentukan string `json:"lampiran_tahun_pembentukan" form:"lampiran_tahun_pembentukan"`
	LamKplDesa          string `json:"diket_kep_dc" form:"diket_kep_dc"`
	SKBupati            string `json:"no_sk_bupati" form:"no_sk_bupati"`
	LamSKBupati         string `json:"lampiran_sk_bupati" form:"lampiran_sk_bupati"`
	AkteNotaris         string `json:"akte_notaris" form:"akte_notaris"`
	LamAkteNotaris      string `json:"lampiran_akte_notaris" form:"lampiran_akte_notaris"`
	NoPendaftaran       string `json:"no_pendaftaran" form:"no_pendaftaran"`
	LamPendaftaran      string `json:"lampiran_pendaftaran" form:"lampiran_pendaftaran"`
}
