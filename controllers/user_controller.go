package controllers

import (
	"context"
	"goauth/database"
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
	UserModel := database.DB.Collection("users")
	var req reqRegister

	if err := c.decodeRequestBody(w, r, &req); err != nil || req.Password == "" || req.Username == "" {
		c.respond(w, BaseResponseBody{
			Data:    nil,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := UserModel.FindOne(ctx, bson.M{"username": req.Username}).Decode(&models.UserSchema{}); err == nil {
		c.respond(w, BaseResponseBody{
			Data:    nil,
			Status:  http.StatusBadRequest,
			Message: "Account existed",
		})
		return
	}

	newUser := models.UserSchema{
		BaseSchema: c.defaultInsertDB(),
		Username:   req.Username,
		Password:   helpers.HashPassword(req.Password),
	}

	if _, err := UserModel.InsertOne(ctx, newUser); err != nil {
		c.respond(w, BaseResponseBody{
			Data:    nil,
			Status:  http.StatusInternalServerError,
			Message: "Insert database error!",
		})
		return
	}

	token, _ := helpers.CreateToken(req.Username, newUser.ID.String())
	c.respond(w, BaseResponseBody{
		Data:    resRegister{Token: token},
		Status:  http.StatusOK,
		Message: "Request success!",
	})
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
	UserModel := database.DB.Collection("users")

	if err := c.decodeRequestBody(w, r, &req); err != nil || req.Password == "" || req.Username == "" {
		c.respond(w, BaseResponseBody{
			Data:    nil,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	user := &models.UserSchema{}
	if err := UserModel.FindOne(ctx, bson.M{"username": req.Username}).Decode(user); err != nil {
		c.respond(w, BaseResponseBody{
			Data:    nil,
			Status:  http.StatusBadRequest,
			Message: "Account not existed",
		})
		return
	}

	if !helpers.ComparePassword(user.Password, req.Password) {
		c.respond(w, BaseResponseBody{
			Data:    nil,
			Status:  http.StatusBadRequest,
			Message: "Password wrong",
		})
		return
	}
	token, _ := helpers.CreateToken(req.Username, user.ID.String())
	c.respond(w, BaseResponseBody{
		Data:    resLogin{Token: token},
		Status:  http.StatusOK,
		Message: "Request success!",
	})
}
