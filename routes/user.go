package routes

import (
	"gostart/controllers"

	"github.com/gorilla/mux"
)

func userRoutes(r *mux.Router) {
	r.HandleFunc("/register", controllers.Register).Methods("POST")
}
