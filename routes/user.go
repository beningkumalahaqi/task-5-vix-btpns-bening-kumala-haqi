package routes

import (
	"task5-vix/controllers"
	"task5-vix/middlewares"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	router := r.PathPrefix("/users").Subrouter()

	router.Use(middlewares.Auth)

	router.HandleFunc("/me", controllers.Me).Methods("GET")
	router.HandleFunc("/update", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/delete", controllers.DeleteUser).Methods("DELETE")
}

func GetAllUser(r *mux.Router) {
	router := r.PathPrefix("/get").Subrouter()

	router.HandleFunc("/all-users", controllers.GetAllProfile).Methods("GET")
}
