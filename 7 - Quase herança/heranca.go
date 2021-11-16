package main

import "fmt"

type pessoa struct {
	nome   string
	idade  uint8
	altura float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{	nome: "João"}
	p2 := estudante{pessoa: p1, curso: "a"}
	fmt.Println(p2)
}
