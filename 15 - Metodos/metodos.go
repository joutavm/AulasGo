package main

import "fmt"

type user struct {
	name  string
	idade uint8
}

func (u user) getName() string {
	return u.name
}

func (u user) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *user) fazerAniversario() uint8 {
	idade := &u.idade
	*idade++
	return *idade
}

func main() {
	usuario := user{
		name:  "João",
		idade: 30,
	}

	fmt.Println(usuario.getName())
	fmt.Println(usuario.maiorDeIdade())
	usuario.fazerAniversario()
	fmt.Println(usuario.idade)
	fmt.Println(usuario.fazerAniversario())

}
