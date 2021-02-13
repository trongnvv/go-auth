package controllers

import (
	"encoding/json"
	"gostart/models"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (c BaseController) decodeRequestBody(w http.ResponseWriter, r *http.Request, req interface{}) error {
	return json.NewDecoder(r.Body).Decode(req)
}

func (c BaseController) respond(w http.ResponseWriter, data BaseResponseBody) {
	w.WriteHeader(data.Status)
	json.NewEncoder(w).Encode(data)
}

func (c BaseController) defaultInsertDB() *models.BaseSchema {
	return &models.BaseSchema{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
