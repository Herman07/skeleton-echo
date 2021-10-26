package request

type TeknikIrigasiReq struct {
	ID           string `json:"id_t_irigasi" form:"id_t_irigasi"`
	Operasi      string `json:"operasi" form:"operasi"`
	Partisipatif string `json:"partisipatif" form:"partisipatif"`
}
