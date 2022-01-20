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
	dogJson := `{"name":"Bolt","age":3,"specie":"dog"}`

	var d dog

	if err := json.Unmarshal([]byte(dogJson), &d); err != nil {
		log.Fatalln(err)
	}

	log.Println(d)

	dog2Json := `{"name":"Budega","age":"3","specie":"dog"}`

	d2 := make(map[string]string)

	if err := json.Unmarshal([]byte(dog2Json), &d2); err != nil {
		log.Fatalln(err)
	}

	log.Println(d2)

}
