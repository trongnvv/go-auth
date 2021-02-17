package controllers

import (
	"context"
	"goauth/helpers"
	"goauth/models"
	"goauth/services"

	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

var ctx = context.TODO()

type UserController struct {
	*BaseController
	service *services.UserService
}

type reqRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type resRegister struct {
	Token string `json:"token"`
}

func NewUserController() *UserController {
	controller := NewBaseController()
	service := services.NewUserService()
	return &UserController{BaseController: controller, service: service}
}

func (c UserController) Register(w http.ResponseWriter, r *http.Request) {
	var req reqRegister

	if err := c.decodeRequestBody(w, r, &req); err != nil || req.Password == "" || req.Username == "" {
		c.respond(w, nil, http.StatusBadRequest, err.Error())
		return
	}

	if err := models.UserModel.FindOne(ctx, bson.M{"username": req.Username}).Decode(&models.UserSchema{}); err == nil {
		c.respond(w, nil, http.StatusBadRequest, "Account existed")
		return
	}

	newUser := models.UserSchema{
		BaseSchema: c.defaultInsertDB(),
		Username:   req.Username,
		Password:   helpers.HashPassword(req.Password),
	}

	if _, err := models.UserModel.InsertOne(ctx, newUser); err != nil {
		c.respond(w, nil, http.StatusInternalServerError, "Insert database error!")
		return
	}

	token, _ := helpers.CreateToken(req.Username, newUser.ID.Hex())
	c.respond(w, resRegister{Token: token}, http.StatusOK, "Request success!")

}

type reqLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type resLogin struct {
	Token string `json:"token"`
}

func (c UserController) Login(w http.ResponseWriter, r *http.Request) {
	var req reqLogin

	if err := c.decodeRequestBody(w, r, &req); err != nil || req.Password == "" || req.Username == "" {
		c.respond(w, nil, http.StatusBadRequest, err.Error())
		return
	}
	user := &models.UserSchema{}
	if err := models.UserModel.FindOne(ctx, bson.M{"username": req.Username}).Decode(user); err != nil {
		c.respond(w, nil, http.StatusBadRequest, "Account not existed")
		return
	}

	if !helpers.ComparePassword(user.Password, req.Password) {
		c.respond(w, nil, http.StatusBadRequest, "Password wrong!")
		return
	}
	token, _ := helpers.CreateToken(req.Username, user.ID.Hex())
	c.respond(w, resLogin{Token: token}, http.StatusOK, "Request success!")
}

type resInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (c UserController) Info(w http.ResponseWriter, r *http.Request) {
	c.respond(
		w,
		resInfo{
			r.Header.Get("user_id"),
			r.Header.Get("username"),
		},
		http.StatusOK, "Request success!",
	)
}
