package request

type KecReq struct {
	ID   string `json:"id_kec" form:"id_kec"`
	Nama string `json:"nama_kec" form:"nama_kec"`
}
