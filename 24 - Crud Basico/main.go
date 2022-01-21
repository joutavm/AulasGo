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

	log.Fatalln(http.ListenAndServe(":8080", router))

}
