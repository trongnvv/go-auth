package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"goauth/helpers"

	"github.com/gorilla/mux"
)

func Setup(PORT string) {
	r := mux.NewRouter()

	v1(r)
	v2(r)

	log.Fatal(http.ListenAndServe(":"+PORT, helpers.Logger(r)))
}

func ping(s *mux.Router, version string) {
	s.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		json.NewEncoder(rw).Encode(
			map[string]string{
				"ping":        "Pong",
				"message":     "Service are running...",
				"api-version": version,
			})
	}).Methods("GET")
}

func v1(r *mux.Router) {
	s := r.PathPrefix("/api/v1").Subrouter()
	ping(s, "v1")

	userRoutes(s.PathPrefix("/users").Subrouter())
}

func v2(r *mux.Router) {
	s := r.PathPrefix("/api/v2").Subrouter()
	ping(s, "v2")
}
