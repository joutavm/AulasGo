package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome": "João",
		"idade": "30",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])
	delete(usuario, "sobrenome")
	fmt.Println(usuario)

	usuario["altura"] = "1.45"
	fmt.Println(usuario)
}