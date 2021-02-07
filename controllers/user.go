package controllers

import (
	"fmt"
	"gostart/helpers"
	"net/http"
)

type reqRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var req reqRegister
	helpers.DecodeRequestBody(w, r, &req)

	// fmt.Println(req)
	fmt.Printf("%+v\n", req)
	helpers.Respond(w, helpers.BaseResponseBody{
		Data:    req,
		Status:  http.StatusOK,
		Message: "Request success!",
	})
}
