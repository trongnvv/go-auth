package controllers

import (
	"fmt"
	"gostart/services"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserController struct {
	*BaseController
	service *services.UserService
}

type reqRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func NewUserController(db *mongo.Database) *UserController {
	controller := NewBaseController(db)
	service := services.NewUserService(db)
	return &UserController{BaseController: controller, service: service}
}

func (c UserController) Register(w http.ResponseWriter, r *http.Request) {
	var req reqRegister
	c.decodeRequestBody(w, r, &req)

	// fmt.Println(req)
	fmt.Printf("%+v\n", req)
	c.service.Register()
	c.respond(w, BaseResponseBody{
		Data:    req,
		Status:  http.StatusOK,
		Message: "Request success!",
	})
}
