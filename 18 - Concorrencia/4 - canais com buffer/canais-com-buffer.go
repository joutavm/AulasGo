package main

import "log"

func main() {

	canal := make(chan string, 2)
	canal <- "Oi"
	canal <- "Oi2"

	message := <-canal
	message2 := <-canal
	log.Println(message)
	log.Println(message2)
}
