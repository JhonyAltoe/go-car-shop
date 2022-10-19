package routes

import (
	"github.com/gorilla/mux"
	carController "github.com/jhony/go-car-shop/api/controllers/car"
)

func CarRoutes() *mux.Router {
	// r.Host("www.example.com")
	r := mux.NewRouter()
	r.HandleFunc("/", carController.YourHandler).Methods("GET")
	// http.Handle("/", r)
	return r
}
