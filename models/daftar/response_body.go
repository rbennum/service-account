package daftar_models

type ResponseBody struct {
	StatusCode   int    `json:"status_code"`
	Account      string `json:"no_rekening,omitempty"`
	ErrorMessage string `json:"remark,omitempty"`
}
