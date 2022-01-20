package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", funcName())

	log.Fatalln(http.ListenAndServe(":5000", nil))
}

func funcName() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("Hello World!"))
		if err != nil {
			return
		}
	}
}
