package main

import (
	"log"
	"net/http"
	"task5-vix/configs"
	"task5-vix/routes"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()
	router := r.PathPrefix("/api").Subrouter()

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.GetAllUser(router)
	routes.PhotoRoutes(router)

	log.Println("Server Runnnig On Port 8000")
	http.ListenAndServe(":8000", router)
}
