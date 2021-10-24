package request

type StatusLegalReq struct {
	ID                  string `json:"id_status_legal"`
	TahunPembentukan    string `json:"tahun_pembentukan"`
	LamTahunPembentukan string `json:"lampiran_tahun_pembentukan"`
	LamKplDesa          string `json:"diket_kep_dc"`
	SKBupati            string `json:"no_sk_bupati"`
	LamSKBupati         string `json:"lampiran_sk_bupati"`
	AkteNotaris         string `json:"akte_notaris"`
	LamAkteNotaris      string `json:"lampiran_akte_notaris"`
	NoPendaftaran       string `json:"no_pendaftaran"`
	LamPendaftaran      string `json:"lampiran_pendaftaran"`
}