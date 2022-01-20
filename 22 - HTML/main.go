package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Idade int
}

func main() {

	templates = template.Must(template.ParseGlob("22 - HTML/*.html"))

	http.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {

		u := usuario{
			Nome:  "Gustavo",
			Idade: 32,
		}

		templates.ExecuteTemplate(writer, "home.html", u)
	})

	log.Fatalln(http.ListenAndServe(":5000", nil))
}
