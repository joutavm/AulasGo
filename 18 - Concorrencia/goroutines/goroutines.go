package main

import (
	"fmt"
	"time"
)

func main() {

	go write("Hello")
	write("Programando em Go")
}

func write(text string) {
	for {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
}
