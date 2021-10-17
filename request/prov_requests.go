package request

type ProvinsiReq struct {
	ID   string `json:"id_prov" form:"id_prov"`
	Nama string `json:"nama_prov" form:"nama_prov"`
}
