package request

import "time"

type RequestInventaris struct {
	Provinsi  string    `json:"provinsi"`
	Kecamatan string    `json:"kecamatan"`
	Daerah    string    `json:"daerah"`
	Luas      string    `json:"luas"`
	CreatedAt time.Time `json:"created_at"`
}
