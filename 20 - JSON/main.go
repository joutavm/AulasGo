package main

import (
	"encoding/json"
	"log"
)

type dog struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Specie string `json:"specie"`
}

func main() {

	c := dog{
		Name:   "Bolt",
		Age:    5,
		Specie: "Dog",
	}
	jsonDog, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(jsonDog))
}
