package routes

import (
	"gostart/controllers"

	"github.com/gorilla/mux"
)

func userRoutes(r *mux.Router) {
	userController := controllers.NewUserController()

	r.HandleFunc("/register", userController.Register).Methods("POST")
	r.HandleFunc("/login", userController.Login).Methods("POST")
}
