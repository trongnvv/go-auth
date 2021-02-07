package helpers

import (
	"encoding/json"
	"net/http"
)

type BaseResponseBody struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

func DecodeRequestBody(w http.ResponseWriter, r *http.Request, req interface{}) {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func Respond(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}
