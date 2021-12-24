package main

func calculoMatematico(n1, n2 int) (soma int, subtracao int, multiplicacao int, divisao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	multiplicacao = n1 * n2
	divisao = n1 / n2
	return
}

func main() {

	soma, subtracao, multiplicacao, divisao := calculoMatematico(10, 5)

	println(soma, subtracao, multiplicacao, divisao)
}
