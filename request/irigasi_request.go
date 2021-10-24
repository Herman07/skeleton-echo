package request

type TeknikIrigasiReq struct {
	ID           string `json:"id_t_irigasi"`
	Operasi      string `json:"operasi"`
	Partisipatif string `json:"partisipatif"`
}
