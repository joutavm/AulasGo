package main

import (
	"crud/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/user", server.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/user", server.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", server.UpdateUser).Methods(http.MethodPut)

	log.Fatalln(http.ListenAndServe(":8080", router))

}
