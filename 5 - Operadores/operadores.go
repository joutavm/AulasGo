package main

import "fmt"

func main() {
	// Aritimético
	// mesmos de python

	var numero1 int16 = 10
	var numero2 int16 = 20

	soma := numero1 + numero2

	fmt.Println(soma)

	// Operadores relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)

	// operadores lógicos
	fmt.Println("--------------------")
	fmt.Println(true && false)
	fmt.Println(true || false)

	// operadores unários
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)
	numero--
	numero -= 15
	fmt.Println(numero)
	numero *= 3
	numero /= 15
	numero %= 15
	fmt.Println(numero)

	// 

}