package routes

import (
	"gostart/controllers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func userRoutes(r *mux.Router, db *mongo.Database) {
	userController := controllers.NewUserController(db)

	r.HandleFunc("/register", userController.Register).Methods("POST")
	// r.HandleFunc("/login", userController.Login).Methods("POST")
}
