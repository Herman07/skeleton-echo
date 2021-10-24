package request

type TeknikTaniReq struct {
	ID        string `json:"id_t_pertanian"`
	PolaTanam string `json:"pola_tanam"`
	UsahaTani string `json:"usaha_tani"`
}
