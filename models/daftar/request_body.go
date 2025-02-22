package daftar_models

type RequestBody struct {
	Name  string `json:"nama"`
	ID    string `json:"nik"`
	Phone string `json:"no_hp"`
}
