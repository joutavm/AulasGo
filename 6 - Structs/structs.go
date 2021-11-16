package main

import "fmt"

type usuario struct {
	nome string
	idade uint8
}

func main() {
	var u usuario
	u.nome = "João"
	u.idade = 30

	fmt.Println(u)

	u2 := usuario{"Maria", 20}
	fmt.Println(u2)

	u3 := usuario{nome: "Pedro"}
	fmt.Println(u3)
}