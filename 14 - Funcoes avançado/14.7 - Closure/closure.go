package main

func closure() func() {
	x := "Dentro da closure"

	funcao := func() {
		println(x)
	}
	return funcao
}

func main() {
	texto := "dentro da funcao main"
	println(texto)

	funcaoNova := closure()
	funcaoNova()

}
