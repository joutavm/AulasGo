package main

func invertendoSinal(numero int) int {
	return numero * -1
}

func invertendoSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20

	numeroInvertido := invertendoSinal(numero)

	println(numeroInvertido)
	println(numero)

	novoNumero := 40
	println(novoNumero)
	invertendoSinalPonteiro(&novoNumero)
	println(novoNumero)

}
