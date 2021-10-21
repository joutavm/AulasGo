package main

import (
	"fmt"
	"errors"
)

func main() {
	// int8, int16, int32, int64
	var numero int64 = 10000000000000
	fmt.Println(numero)

	var numero2 uint = 1000000000
	fmt.Println(numero2)

	// alias 
	// rune == int32
	var numero3 rune = 12345
	fmt.Println(numero3)

	// byte == int8
	var numero4 byte = 255
	fmt.Println(numero4)


	// Numeros reais
	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	// string
	var str string = "texto"
	fmt.Println(str)

	char := '%'
	fmt.Println(char)

	var char2 int = 'a'
	fmt.Println(char2)

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)


}
