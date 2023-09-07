package response

type Response struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
}
