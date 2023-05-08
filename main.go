package main

import (
	"log"
	"net/http"
	"task5-vix/configs"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectDB()

	r := mux.NewRouter()

	router := r.PathPrefix("/api").Subrouter()

	log.Println("Server Runnnig On Port 8080")
	http.ListenAndServe(":8080", router)
}
