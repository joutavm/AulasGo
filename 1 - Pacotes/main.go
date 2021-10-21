package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo dentro da main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("jaog.m.a.il.com")
	fmt.Println(erro)
}
