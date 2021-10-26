package request

type TeknikTaniReq struct {
	ID        string `json:"id_t_pertanian" form:"id_t_pertanian"`
	PolaTanam string `json:"pola_tanam" form:"pola_tanam"`
	UsahaTani string `json:"usaha_tani" form:"usaha_tani"`
}
