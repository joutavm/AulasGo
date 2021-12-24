package main

func soma(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}

func main() {

	println(soma(1, 2, 3, 4, 5))
}
