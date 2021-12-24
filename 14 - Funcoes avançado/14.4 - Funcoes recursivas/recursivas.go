package main

func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}
	return fibonacci(posicao-1) + fibonacci(posicao-2)
}

func main() {
	for i := 0; i < 100; i++ {
		println(fibonacci(uint(i)))
	}
}
