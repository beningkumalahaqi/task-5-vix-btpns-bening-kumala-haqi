package routes

import (
	"task5-vix/controllers"
	"task5-vix/middlewares"

	"github.com/gorilla/mux"
)

func PhotoRoutes(r *mux.Router) {
	router := r.PathPrefix("/photos").Subrouter()

	router.Use(middlewares.Auth)

	router.HandleFunc("/", controllers.GetAllPhoto).Methods("GET")

	router.HandleFunc("/See/{id}", controllers.GetPhoto).Methods("GET")
	router.HandleFunc("/upload/{id}", controllers.UploadPhoto).Methods("POST")
	router.HandleFunc("/update/{id}", controllers.UpdatePhoto).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletePhoto).Methods("DELETE")
}
