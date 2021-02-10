package controllers

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type BaseControllerInterface interface {
}

type BaseController struct {
	db *mongo.Database
}

func NewBaseController(db *mongo.Database) *BaseController {
	return &BaseController{db: db}
}

type BaseResponseBody struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
	Message string      `json:"message"`
}

func (c BaseController) decodeRequestBody(w http.ResponseWriter, r *http.Request, req interface{}) {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (c BaseController) respond(w http.ResponseWriter, data BaseResponseBody) {
	w.WriteHeader(data.Status)
	json.NewEncoder(w).Encode(data)
}
