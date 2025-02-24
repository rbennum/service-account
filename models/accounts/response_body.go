package account_model

type ResponseBody struct {
	StatusCode   int    `json:"status_code"`
	Balance      *int   `json:"saldo,omitempty"`
	ErrorMessage string `json:"remark,omitempty"`
}
