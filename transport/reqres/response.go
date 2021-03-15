package reqres

type ErrorMessage struct {
	StatusCode int         `json:"status_code"`
	Error      interface{} `json:"error"`
}
