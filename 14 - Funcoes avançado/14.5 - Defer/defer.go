package main

func alunoEstaAprovado(n1, n2 float32) bool {
	defer println("Médica calculada o resultado será retornado")
	println("Inicio calculo média")
	media := (n1 + n2) / 2
	return media >= 6
}

func main() {
	defer func() {
		println("Fim da execução do programa")
	}()

	println(alunoEstaAprovado(5, 5))
}
