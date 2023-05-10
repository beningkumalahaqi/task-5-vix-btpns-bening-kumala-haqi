package routes

import (
	"task5-vix/controllers"
	"task5-vix/middlewares"

	"github.com/gorilla/mux"
)

func PhotoRoutes(r *mux.Router) {
	router := r.PathPrefix("/photos").Subrouter()

	router.Use(middlewares.Auth)

	router.HandleFunc("", controllers.GetPhoto).Methods("GET")
	router.HandleFunc("/upload", controllers.UploadPhoto).Methods("POST")
	router.HandleFunc("/update", controllers.UpdatePhoto).Methods("PUT")
	router.HandleFunc("/delete", controllers.DeletePhoto).Methods("DELETE")
}
