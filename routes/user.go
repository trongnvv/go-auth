package routes

import (
	"goauth/controllers"
	"goauth/middlewares"

	"github.com/gorilla/mux"
)

func userRoutes(r *mux.Router) {
	userController := controllers.NewUserController()

	r.HandleFunc("/register", userController.Register).Methods("POST")
	r.HandleFunc("/login", userController.Login).Methods("POST")

	// sub route need have middleware auth
	sub := r.NewRoute().Subrouter()
	sub.Use(middlewares.IsAuthenticated)

	sub.HandleFunc("/info", userController.Info).Methods("GET")
}
