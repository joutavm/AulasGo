package main

import "fmt"

func media(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("MÉDIA É EXATAMENTE 6")

}

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperando execução: ", r)
	}
}

func main() {
	nota1 := 7.0
	nota2 := 5.0

	println(media(nota1, nota2))

	println("Fim do programa")
}
