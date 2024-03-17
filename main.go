package main

import (
	"log"
	"net/http"

	"db/config"

	"db/handlers"

	"github.com/gorilla/mux"
)

func main() {
	config.Init()

	router := mux.NewRouter()

	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
