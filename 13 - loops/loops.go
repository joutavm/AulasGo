package main

func main() {
	//var i int
	//for i < 10 {
	//	println(i)
	//	i++
	//	time.Sleep(time.Second)
	//}
	//
	//for j := 0; j < 10; j++ {
	//	println(j)
	//	time.Sleep(time.Second)
	//}

	nomes := []string{"João", "Maria", "José"}

	for i, nome := range nomes {
		println(i, nome)
	}

	for i, letra := range "Golang" {
		println(i, string(letra))
	}

	user := map[string]string{
		"name":  "João",
		"email": "",
	}

	for key, value := range user {
		println(key, value)
	}

}
