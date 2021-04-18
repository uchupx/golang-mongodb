package reqres

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	StatusCode int         `json:"status_code"`
	Error      interface{} `json:"error"`
}

type BodyRespone struct {
	ContentType *string
	StatusCode  int
	Body        interface{}
}

func returnResponse(w http.ResponseWriter, bodyRespone BodyRespone) {
	w.Header().Set("Content-Type", "application/json")

	if bodyRespone.ContentType != nil {
		w.Header().Set("Content-Type", *bodyRespone.ContentType)
	}

	w.WriteHeader(bodyRespone.StatusCode)
	json.NewEncoder(w).Encode(bodyRespone.Body)
}
