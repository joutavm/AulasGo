package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	canal := multiplexator(write("oi"), write("tudo bem?"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexator(chan1, chan2 <-chan string) <-chan string {
	chanExit := make(chan string)

	go func() {
		for {
			select {
			case msg1 := <-chan1:
				chanExit <- msg1
			case msg2 := <-chan2:
				chanExit <- msg2
			}
		}
	}()

	return chanExit
}

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return channel
}
