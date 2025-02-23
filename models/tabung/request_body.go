package tabung_model

type RequestBody struct {
	AccountNumber string `json:"no_rekening"`
	Transferred   int    `json:"nominal"`
}
