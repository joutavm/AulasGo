package main

import "fmt"


func main() {
	var variavel1 int8 = 10
	var variavel2 int8 = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	var ponteiro *int8 = &variavel1

	fmt.Println(*ponteiro)

}