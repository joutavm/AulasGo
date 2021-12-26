package main

import "fmt"

func generic(interf interface{}) {
	fmt.Println(interf)

}

func main() {
	generic(1)
	generic("Hello")
	generic(true)
}
